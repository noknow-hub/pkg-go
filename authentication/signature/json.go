//////////////////////////////////////////////////////////////////////
// json.go
//////////////////////////////////////////////////////////////////////
package signature

import (
    "bytes"
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "hash"
    "sort"
    "time"
)


const (
    MD5 Algorithm = 1 + iota
    SHA256
    DataKeySignature = "signature"
    DataKeyTimestamp = "timestamp"
    ResultCodeVerified = 0
    ResultCodeInvalid = 1
    ResultCodeExpired = 2
)


type Algorithm int


type JsonClient struct {
    Algorithm
    SecretKey string
}


type Result struct {
    Code int
    Message string
}



//////////////////////////////////////////////////////////////////////
// New JsonClient
//////////////////////////////////////////////////////////////////////
func NewJsonClient(algorithm Algorithm) *JsonClient {
    return &JsonClient{
        Algorithm: algorithm,
    }
}




//////////////////////////////////////////////////////////////////////
// Signature with HMAC
//////////////////////////////////////////////////////////////////////
func (j *JsonClient) SignatureWithHMAC(jsonData []byte, key string) (string, error) {
    var data map[string]interface{}
    if err := json.Unmarshal(jsonData, &data); err != nil {
        return "", err
    }

    // Check the timestamp key
    if data[DataKeyTimestamp] == nil {
        return "", fmt.Errorf("\"%s\" key is required.\n", DataKeyTimestamp)
    }

    // Sort alphabetically
    keys := make([]string, 0, len(data))
    for k := range data {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    sortedData := map[string]interface{}{}
    for _, key := range keys {
        sortedData[key] = data[key]
    }

    // Generate a signature
    buffer := &bytes.Buffer{}
    encoder := json.NewEncoder(buffer)
    encoder.SetEscapeHTML(false)
    if err := encoder.Encode(sortedData); err != nil {
        return "", err
    }
    var h hash.Hash
    if j.Algorithm == MD5 {
        h = hmac.New(md5.New, []byte(j.SecretKey))
    } else if j.Algorithm == SHA256 {
        h = hmac.New(sha256.New, []byte(j.SecretKey))
    } else {
        return "", fmt.Errorf("The specified algorithm is not supported. value=%d\n", j.Algorithm)
    }
    h.Write(buffer.Bytes())
    return hex.EncodeToString(h.Sum(nil)), nil
}




//////////////////////////////////////////////////////////////////////
// Verify with HMAC
//////////////////////////////////////////////////////////////////////
func (j *JsonClient) VerifyWithHMAC(jsonData []byte, key string, expirationSeconds time.Duration) *Result {
    result := &Result{
        Code: ResultCodeVerified,
    }
    var data map[string]interface{}
    if err := json.Unmarshal(jsonData, &data); err != nil {
        result.Code = ResultCodeInvalid
        result.Message = err.Error()
        return result
    }

    // Check the required parameters
    if data[DataKeySignature] == nil || data[DataKeyTimestamp] == nil {
        result.Code = ResultCodeInvalid
        result.Message = fmt.Sprintf("The required parameters are not found. \"%s\" and \"%s\" are required.", DataKeySignature, DataKeyTimestamp)
        return result
    }

    // Check if the timestamp is currently
    var timestamp int64
    switch v := data[DataKeyTimestamp].(type) {
    case int64:
        timestamp = v
    case float64:
        timestamp = int64(v)
    default:
        result.Code = ResultCodeInvalid
        result.Message = fmt.Sprintf("The \"%s\" key is invalid.", DataKeyTimestamp)
        return result
    }
    expiredTime := time.Unix(timestamp, 0).Add(expirationSeconds).UTC()
    if time.Now().UTC().After(expiredTime) {
        result.Code = ResultCodeExpired
        result.Message = "Signature is expired."
        return result 
    }

    // Sort alphabetically
    keys := make([]string, 0, len(data))
    for k := range data {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    sortedData := map[string]interface{}{}
    for _, key := range keys {
        if key == DataKeySignature {
            continue
        }
        sortedData[key] = data[key]
    }

    sig := data[DataKeySignature].(string)
    if sig == "" {
        result.Code = ResultCodeInvalid
        result.Message = "Signature is invalid."
        return result
    }

    // Signature with the sorted data
    sortedJsonData, err := json.Marshal(sortedData)
    if err != nil {
        result.Code = ResultCodeInvalid
        result.Message = err.Error()
        return result
    }
    sortedJsonSig, err := j.SignatureWithHMAC(sortedJsonData, key)
    if err != nil {
        result.Code = ResultCodeInvalid
        result.Message = err.Error()
        return result
    }

    if sig != sortedJsonSig {
        result.Code = ResultCodeInvalid
        result.Message = "Signature is invalid."
        return result
    }

    return result
}
