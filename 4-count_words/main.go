package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type node struct {
	word string
	//cnt         int
	left, right *node
}

func insert(t *node, a string) *node {
	if t == nil {
		return &node{word: a}
	}
	if a < t.word {
		t.left = insert(t.left, a)

	} else {
		t.right = insert(t.right, a)
	}
	return t
}

func CreateBST(s []string) *node {
	var t *node
	for _, i := range s {
		t = insert(t, i)
	}
	return t
}

func inorder(t, prev *node, f func(string, int), cnt int) int {
	if t == nil {
		return 1
	}
	cnt = inorder(t.left, t, f, cnt)
	fmt.Println("t:", t.word)
	fmt.Println("prev:", prev.word)
	if t.word != prev.word {
		f(t.word, cnt)
		cnt = 1
	} else {
		cnt++
	}
	return inorder(t.right, t, f, cnt)
}

func main() {
	s, err := os.ReadFile("4-count_words/text_file.txt")
	if err != nil {
		panic(err)
	}
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(string(s), f)
	fmt.Println(words)

	t := CreateBST(words)
	c, err := os.Create("4-count_words/words_and_count_file.txt")
	if err != nil {
		panic(err)
	}
	inorder(t, &node{}, func(v string, i int) { c.WriteString(fmt.Sprintf("%v : %v\n", v, i)) }, 1)
}
