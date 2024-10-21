//////////////////////////////////////////////////////////////////////
// yaml.go
//////////////////////////////////////////////////////////////////////
package yaml

import (
    "flag"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)


var (
    Config  = make(map[string]interface{})
)


//////////////////////////////////////////////////////////////////////
// Initialize.
//////////////////////////////////////////////////////////////////////
func Init() {
    var yamlPaths []string
    flag.Func("yaml", "Path to yaml config file (can be specified multiple times)", func(s string) error {
        yamlPaths = append(yamlPaths, s)
        return nil
    })
    flag.Parse()

    if len(yamlPaths) == 0 {
        log.Fatalf("[FATAL] Need at least one -yaml option. The value must be yaml type. Usage: BUILDED_APP_FILE -yaml /yaml_path/config.yaml. -yaml /yaml_path/config2.yaml.\n")
    }

    for _, path := range yamlPaths {
        bytes, err := ioutil.ReadFile(path)
        if err != nil {
            log.Fatalf("[FATAL] %s\n", err)
        }
        tmpConfig := make(map[string]interface{})
        if err = yaml.Unmarshal(bytes, &tmpConfig); err != nil {
            log.Fatalf("[FATAL] %s\n", err)
        }
        mergeConfig(Config, tmpConfig)
    }

//    env := flag.String("yaml", "", "/yaml_path/config.yaml")
//    envList := flag.String("yamls", "", "/yaml_path/config.yaml,/yaml_path/config.yaml")
//    flag.Parse()
//    if *env == "" && *envList == "" {
//    }
// 
//    if *env != "" {
//        bytes, err := ioutil.ReadFile(*env)
//        if err != nil {
//            log.Fatalf("[FATAL] %s\n", err)
//        }
//        if err = yaml.Unmarshal(bytes, &Config); err != nil {
//            log.Fatalf("[FATAL] %s\n", err)
//        }
//    } else {
//        for _, filePath := range strings.Split(*envList, ",") {
//            bytes, err := ioutil.ReadFile(filePath)
//            if err != nil {
//                log.Fatalf("[FATAL] %s\n", err)
//            }
//            tempConfig := make(map[interface{}]interface{})
//            if err = yaml.Unmarshal(bytes, &tempConfig); err != nil {
//                log.Fatalf("[FATAL] %s\n", err)
//            }
//            mergeConfig(Config, tempConfig)
//        }
//    }
}


//////////////////////////////////////////////////////////////////////
// Initialize from the specified file path
//////////////////////////////////////////////////////////////////////
func InitFromFilePath(filePath string) {
    bytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatalf("[FATAL] %s\n", err)
    }
    if err = yaml.Unmarshal(bytes, &Config); err != nil {
        log.Fatalf("[FATAL] %s\n", err)
    }
}


//////////////////////////////////////////////////////////////////////
// Get boolean value from key.
//////////////////////////////////////////////////////////////////////
func GetBool(key string) bool {
    v := getInterfaceValue(key)
    if v == nil {
        return false
    }
    switch result := v.(type) {
    case bool:
        return result
    default:
        return false
    }
}


//////////////////////////////////////////////////////////////////////
// Get string value from key.
//////////////////////////////////////////////////////////////////////
func GetString(key string) string {
    v := getInterfaceValue(key)
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
// Get int value from key.
//////////////////////////////////////////////////////////////////////
func GetInt(key string) int {
    v := getInterfaceValue(key)
    if v == nil {
        return 0
    }
    switch result := v.(type) {
    case int:
        return result
    case int8:
        return int(result)
    case int16:
        return int(result)
    case int32:
        return int(result)
    case int64:
        return int(result)
    case uint:
        return int(result)
    case uint8:
        return int(result)
    case uint16:
        return int(result)
    case uint32:
        return int(result)
    case uint64:
        return int(result)
    case float32:
        return int(result)
    case float64:
        return int(result)
    default:
        return 0
    }
}


//////////////////////////////////////////////////////////////////////
// Get int64 value from key.
//////////////////////////////////////////////////////////////////////
func GetInt64(key string) int64 {
    v := getInterfaceValue(key)
    if v == nil {
        return 0
    }
    switch result := v.(type) {
    case int64:
        return result
    case int:
        return int64(result)
    case int8:
        return int64(result)
    case int16:
        return int64(result)
    case int32:
        return int64(result)
    case uint8:
        return int64(result)
    case uint16:
        return int64(result)
    case uint32:
        return int64(result)
    case uint64:
        return int64(result)
    case float32:
        return int64(result)
    case float64:
        return int64(result)
    default:
        return 0
    }
}


//////////////////////////////////////////////////////////////////////
// Get float64 value from key.
//////////////////////////////////////////////////////////////////////
func GetFloat64(key string) float64 {
    v := getInterfaceValue(key)
    if v == nil {
        return 0
    }
    switch result := v.(type) {
    case float64:
        return result
    case float32:
        return float64(result)
    case int:
        return float64(result)
    case int8:
        return float64(result)
    case int16:
        return float64(result)
    case int32:
        return float64(result)
    case int64:
        return float64(result)
    case uint:
        return float64(result)
    case uint8:
        return float64(result)
    case uint16:
        return float64(result)
    case uint32:
        return float64(result)
    case uint64:
        return float64(result)
    default: 
        return 0
    }
}


//////////////////////////////////////////////////////////////////////
// Get array value from key.
//////////////////////////////////////////////////////////////////////
func GetArray(key string) []interface{} {
    v := getInterfaceValue(key)
    if v == nil {
        return nil
    }
    switch result := v.(type) {
    case []interface{}:
        return result
    default:
        return nil
    }
}


//////////////////////////////////////////////////////////////////////
// Get array int value from key.
//////////////////////////////////////////////////////////////////////
func GetArrayInt(key string) []int {
    v := GetArray(key)
    if v == nil {
        return nil
    }
    var result []int
    for _, vv := range v {
        vvv, ok := vv.(int)
        if ok {
            result = append(result, vvv)
        }
    }
    return result
}


//////////////////////////////////////////////////////////////////////
// Get array string value from key.
//////////////////////////////////////////////////////////////////////
func GetArrayString(key string) []string {
    v := GetArray(key)
    if v == nil {
        return nil
    }
    var result []string
    for _, vv := range v {
        switch vvv := vv.(type) {
        case string:
            result = append(result, vvv)
        case bool:
            result = append(result, strconv.FormatBool(vvv))
        case int:
            result = append(result, strconv.FormatInt(int64(vvv), 10))
        case int8:
            result = append(result, strconv.FormatInt(int64(vvv), 10))
        case int16:
            result = append(result, strconv.FormatInt(int64(vvv), 10))
        case int32:
            result = append(result, strconv.FormatInt(int64(vvv), 10))
        case int64:
            result = append(result, strconv.FormatInt(vvv, 10))
        case float32:
            result = append(result, strconv.FormatFloat(float64(vvv), 'f', -1, 32))
        case float64:
            result = append(result, strconv.FormatFloat(vvv, 'f', -1, 64))
        case uint:
            result = append(result, strconv.FormatUint(uint64(vvv), 10))
        case uint8:
            result = append(result, strconv.FormatUint(uint64(vvv), 10))
        case uint16:
            result = append(result, strconv.FormatUint(uint64(vvv), 10))
        case uint32:
            result = append(result, strconv.FormatUint(uint64(vvv), 10))
        case uint64:
            result = append(result, strconv.FormatUint(vvv, 10))
        default:
        }
    }
    return result
}


//////////////////////////////////////////////////////////////////////
// Get interface value from key.
//////////////////////////////////////////////////////////////////////
func getInterfaceValue(key string) interface{} {
    var v interface{}
    for i, k := range strings.Split(key, ".") {
        if i == 0 {
            v = Config[k]
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


//////////////////////////////////////////////////////////////////////
// Merge Config
//////////////////////////////////////////////////////////////////////
func mergeConfig(srcConfig, overrideConfig map[string]interface{}) {
    for key, value := range overrideConfig {
        if srcVal, ok := srcConfig[key]; ok {
            // If the key exists, cast data to map type.
            if srcMap, ok := srcVal.(map[string]interface{}); ok {
                // If the override value is map, merge this map data recursively.
                if overrideMap, ok := value.(map[string]interface{}); ok {
                    mergeConfig(srcMap, overrideMap)
                    continue
                }
            }
        }
        // Set data if the key does not exists.
        srcConfig[key] = value
    }
}
