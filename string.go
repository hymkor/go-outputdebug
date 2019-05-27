package outputdebug

import (
	"bytes"
)

func String(s string) {
	outputdebugstring(s)
}

type out struct {
	buffer []byte
}

func (d *out) Write(b []byte) (int, error) {
	d.buffer = append(d.buffer, b...)
	for {
		eol := bytes.IndexByte(d.buffer, '\n')
		if eol < 0 {
			return len(b), nil
		}
		String(string(d.buffer[:eol]))

		left := len(d.buffer) - (eol + 1)
		copy(d.buffer, d.buffer[eol+1:])
		d.buffer = d.buffer[:left]
	}
}

var Out = &out{}
