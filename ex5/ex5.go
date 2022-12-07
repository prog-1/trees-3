package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Sp0ozy/bst"
)

func MinIK(uzd []int, m int) {
	var win *bst.Node
	for i := 0; i < m && len(uzd) != 0; i++ {
		win = win.Insert(uzd[0])
		uzd = uzd[1:]
	}
	minIk := 1
	ik := minIk
	for len(uzd) > 0 || win != nil {
		min := win.FindMin()
		if min <= ik {
			win = win.DeleteNode(min)
			ik++

			if len(uzd) > 0 {
				win = win.Insert(uzd[0])
				uzd = uzd[1:]
			}

			continue
		}
		minIk += (min - ik)
		ik += (min - ik)

	}
	fmt.Println(minIk)
}
func MinDays(uzd []int, m, ik int) {
	var win *bst.Node
	for i := 0; i < m && len(uzd) != 0; i++ {
		win = win.Insert(uzd[0])
		uzd = uzd[1:]
	}
	for days:=1;len(uzd) > 0 || win != nil; {
		for win != nil{
			min := win.FindMin()
		if min <= ik {
			win = win.DeleteNode(min)
			ik++
		}
		}
		
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

}
