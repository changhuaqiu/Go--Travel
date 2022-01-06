package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type LimitedReader struct {
	underlyingReader io.Reader
	remainBytes      int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.remainBytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.remainBytes {
		p = p[:l.remainBytes]
	}
	n, err = l.underlyingReader.Read(p)
	l.remainBytes -= int64(len(p))
	return
}
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
func main() {
	lr := LimitReader(strings.NewReader("12345"), 1)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
	fmt.Printf("%s\n", b)
}
