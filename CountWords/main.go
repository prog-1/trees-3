package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type node struct {
	val         string
	count       int
	left, right *node
}

func countWords(root *node, s []string) *node {
	var count func(root *node, v string) *node
	count = func(root *node, v string) *node {
		if root == nil {
			return &node{val: v, count: 1}
		}
		if v < root.val {
			root.left = count(root.left, v)
		} else if v > root.val {
			root.right = count(root.right, v)
		} else {
			root.count++
		}
		return root
	}
	for _, v := range s {
		root = count(root, v)
	}
	return root
}

func inorderTraversal(root *node, sink func(v *node)) {
	if root == nil {
		return
	}
	inorderTraversal(root.left, sink)
	sink(root)
	inorderTraversal(root.right, sink)
}

func main() {
	file, err := os.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}
	fileNew := strings.ToLower(string(file))
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(fileNew, f)
	bst := countWords(nil, words)
	res, err := os.Create("res.txt")
	if err != nil {
		log.Fatal(err)
	}
	inorderTraversal(bst, func(v *node) { fmt.Fprintf(res, "%v: %v\n", v.val, v.count) })
}
