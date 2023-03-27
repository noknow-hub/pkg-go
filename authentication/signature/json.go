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
)


type Algorithm int


type JsonClient struct {
    Algorithm
    SecretKey string
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
func (j *JsonClient) VerifyWithHMAC(jsonData []byte, key string, expirationSeconds time.Duration) (bool, error) {
    var data map[string]interface{}
    if err := json.Unmarshal(jsonData, &data); err != nil {
        return false, err
    }

    // Check the required parameters
    if data[DataKeySignature] == nil || data[DataKeyTimestamp] == nil {
        return false, fmt.Errorf("The required parameters are not found. \"%s\" and \"%s\" are required.\n", DataKeySignature, DataKeyTimestamp)
    }

    // Check if the timestamp is currently
    var timestamp int64
    switch v := data[DataKeyTimestamp].(type) {
    case int64:
        timestamp = v
    case float64:
        timestamp = int64(v)
    default:
        return false, fmt.Errorf("The \"%s\" key is invalid.\n", DataKeyTimestamp)
    }
    expiredTime := time.Unix(timestamp, 0).Add(expirationSeconds).UTC()
    if time.Now().UTC().After(expiredTime) {
        return false, fmt.Errorf("Signature is expired.\n")
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
        return false, fmt.Errorf("Signature is invalid.\n")
    }

    // Signature with the sorted data
    sortedJsonData, err := json.Marshal(sortedData)
    if err != nil {
        return false, err
    }
    sortedJsonSig, err := j.SignatureWithHMAC(sortedJsonData, key)
    if err != nil {
        return false, err
    }

    if sig != sortedJsonSig {
        return false, fmt.Errorf("Signature is invalid.\n")
    }

    return true, nil
}
