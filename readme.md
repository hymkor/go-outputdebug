OutputDebugString for Go
========================

We can view debug-log with debugger or viewer.(Recommend [DebugView for Windows](https://technet.microsoft.com/ja-jp/sysinternals/debugview.aspx) )

[Example](./example/example.go)
--------------------------------

```
package main

import (
	dbg "../../goutputdebugstring"
)

func main() {
	dbg.Print("example.go: main() Call dbg.Print")
}
```
