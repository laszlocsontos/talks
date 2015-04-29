package main

import (
	"io"
	"os"
	"strings"
)

type DecipherReader struct {
	r io.Reader
}

const OFFSET = 13

func (dr *DecipherReader) Read(p []byte) (n int, err error) {
	dr.r.Read(p)

	for i, b := range p {
		switch {
		case b >= 'a' && b <= 'z':
			p[i] = (b-'a'+OFFSET)%26 + 'a'
		case b >= 'A' && b <= 'Z':
			p[i] = (b-'A'+OFFSET)%26 + 'A'
		default:
			continue
		}
	}

	return len(p), io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := DecipherReader{s}
	io.Copy(os.Stdout, &r)
}

