package main

import "fmt"

func partition(s []int, pivot int) (i int) {
	j := 0
	for ; j < len(s); j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	return i
}

func IngusKoef(m, ik int, ex []int) (int, int) {
	i := partition(ex, ik)
	return i, 0
}

func main() {
	// n := 6
	m := 3
	ex := []int{3, 1, 7, 2, 6, 3} // len n
	fmt.Println(IngusKoef(m, 1, ex))
}
