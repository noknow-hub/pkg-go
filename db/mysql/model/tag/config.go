//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package tag

const (
    COL_ID = "id"
    COL_LABEL = "label"
    COL_NAME = "name"
    NUM_COLS = 3
    TABLE_NAME = "tags"
)

type Tag struct {
    Id string
    Name string
    Label string
}
