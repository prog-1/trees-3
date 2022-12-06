package main

import "fmt"

//iK - Ingus koeficient

//n - total task amount
//m - task amount that Ingus check every day

// Ingus is taking m task in day
// Ingus is trying to solve then in any order
// With every solved task, IK raises on 1
// If some task is too hard, he keeps it in next day
// if we can't solve all tasks of the day, we end

//Result:
//which starter IK should be to solve all tasks
//How many days is needed to solve all tasks with selected starter IK

//Input
// N M  // N - total task amount  M - tasks per day
// 3 1 7 2 6 3 // each task difficulty level

//Output
// sIK D
// sIK - starter IK | D - day count


func main() {
	//Works at least somehow
	book := []int{3, 1, 7, 2, 6, 3}
	TperD := 3
	//Don't work at all
	// book := []int{7, 6, 5, 4, 3}
	// TperD := 9
	fmt.Println(Ingus(TperD, book))
}

//minTI - min Task Level Index
//tA - task Amount

func Ingus(TperD int, book []int) (int, int) {

	//Variable declaration
	//var solveMinTask func() //pre declaration to use before implementation
	//var nextDay func()
	var days int           // days count
	if TperD > len(book) { //if task per day amount is bigger that task amount in book
		TperD = len(book) //make it equal to book task amount
	}
	dayTasks := book[:TperD] //taking first TperD elements
	//var dayTasks []int
	book = book[TperD:] //deleting first TperD elements from book
	//minTI := takeMin(dayTasks) //taking min task index (index to be able to delete it)
	var minTI int

	startIK := dayTasks[takeMin(dayTasks)] // making that iK is min task lvl, because it's necessary starting point
	//iK := startIK
	//dayTasks = append(dayTasks[:minTI], dayTasks[minTI+1:]...) //delete 1st min task (mb swap with 1st and delete 1st element?)

	for len(dayTasks) != 0 {
		b := book
		iK := startIK
		days = 0
		for len(b) != 0 { //while we haven't found the ik which will solve all tasks

			//next day
			tA := TperD - len(dayTasks) //calculate how much tasks we solved and need to take
			if tA > len(b) {            //if we want to take more task then is in book
				tA = len(b) //take the amount of task that left in book
			}
			dayTasks = append(dayTasks, b[:tA]...) //add new tasks in day tasks
			b = b[tA:]                             //delete taken tasks from book

			for range dayTasks {
				minTI = takeMin(dayTasks)
				if dayTasks[minTI] <= iK { // if we can solve task
					dayTasks = append(dayTasks[:minTI], dayTasks[minTI+1:]...) //delete this task (mb swap with 1st and delete 1st element?)
					iK++                                                       // increase iK
				} else {
					break //else we can't solve any more tasks today
				}
			}

			// //if we don't have any changes since previous day
			// if prevDay == len(dayTasks){
			// 	break
			// }
			// prevDay = len(dayTasks)//updating day progress
			days++ //increasing day count
		}
		startIK++ //try one more time with koefficient +1
		dayTasks = nil
	}
	return startIK, days

	//we need to repeat solving many times
	//how to act when we stopped? When nothing is changing day after day

}

//return index
func takeMin(s []int) int {
	res := 0
	for i := range s {
		if s[i] < s[res] {
			res = i
		}
	}
	return res
}
