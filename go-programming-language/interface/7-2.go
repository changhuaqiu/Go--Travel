package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	Writer io.Writer
	Count  int
}

func (c *CountWriter) Write(p []byte) (n int, err error) {
	n, err = c.Writer.Write(p)
	if err != nil {
		return
	}
	c.Count += n
	return
}

func main() {
	cw, counter := CountingWriter(os.Stdout)
	cw.Write([]byte("Append soething..."))
	fmt.Println(*counter)
}

func CountingWriter(w io.Writer) (io.Writer, *int) {
	c := CountWriter{
		Writer: w,
	}
	return &c, &(c.Count)
}
