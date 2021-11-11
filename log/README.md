# noknow-hub/pkg-go/log

## Getting Started

### Import

```go
import (
    "github.com/noknow-hub/pkg-go/log"
)
```

## Usage

### Logging for DEBUG

```go
log.Debug("Hello World")
// 2021/11/07 00:25:25 main.go:14: [DEBUG] Hello World
```

### Logging for INFO

```go
log.Info("Hello World")
// 2021/11/07 00:25:25 main.go:14: [INFO] Hello World
```

### Logging for WARN

```go
log.Warn("Hello World")
// 2021/11/07 00:25:25 main.go:14: [WARN] Hello World
```

### Logging for ERROR

```go
log.Error("Hello World")
// 2021/11/07 00:25:25 main.go:14: [ERROR] Hello World
```

### Logging for FATAL

```go
log.Fatal("Hello World")
// 2021/11/07 00:25:25 main.go:14: [FATAL] Hello World
// exit status 1
```

### Set output with file path

```go
filePath := "app.log"
log.SetOutputWithFilePath(filePath)
```
