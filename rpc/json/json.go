//////////////////////////////////////////////////////////////////////
// json.go
//////////////////////////////////////////////////////////////////////
package json

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
)

var (
    JSON_RPC_VERSION = "2.0"
)

type ReqData struct {
    Id int              `json:"id"`
    Method string       `json:"method"`
    JsonRpc string      `json:"jsonrpc"`
    Params interface{}  `json:"params"`
}


//////////////////////////////////////////////////////////////////////
// Get ReqData from json request.
//////////////////////////////////////////////////////////////////////
func GetReqDataFromJsonRequest(r *http.Request) (*ReqData, error) {
    if r.Body == nil {
        return nil, errors.New("HTTP body is invalid.")
    }
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return nil, err
    }
    var reqData *ReqData
    if err := json.Unmarshal(body, &reqData); err != nil {
        return nil, err
    }
    return reqData, nil
}


//////////////////////////////////////////////////////////////////////
// Get the string value of a key in params.
//////////////////////////////////////////////////////////////////////
func GetParamsString(params interface{}, key string) (string, error) {
    paramsMap := params.(map[string]interface{})
    v, ok := paramsMap[key].(string)
    if !ok {
        return "", errors.New("JSON RPC format for params is invalid.")
    }
    return v, nil
}


//////////////////////////////////////////////////////////////////////
// Response Empty Data.
//////////////////////////////////////////////////////////////////////
func ResEmpty(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(""))
}


//////////////////////////////////////////////////////////////////////
// Response Success.
//////////////////////////////////////////////////////////////////////
func ResSuccess(id int, success interface{}, w http.ResponseWriter) {
    jsonResMap := map[string]interface{}{
        "id": id,
        "jsonrpc": JSON_RPC_VERSION,
        "result": success,
    }
    jsonRes, err := json.Marshal(jsonResMap)
    if err != nil {
        ResEmpty(w)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonRes)
}


//////////////////////////////////////////////////////////////////////
// Response Error.
//////////////////////////////////////////////////////////////////////
func ResError(id int, error interface{}, w http.ResponseWriter) {
    jsonResMap := map[string]interface{}{
        "id": id,
        "jsonrpc": JSON_RPC_VERSION,
        "error": error,
    }
    jsonRes, err := json.Marshal(jsonResMap)
    if err != nil {
        ResEmpty(w)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonRes)
}


//////////////////////////////////////////////////////////////////////
// Response Error for Authenticate Bearer.
// It conforms to RFC6750.
//////////////////////////////////////////////////////////////////////
func ResErrorBearer(w http.ResponseWriter, statusCode int, realm string, errDesc string, errUri string) {
    buffer := make([]byte, 0)
    buffer = append(buffer, "Bearer realm=\"" + realm + "\""...)

    if statusCode == http.StatusBadRequest {
        buffer = append(buffer, ", error=\"invalid_request\""...)
    } else if statusCode == http.StatusUnauthorized {
        buffer = append(buffer, ", error=\"invalid_token\""...)
    } else if statusCode == http.StatusForbidden {
        buffer = append(buffer, ", error=\"insufficient_scope\""...)
    } else if statusCode == http.StatusMethodNotAllowed {
        buffer = append(buffer, ", error=\"invalid_request\""...)
    } else {
        buffer = append(buffer, ", error=\"invalid_request\""...)
    }
    if errDesc != "" {
        buffer = append(buffer, ", error_description=\"" + errDesc + "\""...)
    }
    if errUri != "" {
        buffer = append(buffer, ", error_uri=\"" + errUri + "\""...)
    }
    w.Header().Set("WWW-Authenticate", string(buffer[:]))
    if statusCode == http.StatusBadRequest {
        w.WriteHeader(http.StatusBadRequest)
    } else if statusCode == http.StatusUnauthorized {
        w.WriteHeader(http.StatusUnauthorized)
    } else if statusCode == http.StatusForbidden {
        w.WriteHeader(http.StatusForbidden)
    } else if statusCode == http.StatusMethodNotAllowed {
         w.WriteHeader(http.StatusMethodNotAllowed)
    } else {
        w.WriteHeader(http.StatusBadRequest)
    }
    w.Write([]byte(""))
}


//////////////////////////////////////////////////////////////////////
// Set JSON RPC version.
//////////////////////////////////////////////////////////////////////
func SetJsonRpcVersion(version string) {
    JSON_RPC_VERSION = version
}


//////////////////////////////////////////////////////////////////////
// Get JSON RPC version.
//////////////////////////////////////////////////////////////////////
func GetJsonRpcVersion() string {
    return JSON_RPC_VERSION
}
