go-outputdebug
==============

We can view debug-log with debugger or viewer.(Recommend [DebugView for Windows](https://technet.microsoft.com/ja-jp/sysinternals/debugview.aspx) )

[Example](./example/example.go)
--------------------------------

```go
package main

import (
    "fmt"
    "github.com/zetamatta/go-outputdebug"
)

var dbg = outputdebug.Out

func main() {
    outputdebug.String("outputdebug.String()")
    fmt.Fprintf(dbg, "example.go: main() Call dbg.Print\n")
    fmt.Fprintf(dbg, "foo\nbar\ndddddd\n")
}
```
