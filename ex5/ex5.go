package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Sp0ozy/bst"
)

func MinIK(uzd []int, m int) int {
	var win *bst.Node
	minIk := 1
	ik := minIk

	for len(uzd) > 0 || win != nil {
		for win.NodeCount() < m && len(uzd) > 0 {
			win = win.Insert(uzd[0])
			uzd = uzd[1:]
		}
		change := false
		for i := m; i > 0 && win != nil; i-- {
			min := win.FindMin()
			// fmt.Println(min, ik, "             ", days)
			if min <= ik {
				win = win.DeleteNode(min)
				ik++
				change = true
				continue
			}
			if !change {
				minIk += (min - ik)
				ik += (min - ik)
			}
			break
		}
	}
	return minIk
}
func MinDays(uzd []int, m, ik int) {
	var win *bst.Node
	ikk := ik
	days := 0
	for len(uzd) > 0 || win != nil {
		days++
		for win.NodeCount() < m && len(uzd) > 0 {
			win = win.Insert(uzd[0])
			uzd = uzd[1:]
		}
		// fmt.Println(win.Left.Val, win.Val)
		for i := m; i > 0 && win != nil; i-- {
			min := win.FindMin()
			// fmt.Println(min, ik, "             ", days)
			if min <= ik {
				win = win.DeleteNode(min)
				ik++
				continue
			}
			break
		}
	}

	fmt.Println(ikk, days)
}
func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	uzd := make([]int, n)
	for i := range uzd {
		fmt.Fscan(r, &uzd[i])
	}
	// fmt.Println(uzd)
	ik := MinIK(uzd, m)
	// fmt.Println(ik)
	MinDays(uzd, m, ik)
}
