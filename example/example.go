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
