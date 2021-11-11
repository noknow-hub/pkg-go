//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package localization

import (
    "encoding/json"
    "io/ioutil"
    "strconv"
    "strings"
)

var (
    Locales = map[string]map[string]interface{}{}
)


//////////////////////////////////////////////////////////////////////
// Set a localization file.
//////////////////////////////////////////////////////////////////////
func SetWithJsonFile(langCode, filePath string) error {
    bytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return err
    }
    var v map[string]interface{}
    if err = json.Unmarshal(bytes, &v); err != nil {
        return err
    }
    Locales[langCode] = v
    return nil
}


//////////////////////////////////////////////////////////////////////
// Get localized strings.
//////////////////////////////////////////////////////////////////////
func Strings(langCode string) map[string]interface{} {
    if targetStrings, ok := Locales[langCode]; ok {
        return targetStrings
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Get string value from key.
//////////////////////////////////////////////////////////////////////
func String(key, langCode string) string {
    if Locales[langCode] == nil {
        return key
    }
    return getString(key, langCode)
}


//////////////////////////////////////////////////////////////////////
// Get string value from key.
//////////////////////////////////////////////////////////////////////
func getString(key, langCode string) string {
    v := getInterfaceValue(key, langCode)
    if v == nil {
        return ""
    }
    switch result := v.(type) {
    case string:
        return result
    case bool:
        return strconv.FormatBool(result)
    case int:
        return strconv.FormatInt(int64(result), 10)
    case int8:
        return strconv.FormatInt(int64(result), 10)
    case int16:
        return strconv.FormatInt(int64(result), 10)
    case int32:
        return strconv.FormatInt(int64(result), 10)
    case int64:
        return strconv.FormatInt(result, 10)
    case float32:
        return strconv.FormatFloat(float64(result), 'f', -1, 32)
    case float64:
        return strconv.FormatFloat(result, 'f', -1, 64)
    case uint:
        return strconv.FormatUint(uint64(result), 10)
    case uint8:
        return strconv.FormatUint(uint64(result), 10)
    case uint16:
        return strconv.FormatUint(uint64(result), 10)
    case uint32:
        return strconv.FormatUint(uint64(result), 10)
    case uint64:
        return strconv.FormatUint(result, 10)
    default:
        return ""
    }
}


//////////////////////////////////////////////////////////////////////
// Get interface value from key.
//////////////////////////////////////////////////////////////////////
func getInterfaceValue(key, langCode string) interface{} {
    var v interface{}
    for i, k := range strings.Split(key, ".") {
        if i == 0 {
            v = Locales[langCode][k]
        } else {
            switch result := v.(type) {
            case map[interface{}]interface{}:
            case map[string]interface{}:
                v = result[k]
            default:
                v = nil
                break
            }
        }
    }
    return v
}
