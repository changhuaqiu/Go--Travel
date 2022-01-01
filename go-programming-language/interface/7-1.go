package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func main() {
	var b ByteCounter
	var w WordCounter
	var l LineCounter

	b.Write([]byte("hello world"))
	fmt.Println(b)

	w.Write([]byte("hello world"))
	fmt.Println(w)

	l.Write([]byte("hello \n world \n hello \n"))
	fmt.Println(l)

}

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanWords)
	*w += WordCounter(count)
	return count, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanLines)
	*l += LineCounter(count)
	return count, nil
}
func retCount(p []byte, fn bufio.SplitFunc) (count int) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(fn)
	count = 0
	for scanner.Scan() {
		count++
	}
	if error := scanner.Err(); error != nil {
		fmt.Printf("Invalid input: %s", error)
	}
	return
}
