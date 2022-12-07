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
	for t.left != nil {
		t = t.left
	}
	return t.val
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
	for *t != nil {
		if val <= (*t).val {
			t = &(*t).left
		} else {
			t = &(*t).right
		}
	}
	*t = &node{val: val}
	return
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
	IK := findIngusCofficent(tasks, tasksPerDay)
	fmt.Println(IK, calculateDays(IK, tasks, tasksPerDay))
}

func findIngusCofficent(tasks []int, tasksPerDay int) int {
	var ik, startIK, i int
	var BST *node
	for ; i < tasksPerDay; i++ {
		Insert(&BST, tasks[i])
	}
	for BST != nil {
		if d := BST.min() - ik; d > 0 {
			startIK += d
			ik += d
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

func calculateDays(cofficent int, tasks []int, tasksPerDay int) int {

	//Init
	var day, i int
	var BST *node
	var removed int
	for ; i < tasksPerDay; i++ {
		Insert(&BST, tasks[i])
	}
	// Solve
	for i < len(tasks) {
		if BST == nil || BST.min() > cofficent {
			day++
			addNewTasks(&i, tasks, &removed, &BST)
			continue
		}
		removeMin(&BST)
		cofficent++
		removed++
	}

	return day + 1
}

func addNewTasks(i *int, tasks []int, removed *int, BST **node) {
	for ; *removed != 0 && *i < len(tasks); *removed-- {
		Insert(BST, tasks[*i])
		*i++
	}
}
