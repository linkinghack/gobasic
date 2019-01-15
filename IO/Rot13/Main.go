package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(v byte) byte {
	var rtn byte

	if v < 'A' || v > 'z' {
		return v
	}

	switch {
	case v >= 'A' && v <= 'Z':
		rtn = v + 13
	case v >= 'a' && v <= 'z':
		rtn = v + 13
	default:
		break
	}

	if rtn > 'Z' && rtn < 'a' || rtn > 'z' {
		rtn = rtn - 26
	}
	return rtn
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err == nil {
		for i := 0; i < n; i++ {
			b[i] = rot13(b[i])
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!?")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
