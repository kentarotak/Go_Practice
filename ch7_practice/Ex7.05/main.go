package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

var EOF = errors.New("EOF")

type Limit struct {
	R io.Reader
	N int64
}

func (l *Limit) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &Limit{r, n}
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
