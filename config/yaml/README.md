# noknow-hub/pkg-go/config/yaml

## Getting Started

### Import

```go
import (
    v1 "github.com/noknow-hub/pkg-go/config/yaml"
)
```

### Initialize

You must option `-yaml [your_config_yaml_file_path]` when executing your application.

```go
v1.Init()
```



## Example

Here is sample yaml file.

```yaml
---
bool: true
int: 100
int64: 1234567890123456789
string:
  one_line: string text
  block1: |
      this is block string
      you can break lines.
  block2: >
      this is another block string.
       you can break lines.
array:
  int: [1, 2, 3, 4, 5]
  string: [hello, world]
items:
  - no: 1
    name: aaa
  - no: 2
    name: bbb
```

When executing the sample go.

```console
go run main.go -yaml ./local.yaml
```

```go
[INFO] Config:
map[array:map[int:[1 2 3 4 5] string:[hello world]] bool:true int:100 int64:1234567890123456789 items:[map[name:aaa no:1] map[name:bbb no:2]] string:map[block1:this is block string
you can break lines.
 block2:this is another block string.
 you can break lines.
 one_line:string text]]
----------
[INFO] GetBool("bool"): true
----------
[INFO] GetInt("int"): 100
----------
[INFO] GetInt("int64"): 1234567890123456789
----------
[INFO] GetString("string.one_line"): string text
----------
[INFO] GetString("string.block1"): this is block string
you can break lines.

----------
[INFO] GetString("string.block2"): this is another block string.
 you can break lines.

----------
[INFO] GetArrayInt("array.int"): [1 2 3 4 5]
----------
[INFO] GetArrayString("array.string"): [hello world]
----------
[INFO] GetArray("items"): [map[name:aaa no:1] map[name:bbb no:2]]
----------
```
