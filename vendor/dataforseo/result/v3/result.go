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
