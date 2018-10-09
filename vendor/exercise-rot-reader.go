package main

import (
	"io"
	"os"
	"strings"
)

func exchange_rot13_convert(b byte) byte  {
	var a,z byte
	switch  {
	case 'a' <= b && b <= 'z':
		a,z = 'a','z'
	case 'A' <= b && b <= 'Z':
		a,z = 'A','Z'
	default:
		return b
	}

	return (b-a+13)%(z-a+1) + a

}

type rot13Reader struct {
	r io.Reader
}

func (read rot13Reader) Read(b []byte) (n int,err error)  {
	n,err = read.r.Read(b)
	for i:=0 ; i<n;i++ {
		b[i] = exchange_rot13_convert(b[i])
	}
	
	return 
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
