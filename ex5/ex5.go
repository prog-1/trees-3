package main

import (
	"bufio"
	"fmt"
	"os"
)

func solutions(ik int, lowlvlex []int) bool {
	for len(lowlvlex) > 0 {
		if lowlvlex[0] > ik {
			return false
		}
		// lowlvlex = append(lowlvlex[:0], lowlvlex[1:]...)
		lowlvlex = lowlvlex[1:]
		ik++
	}
	return true
}

func find(experday int, ex []int) ([]int, []int) {
	var lowlvlex []int
	for a := 0; a < experday && len(ex) > 0; a++ {
		min := ex[0]
		mini := 0
		for i, j := range ex {
			if min > j {
				min = j
				mini = i
			}
		}
		lowlvlex = append(lowlvlex, min)
		ex = append(ex[:mini], ex[mini+1:]...)
		fmt.Println(lowlvlex, ex)
	}
	return ex, lowlvlex
}

func findminIK(experday int, ex []int) int {
	ik := 1
	excopy := ex
	for {
		var lowlvlex []int
		ex, lowlvlex = find(experday, ex)
		if !solutions(ik, lowlvlex) {
			ik++
			ex = excopy
		} else if len(ex) == 0 {
			return ik
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, experday, dayneed int
	var ex []int
	fmt.Println("First two")
	fmt.Scan(&n, &experday)
	fmt.Println("Enter ex")
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		ex = append(ex, x)
	}
	if n/experday == 0 {
		dayneed = 1
	} else {
		dayneed = (n / experday) + (n % experday)
	}
	fmt.Println(findminIK(experday, ex), dayneed)
}

// 3 7 1 2 6 3
// fmt.Println("prst1")
