package outputdebug

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")
var outputDebugStringW = kernel32.NewProc("OutputDebugStringW")

func String(s string) {
	p, err := syscall.UTF16PtrFromString(s)
	if err == nil {
		outputDebugStringW.Call(uintptr(unsafe.Pointer(p)))
	}
}
