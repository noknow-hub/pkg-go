//////////////////////////////////////////////////////////////////////
// result.go
//////////////////////////////////////////////////////////////////////
package v3

import (
    "reflect"
)

const (
    STATUS_CODE_OK = 20000
    STATUS_CODE_TASK_CREATED = 20100
)

type General struct {
    Version string        `json:"version"`
    StatusCode int        `json:"status_code"`
    StatusMessage string  `json:"status_message"`
    Time string           `json:"time"`
    Cost float64          `json:"cost"`
    TasksCount int        `json:"tasks_count"`
    TasksError int        `json:"tasks_error"`
}

type Task struct {
    Id string                    `json:"id"`
    StatusCode int              `json:"status_code"`
    StatusMessage string        `json:"status_message"`
    Time string                  `json:"time"`
    Cost float64                 `json:"cost"`
    ResultCount int              `json:"result_count"`
    Path []string                `json:"path"`
    Data map[string]interface{}  `json:"data"`
}


//////////////////////////////////////////////////////////////////////
// Is status code "OK"
//////////////////////////////////////////////////////////////////////
func (t *Task) IsOk() bool {
    return t.StatusCode == STATUS_CODE_OK
}


//////////////////////////////////////////////////////////////////////
// Is status code "task created"
//////////////////////////////////////////////////////////////////////
func (t *Task) IsTaskCreated() bool {
    return t.StatusCode == STATUS_CODE_TASK_CREATED
}


//////////////////////////////////////////////////////////////////////
// Get "device" from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetDeviceFromData() string {
    if t.Data == nil {
        return ""
    }
    v, exist := t.Data["device"]
    if !exist {
        return ""
    }
    vv, ok := v.(string)
    if !ok {
        return ""
    }
    return vv
}


//////////////////////////////////////////////////////////////////////
// Get "item_types" slice from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetItemTypesSliceFromData() []string {
    if t.Data == nil {
        return nil
    }
    v, exist := t.Data["item_types"]
    if !exist {
        return nil
    }
    vv := reflect.ValueOf(v)
    if vv.Kind() != reflect.Slice {
        return nil
    }
    var result []string
    for i := 0; i < vv.Len(); i++ {
        vvv := vv.Index(i)
        if vvv.Kind() == reflect.Interface {
            if vvvv, ok := vvv.Interface().(string); ok {
                result = append(result, vvvv)
            }
        }
    }
    return result
}


//////////////////////////////////////////////////////////////////////
// Get "keyword" from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetKeywordFromData() string {
    if t.Data == nil {
        return ""
    }
    v, exist := t.Data["keyword"]
    if !exist {
        return ""
    }
    vv, ok := v.(string)
    if !ok {
        return ""
    }
    return vv
}


//////////////////////////////////////////////////////////////////////
// Get "keywords" slice from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetKeywordsSliceFromData() []string {
    if t.Data == nil {
        return nil
    }
    v, exist := t.Data["keywords"]
    if !exist {
        return nil
    }
    vv := reflect.ValueOf(v)
    if vv.Kind() != reflect.Slice {
        return nil
    }
    var result []string
    for i := 0; i < vv.Len(); i++ {
        vvv := vv.Index(i)
        if vvv.Kind() == reflect.Interface {
            if vvvv, ok := vvv.Interface().(string); ok {
                result = append(result, vvvv)
            }
        }
    }
    return result
}


//////////////////////////////////////////////////////////////////////
// Get "language_code" from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetLanguageCodeFromData() string {
    if t.Data == nil {
        return ""
    }
    v, exist := t.Data["language_code"]
    if !exist {
        return ""
    }
    vv, ok := v.(string)
    if !ok {
        return ""
    }
    return vv
}


//////////////////////////////////////////////////////////////////////
// Get "location_code" from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetLocationCodeFromData() int {
    if t.Data == nil {
        return 0
    }
    v, exist := t.Data["location_code"]
    if !exist {
        return 0
    }
    vv, ok := v.(float64)
    if !ok {
        return 0
    }
    return int(vv)
}


//////////////////////////////////////////////////////////////////////
// Get "os" from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetOsFromData() string {
    if t.Data == nil {
        return ""
    }
    v, exist := t.Data["os"]
    if !exist {
        return ""
    }
    vv, ok := v.(string)
    if !ok {
        return ""
    }
    return vv
}


//////////////////////////////////////////////////////////////////////
// Get "se" from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetSeFromData() string {
    if t.Data == nil {
        return ""
    }
    v, exist := t.Data["se"]
    if !exist {
        return ""
    }
    vv, ok := v.(string)
    if !ok {
        return ""
    }
    return vv
}


//////////////////////////////////////////////////////////////////////
// Get "se_type" from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetSeTypeFromData() string {
    if t.Data == nil {
        return ""
    }
    v, exist := t.Data["se_type"]
    if !exist {
        return ""
    }
    vv, ok := v.(string)
    if !ok {
        return ""
    }
    return vv
}
