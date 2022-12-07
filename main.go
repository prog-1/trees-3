package main

import (
	"fmt"
)

func main() {
	//fmt.Println(IngusCoefficient(processInput()))
	fmt.Println(IngusCoefficient(3, []int{3, 1, 7, 2, 6, 3}))
	//fmt.Println(IngusCoefficient(9, []int{7, 6, 5, 4, 3}))
}

type node[T any] struct {
	val         T
	count       int
	left, right *node[T]
}

func insert[T int](t *node[T], val T) *node[T] {
	p := &t
	for *p != nil {
		if val < (*p).val {
			p = &(*p).left
		} else {
			p = &(*p).right
		}
	}
	*p = &node[T]{val: val}
	return t
}

func processInput() (m int, tasks []int) {
	var taskCount int
	fmt.Scanln(&taskCount, &m)
	tasks = make([]int, taskCount)
	for i := range tasks {
		fmt.Scan(&tasks[i])
	}
	return m, tasks
}

func (root *node[T]) smallest() *node[T] {
	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

func deleteNode(root *node[int], k int) *node[int] {
	if root == nil {
		return nil
	}
	p := &root
	for (*p) != nil && k < (*p).val {
		p = &(*p).left
	}
	for (*p) != nil && k > (*p).val {
		p = &(*p).right
	}
	if (*p) == nil {
		return nil
	}

	if (*p).left == nil && (*p).right == nil {
		(*p) = nil
		return root
	}
	if (*p).left != nil && (*p).right == nil {
		(*p) = (*p).left
		return root
	}
	if (*p).left == nil && (*p).right != nil {
		(*p) = (*p).right
		return root
	}

	sr := &(*p).right // sr - smallest right
	for (*sr).left != nil {
		sr = &(*sr).left
	}
	(*p).val = (*sr).val
	(*sr) = nil
	return root
}

func IngusCoefficient(m int, compl []int) (minIC, days int) {
	tasks := make([]int, len(compl))
	copy(tasks, compl)
	// Returns index of the smallest element
	smallest := func(i, j int) int {
		l := i
		for k := i + 1; k < j; k++ {
			if tasks[k] < tasks[l] {
				l = k
			}
		}
		return l
	}
	for i, j, curIC := 0, m, minIC; i < len(tasks); i, j, curIC = i+1, j+1, curIC+1 {
		if j > len(tasks) {
			j = len(tasks)
		}
		l := smallest(i, j)
		if tasks[l] > curIC {
			minIC = minIC + (tasks[l] - curIC)
			curIC = tasks[l]
		}
		tasks[i], tasks[l] = tasks[l], tasks[i]
	}

	// Calculating days:
	// Increase day count
	// Insert first m tasks into BST
	// Remove smallest one
	// Try to solve next smallest
	//

	var root *node[int]
	curIC := minIC
	var solved int
	for i, j := 0, m; i < len(compl); i, j = i+solved, j+solved {
		solved = 0
		days++
		if j > len(compl) {
			j = len(compl)
		}
		for k := i; k < j; k++ {
			root = insert(root, compl[k])
		}
		for root != nil && root.smallest().val <= curIC {
			root = deleteNode(root, root.smallest().val)
			curIC++
			solved++
		}
	}

	return minIC, days
}
