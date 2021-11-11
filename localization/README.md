# noknow-hub/pkg-go/localization

## Getting Started

### Import

```go
import (
    "github.com/noknow-hub/pkg-go/localization"
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
localization.SetWithJsonFile(langCode, enFilePath)

fmt.Println(localization.String("app.lang", "en"))
// Output: English
fmt.Println(localization.String("app.lang", "ja"))
// Output: app.lang
```
