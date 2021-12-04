# noknow-hub/pkg-go/i18n

## Getting Started

### Import

```go
import (
    "github.com/noknow-hub/pkg-go/i18n"
)
```

## Usage

### Set an English localization file

Here is `en.json` .

```json
{
  "app": {
    "lang": "English"
  }
}
```

```go
langCode := "en"
enFilePath := "/path/to/en.json"
i18n.SetWithJsonFile(langCode, enFilePath)

fmt.Println(i18n.String("app.lang", "en"))
// Output: English
fmt.Println(i18n.String("app.lang", "ja"))
// Output: app.lang
```
