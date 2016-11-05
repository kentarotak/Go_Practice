package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWrite struct {
	w   io.Writer
	cnt int64
}

func (c *countingWrite) Write(p []byte) (int, error) {

	c.cnt += int64(len(p))
	n, err := c.w.Write(p)

	return n, err

}

func (c *countingWrite) Set(w io.Writer) {
	c.w = w
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c countingWrite
	c.Set(w)
	return &c, &c.cnt
}

func main() {

	buf := &bytes.Buffer{}

	w, cnt := CountingWriter(buf)

	str := "test!"
	w.Write([]byte(str))

	fmt.Printf("word = %s : cnt = %d", buf.String(), *cnt)

}
