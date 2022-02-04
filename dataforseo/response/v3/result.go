//////////////////////////////////////////////////////////////////////
// general.go
//////////////////////////////////////////////////////////////////////
package v3

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
    Status_code int              `json:"status_code"`
    Status_message string        `json:"status_message"`
    Time string                  `json:"time"`
    Cost float64                 `json:"cost"`
    ResultCount int              `json:"result_count"`
    Path []string                `json:"path"`
    Data map[string]interface{}  `json:"data"`
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
// Get "keyword" slice from "Task.Data".
//////////////////////////////////////////////////////////////////////
func (t *Task) GetKeywordSliceFromData() []string {
    if t.Data == nil {
        return nil
    }
    v, exist := t.Data["keyword"]
    if !exist {
        return nil
    }
    vv, ok := v.([]string)
    if !ok {
        return nil
    }
    return vv
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
