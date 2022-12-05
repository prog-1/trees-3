package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"unicode"
)

type node struct {
	left, right *node
	word        []byte
	repeates    int
}

func less(a, b []byte) bool {
	var invert bool
	if len(a) > len(b) {
		a, b = b, a
		invert = true
	}
	for i := range a {
		if a[i] < b[i] {
			return !invert
		}
		if a[i] > b[i] {
			return invert
		}
	}
	return true
}

func insert(t *node, v []byte) *node {
	if t == nil {
		t = &node{word: v, repeates: 1}
		return t
	}
	if reflect.DeepEqual(t.word, v) {
		t.repeates++
		return t
	}
	if less(v, t.word) {
		t.left = insert(t.left, v)
		return t
	}
	t.right = insert(t.right, v)
	return t
}

func createBSTFromString(t []byte) (BST *node) {
	t = bytes.ToLower(t)
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := bytes.FieldsFunc(t, f)
	for _, v := range words {
		BST = insert(BST, v)
	}
	return
}

func main() {
	file, err := os.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}
	BST := createBSTFromString(file)
	resultFile, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	inorderTraversal(BST, resultFile)
}

func inorderTraversal(t *node, file *os.File) {
	if t == nil {
		return
	}
	inorderTraversal(t.left, file)
	file.WriteString(fmt.Sprintf("%s : %v\n", t.word, t.repeates))
	inorderTraversal(t.right, file)
}
