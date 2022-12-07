package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solutions(ik, experday int, lowlvlex, ex []int) (bool, []int, int) {
	// fmt.Println(ik, lowlvlex, ex)
	for len(lowlvlex) > 0 {
		if lowlvlex[0] > ik && len(lowlvlex) == experday {
			return false, lowlvlex, ik
		} else if lowlvlex[0] > ik && len(ex) == 0 {
			return false, lowlvlex, ik
		} else if lowlvlex[0] > ik {
			return true, lowlvlex, ik
		}
		lowlvlex = lowlvlex[1:]
		ik++
	}
	return true, lowlvlex, ik
}

func find(experday int, ex, lowlvlex []int) ([]int, []int) {
	if len(lowlvlex) == 0 && len(ex) >= experday {
		lowlvlex = ex[:experday]
		ex = ex[experday:]
	}
	for len(lowlvlex) < experday && len(ex) > 0 {
		lowlvlex = append(lowlvlex, ex[0])
		ex = ex[1:]
	}
	sort.Ints(lowlvlex)
	return ex, lowlvlex
}

func copy(ex, excopy []int) []int {
	ex = nil
	ex = append(ex, excopy...)
	return ex
}

func findminIK(experday int, ex, excopy []int) (int, int) {
	ik, dayneed, ikcopy := 1, 0, 1
	var lowlvlex []int
	var pass bool
	for {
		// fmt.Println(ex)
		// fmt.Println(excopy)
		ex, lowlvlex = find(experday, ex, lowlvlex)
		pass, lowlvlex, ik = solutions(ik, experday, lowlvlex, ex)
		dayneed++
		if !pass {
			ikcopy++
			ik = ikcopy
			ex = copy(ex, excopy)
			lowlvlex = nil
			dayneed = 0
		} else if len(ex) == 0 {
			return ikcopy, dayneed
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, experday int
	var ex, excopy []int
	fmt.Println("First two")
	fmt.Scan(&n, &experday)
	fmt.Println("Enter ex")
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		ex = append(ex, x)
		excopy = append(excopy, x)
	}
	fmt.Println(findminIK(experday, ex, excopy))
}

// var ex []int
// excopy := ex
// lowlvlex := ex[:3]
// sort.Ints(lowlvlex)
// excopy[:3] <- this part of slice is sorted similar to lowlvlex
