package LC946

func validateStackSequences(pushed []int, popped []int) bool {
	length := len(pushed)
	for _, val := range popped {
		t := indexOf(pushed, val)
		if t == -1 {
			return false
		}
		if length == 0 {
			return false
		}
		
	}
	return false
}

func indexOf(a []int, t int) int {

	for i, val := range a {
		if val == t {
			return i
		}
	}
	return -1
}
