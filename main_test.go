package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func tree1() *node {
	return nil
}

func tree2() *node {
	return &node{val: 1}
}

func tree3() *node {
	return &node{
		val: 1,
		left: &node{
			val: -5,
		},
		right: &node{
			val: 3,
		},
	}
}
func tree4() *node {
	return &node{
		val: 1,
		left: &node{
			val: 2,
		},
		right: &node{
			val: 1,
		},
	}
}
func tree5() *node {
	return &node{
		val: 1,
		left: &node{
			val: 2,
		},
		right: &node{
			val: -8,
			left: &node{
				right: &node{val: 9},
			},
			right: &node{val: 4},
		},
	}
}

func TestDuplicates(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want bool
	}{
		{"1", tree1(), false},
		{"2", tree2(), false},
		{"3", tree3(), false},
		{"4", tree4(), true},
		{"5", tree5(), true},
	} {

		if got := Duplicates(tc.root); got != tc.want {
			var buf bytes.Buffer
			printTree(tc.root, &buf)
			t.Errorf("levelOrderTraversal(#%v) = %v, want %v\n%v", tc.name, got, tc.want, buf.String())
		}
	}
}

func TestLevelMax(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{"1", tree1(), []int{}},
		{"2", tree2(), []int{1}},
		{"3", tree3(), []int{1, 3}},
		{"4", tree4(), []int{1, 2}},
		{"5", tree5(), []int{1, 2, 4, 9}},
	} {

		if got := levelMax(tc.root); !reflect.DeepEqual(got, tc.want) {
			var buf bytes.Buffer
			printTree(tc.root, &buf)
			t.Errorf("levelOrderTraversal(#%v) = %v, want %v\n%v", tc.name, got, tc.want, buf.String())
		}
	}
}

func printTree(n *node, w io.Writer) {
	printer(n, w, nil)
}

func printNonemptyHistory(prev []int, w io.Writer, printLast bool) {
	if len(prev) == 0 {
		panic("must not happen")
	}
	last := len(prev)
	if !printLast {
		last--
	}
	for _, p := range prev[:last] {
		if p > 0 {
			fmt.Fprint(w, "| ")
		} else {
			fmt.Fprint(w, "  ")
		}
	}
}

func directChildren(n *node) (num int) {
	if n == nil {
		return 0
	}
	if n.left != nil {
		num++
	}
	if n.right != nil {
		num++
	}
	return num
}

func printer(n *node, w io.Writer, prev []int) {
	if n == nil {
		return
	}
	if len(prev) > 0 {
		printNonemptyHistory(prev, w, true /*printLast*/)
		fmt.Fprintln(w)
		printNonemptyHistory(prev, w, false /*printLast*/)
		prev[len(prev)-1]--
	}
	fmt.Fprintln(w, "+->", n.val)
	prev = append(prev, directChildren(n))
	if n.left != nil {
		printer(n.left, w, prev)
	}
	if n.right != nil {
		printer(n.right, w, prev)
	}
}
