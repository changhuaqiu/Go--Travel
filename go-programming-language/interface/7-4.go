package main

//破题
//题目告诉你需要思考的是重新实现Reader类型
//
import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type StringReader string

func (s *StringReader) Read(p []byte) (n int, err error) {
	copy(p, *s)
	return len(*s), io.EOF
}

func NewReader(s string) io.Reader {
	t := StringReader(s)

	return &t
}

func main() {
	doc, err := html.Parse(NewReader("<a herf='/value'>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse err: %v", err)
		os.Exit(1)
	}
	fmt.Println(doc.Data)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
