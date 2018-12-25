package lc20

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	t := []rune(s)
	println(s)
	for i := 1; i < len(t); i++ {
		if t[i] == 125 || t[i] == 93 || t[i] == 41 {
			if isMatch(t[i], t[i-1]) == false {
				return false
			}
			t = deleteIndexAt(t, i)
			t = deleteIndexAt(t, i-1)
			i -= 2
			if i < 0 {
				i = 0
			}
		}
	}
	s = string(t)
	println(s)
	return len(t) == 0
}
func isMatch(end rune, prev rune) bool {
	if end == 125 {
		return prev == 123
	} else if end == 93 {
		return prev == 91
	} else if end == 41 {
		return prev == 40
	}
	return false
}

// 123 {
// 125 }
// 91 [
// 93 ]
// 40 (
// 41 )
func deleteIndexAt(arr []rune, index int) []rune {
	if index >= len(arr) || index < 0 {
		return nil
	}

	copy(arr[index:], arr[index+1:])
	arr = arr[:len(arr)-1]
	return arr
}
