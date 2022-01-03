package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func (t *tree) String() string {
	var s string
	if t == nil {
		return s
	}
	s += t.left.String()
	s += strconv.Itoa(t.value)
	s += t.right.String()
	return s
}
func main() {
	root := &tree{0, nil, nil}
	add(root, 2)
	fmt.Println(root.String())
}
