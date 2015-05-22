package ProjectEular

import "strconv"

func IsItPalindrome(param int) bool {
	normal := strconv.Itoa(param)
	reverse := reverse(normal)
	if normal == reverse {
		return true
	}
	return false
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}