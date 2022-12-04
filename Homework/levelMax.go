package main

func LevelMax(t *Node) (resValues []int) {
	var levelMax func(t *Node, lvl int)
	levelMax = func(t *Node, lvl int) {
		if t == nil {
			return
		}

		if lvl == len(resValues) { // if we are on another level
			//append another value to result slice
			resValues = append(resValues, t.val)
		} else if t.val > resValues[lvl] {
			// if we are on the same level & current value is bigger than last stored
			//update last level element
			resValues[lvl] = t.val
		}

		//Traversal
		levelMax(t.left, lvl+1)
		levelMax(t.right, lvl+1)
	}
	levelMax(t, 0)
	return resValues
}

// func main() {
// 	tree := Tree2()
// 	fmt.Println(LevelMax(tree))
// }
