//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package tag

const (
    COL_NAME = "name"
    COL_PARENT_SLUG = "parent_slug"
    COL_SLUG = "slug"
    NUM_COLS = 3
    TABLE_NAME = "tags"
)

type Tag struct {
    Slug string
    Name string
    ParentSlug string
}
