package main

import (
	"fmt"
)

func main() {
	//fmt.Println(IngusCoefficient(processInput()))
	fmt.Println(IngusCoefficient(3, []int{3, 1, 7, 2, 6, 3}))
}

type node[T any] struct {
	val         T
	count       int
	left, right *node[T]
}

func insert[T int](root **node[T], in T) {
	if (*root) == nil {
		*root = &node[T]{in, 1, nil, nil}
	} else if in == (*root).val {
		(*root).count++
	} else if in < (*root).val {
		insert(&(*root).left, in)
	} else /*in > root.val*/ {
		insert(&(*root).right, in)
	}
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

func (root *node[T]) lowest() *node[T] {
	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

func IngusCoefficient(m int, compl []int) (minIC, days int) {
	// Returns index of the smallest element
	smallest := func(i, j int) int {
		l := i
		for k := i + 1; k < j; k++ {
			if compl[k] < compl[l] {
				l = k
			}
		}
		return l
	}
	// Calculate minimal Ingus coeffecient
	// Take first m elements
	// Find lowest
	// minIK = lowest
	// if lowest <= cur then delete lowest and cur++ break
	// else /*lowest > cur*/ then minIK = minIK + (lowest-cur) and cur = lowest + 1
	// Delete lowest
	// take one more element
	tasks := compl
	curIC := minIC
	for i, j := 0, m; i < len(tasks); i, j = i+1, j+1 {
		if j > len(tasks) {
			j = len(tasks)
		}
		l := smallest(i, j)
		if tasks[l] > curIC {
			minIC = minIC + (tasks[l] - curIC)
			curIC = tasks[l]
		}
		curIC++
		tasks[i], tasks[l] = tasks[l], tasks[i]

	}

	return minIC, 0
}
