package main

import "fmt"

func main() {
	// fmt.Println(solve(processInput()))
	c := []int{3, 1, 7, 2, 6, 3}
	fmt.Println(solve(3, c))

	// c := []int{7, 6, 5, 4, 3}
	// fmt.Println(solve(9, c))

	//c := []int{5, 3, 1, 6, 3, 5, 3}
	//fmt.Println(solve(2, c))

}

func processInput() (maxDailyTasks int, complexities []int) {
	var taskCount int
	fmt.Scanln(&taskCount, &maxDailyTasks)
	complexities = make([]int, taskCount)
	for i := range complexities {
		fmt.Scan(&complexities[i])
	}
	return maxDailyTasks, complexities
}

// mit - maximal inspected tasks(a day)
// miic - minimal initial Ingus coefficient
func solve(mit int, complexities []int) (miic, days int) {
	lowest := func(s []int) int {
		lowest := s[0]
		for _, el := range s {
			if el < lowest {
				lowest = el
			}
		}
		return lowest
	}
	if mit > len(complexities) {
		mit = len(complexities)
	}
	miic = lowest(complexities[:mit])
	var ct []int // current tasks(task complexities)
	var c []int
	// Appends recquired amount of tasks from complexities into ct
	updateCT := func() {
		if len(ct) == mit {
			return
		}
		d := mit - len(ct)
		// if d > len(c) {
		// 	d = len(c)
		// }
		ct = append(ct, c[:d]...)
		c = c[d:]
	}

	for {
		c = complexities
		ic := miic
		for len(c) != 0 {
			updateCT()
			for i := 0; i < len(ct); {
				if ct[i] <= ic {
					ct = append(ct[:i], ct[i+1:]...)
					ic++
					i = 0
				} else {
					i++
				}
			}
			days++
		}
		if len(ct) == 0 && len(c) == 0 {
			return miic, days
		}
		miic++
		ct = nil
		days = 0
	}

}
