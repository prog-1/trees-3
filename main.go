package main

import (
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {

	f, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	printTree(countWords(f))
}

type node[T any] struct {
	val         T
	count       int
	left, right *node[T]
}

func n[T any](v T, count int) *node[T]      { return &node[T]{v, count, nil, nil} }
func (root *node[T]) l(n *node[T]) *node[T] { root.left = n; return root }
func (root *node[T]) r(n *node[T]) *node[T] { root.right = n; return root }

func hasDublicates[T comparable](root *node[T]) bool {
	contains := func(s []T, val T) bool {
		for _, el := range s {
			if el == val {
				return true
			}
		}
		return false
	}
	var values []T
	var hd func(*node[T]) bool
	hd = func(r *node[T]) bool {
		if r == nil {
			return false
		}
		if contains(values, r.val) {
			return true
		}
		values = append(values, r.val)
		return hd(r.left) || hd(r.right)
	}
	return hd(root)
}

func levelMax(root *node[int]) []int {
	var values [][]int
	var traverse func(*node[int], int)
	traverse = func(root *node[int], lvl int) {
		if root == nil {
			return
		}
		if len(values) <= lvl {
			tmp := make([]int, 1)
			values = append(values, tmp)
		}
		values[lvl] = append(values[lvl], root.val)
		traverse(root.left, lvl+1)
		traverse(root.right, lvl+1)
	}
	traverse(root, 0)

	lm := make([]int, len(values))
	for i, lvl := range values {
		for _, val := range lvl {
			if lm[i] < val {
				lm[i] = val
			}
		}
	}

	return lm
}

func countWords(file []byte) *node[string] {
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(string(file), f)

	var root *node[string]
	for _, word := range words {
		insert(&root, word)
	}
	return root
}

func insert(root **node[string], in string) {
	if (*root) == nil {
		*root = &node[string]{in, 1, nil, nil}
	} else if in == (*root).val {
		(*root).count++
	} else if in < (*root).val {
		insert(&(*root).left, in)
	} else /*in > root.val*/ {
		insert(&(*root).right, in)
	}
}
