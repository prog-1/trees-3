package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
//Todo: add len of BST
// less recursion
// avl, red-black, random trees, heaps, push,pop
type node struct {
	v           int
	left, right *node
}

func ik(r io.Reader) (int, int) {
	var n, m, days, minIk int
	fmt.Fscan(r, &n, &m)
	allEx := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &allEx[i])
	}
	fmt.Println(n, m, allEx)
	var dailyEx *node
	ik := 1
	var currTask int
	tasksLeftForDay := m
	for nodeCount(dailyEx) < m {
		for ; currTask < m; currTask++ {
			insertIntoBST(dailyEx, allEx[currTask])
		}
	}
	if minIk <= minValued(dailyEx).v {
		minIk = minValued(dailyEx).v
	} else {
		ik++
		dailyEx = Delete(dailyEx, minValued(dailyEx).v)
		tasksLeftForDay--
	}
	for i := m - 1; i != n-1; i++ {
		for tasksLeftForDay != 0 {
			if nodeCount(dailyEx) < m {
				insertIntoBST(dailyEx, allEx[currTask])
			}
			if ik >= minValued(dailyEx).v {
				ik++
				dailyEx = Delete(dailyEx, minValued(dailyEx).v)
				tasksLeftForDay--
			}
			return minIk, 1

		}
		days++
		tasksLeftForDay = m
	}
	return ik, days
}

func main() {
	r := bufio.NewReader(os.Stdin)
	a, b := ik(r)
	fmt.Println(a, b)
}

/*
	for _, i := range allEx {
		if nodeCount(dailyEx) < m {
			insertIntoBST(dailyEx, allEx[i])
		}
		if nodeCount(dailyEx) == m {
			days++
			for dailyEx.left != nil {
				dailyEx = dailyEx.left
				if minValued(dailyEx).v <= ik {
					ik++
				}
				Delete(dailyEx, minValued(dailyEx).v)
			}
		}
	}

ik = minValued(dailyEx).v
days = 1
*/
func insertIntoBST(root *node, val int) *node {
	if root == nil {
		return &node{val, nil, nil}
	}
	if val < root.v {
		root.left = insertIntoBST(root.left, val)
	} else {
		root.right = insertIntoBST(root.right, val)
	}
	return root
}

func nodeCount(root *node) int {
	if root == nil {
		return 0
	}
	return nodeCount(root.left) + nodeCount(root.right) + 1
}

func minValued(root *node) *node {
	temp := root
	for nil != temp && temp.left != nil {
		temp = temp.left
	}
	return temp
}

func Delete(root *node, val int) *node {
	if nil == root {
		return root
	}
	if root.v > val {
		root.left = Delete(root.left, val)
	}
	if root.v < val {
		root.right = Delete(root.right, val)
	}
	if root.v == val {
		if root.left == nil && root.right == nil {
			root = nil
			return root
		}
		if root.left == nil && root.right != nil {
			temp := root.right
			root = nil
			root = temp
			return root
		}
		if root.left != nil && root.right == nil {
			temp := root.left
			root = nil
			root = temp
			return root
		}

		tempNode := minValued(root.right)
		root.v = tempNode.v
		root.right = Delete(root.right, tempNode.v)
	}
	return root
}
