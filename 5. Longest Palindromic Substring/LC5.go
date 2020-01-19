package lc5

func longestPalindrome(s string) string {

	for long := len(s); long >= 2; long-- {
		for shift := 0; shift <= len(s)-long; shift++ {
			t := []rune(s)
			f := string(t[shift : shift+long])
			if isPalindromic(f) {
				return f
			}
		}
	}
	t1 := []rune(s)
	return string(t1[0])
}
func isPalindromic(s string) bool {
	var tmp string
	ran := len(s) / 2
	if len(s)%2 == 1 {
		t := []rune(s)
		tmp = reverseString(string(t[ran+1:]))
	} else {
		t := []rune(s)
		tmp = reverseString(string(t[ran:]))
	}
	front := []rune(s)
	if string(front[:ran]) == tmp {
		return true
	}
	return false
}

func reverseString(s string) string {
	runes := []rune(s)
	result := make([]rune, len(runes))
	for i, j := len(runes)-1, 0; i >= 0; i, j = i-1, j+1 {
		result[j] = runes[i]
	}
	return string(result)
}
