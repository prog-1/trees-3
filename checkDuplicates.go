package main

//function that verifies whether there are duplicate elements in a binary tree
func CheckDuplicates(t *Node) bool {
	var s []int
	var checkDuplicates func(t *Node, s *[]int) bool
	checkDuplicates = func(t *Node, s *[]int) bool {
		if t == nil {
			return false
		}
		for _, el := range *s {
			if el == t.val {
				return true
			}
		}
		*s = append(*s, t.val)
		return checkDuplicates(t.left, s) || checkDuplicates(t.right, s)
	}
	return checkDuplicates(t, &s)

}

// func main() {
// 	tree := tree1()
// 	var s []int
// 	fmt.Println(checkDuplicates(tree, &s))
// }
