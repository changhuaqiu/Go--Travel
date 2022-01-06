package main

import "time"

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}
type customSort struct {
	t    []*Track
	l    *list
	less func(x, y *Track) bool
}

func main() {

}

func (x customSort) Len() int      { return len(x.t) }
func (x customSort) Less(i, j int) { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
