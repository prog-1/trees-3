package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	left, right *node
	val         int
}

func (t *node) min() int {
	if t == nil {
		return -1
	}
	if t.left == nil {
		return t.val
	}
	return (t.left).min()
}

func removeMin(t **node) {
	i := (*t)
	if i.left == nil {
		*t = (*t).right
		return
	}
	for ; i.left.left != nil; i = i.left {
	}
	i.left = nil
}

func Insert(t **node, val int) {
	if *t == nil {
		*t = &node{val: val}
		return
	}
	if val < (*t).val {
		Insert(&(*t).left, val)
		return
	}
	if val > (*t).val {
		Insert(&(*t).right, val)
		return
	}
}

func main() {
	// INPUT
	var taskNum, tasksPerDay int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &taskNum, &tasksPerDay)
	tasks := make([]int, taskNum)
	for i := 0; i < taskNum; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)
		tasks[i] = tmp
	}
	if tasksPerDay > taskNum {
		tasksPerDay = taskNum
	}
	fmt.Println(findIngusCofficent(tasks, tasksPerDay))
}

func findIngusCofficent(tasks []int, tasksPerDay int) int {
	var ik, startIK, i int
	var BST *node
	for ; i < tasksPerDay; i++ {
		Insert(&BST, tasks[i])
	}
	for BST != nil {
		if ik < BST.min() {
			fmt.Println(BST.min(), ik, BST.min()-ik)
			startIK += BST.min() - ik
			ik += BST.min() - ik
		}
		removeMin(&BST)
		if i < len(tasks) {
			Insert(&BST, tasks[i])
			i++
		}
		ik++
	}
	return startIK
}
