package main

import (
	"bufio"
	"fmt"
	"os"
)

type uzdevums struct {
	dif  int
	done bool
}

func ingus(uzd []uzdevums, m int) {
	var cnt2 int
	var ikk int
	for ikk = 1; ; ikk++ {
		ik := ikk
		for i := range uzd {
			uzd[i].done = false
		}
		cnt2 = 0
		for {
			var changed = false
			for i, cnt := 0, 1; i < len(uzd) && cnt <= m; i++ {
				if !uzd[i].done {
					cnt++
					if uzd[i].dif <= ik {
						fmt.Println(ik)
						ik++

						fmt.Println(uzd)
						uzd[i].done = true
						changed = true
					}
				}

			}
			if !changed {
				break
			}
			cnt2++
		}
		var a bool
		for _, v := range uzd {
			if !v.done {
				a = true

			}
		}
		if a {
			continue
		}
		break

	}

	fmt.Println(ikk, cnt2)
}
func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	uzd := make([]uzdevums, n)
	for i := range uzd {
		fmt.Fscan(r, &uzd[i].dif)
	}
	ingus(uzd, m)

}
