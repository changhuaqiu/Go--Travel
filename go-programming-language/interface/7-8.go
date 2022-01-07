package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byFunc func(i, j int) bool
type customSort struct {
	t        []*Track
	lessFunc []byFunc
	optInd   []int
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func (x customSort) Len() int { return len(x.t) }
func (x customSort) Less(i, j int) bool {
	for _, opt := range x.optInd {
		if x.lessFunc[opt](i, j) {
			return true
		} else if !x.lessFunc[opt](j, i) {
			continue
		} else {
			return false
		}
	}
	return false
}
func (x customSort) byTitle(i, j int) bool {
	return x.t[i].Title < x.t[j].Title
}
func (x customSort) byArtist(i, j int) bool {
	return x.t[i].Artist < x.t[j].Artist
}
func (x customSort) byAlbum(i, j int) bool {
	return x.t[i].Album < x.t[j].Album
}

func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (x customSort) UpdateSord(idx int) {
	//更新
	optIndSize := len(x.optInd)
	if idx >= optIndSize {
		return
	}
	val := x.optInd[idx]
	for j := idx - 1; j >= 0; j-- {
		x.optInd[j+1] = x.optInd[j]
	}
	x.optInd[0] = val
}
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
func main() {
	table := customSort{t: tracks}
	table.lessFunc = []byFunc{table.byTitle, table.byArtist, table.byAlbum}
	for idx, _ := range table.lessFunc {
		table.optInd = append(table.optInd, idx)
	}
	sort.Sort(table)
	printTracks(table.t)
	table.UpdateSord(1)
	sort.Sort(table)
	printTracks(table.t)
}
