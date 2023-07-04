package palindrome

import "strconv"

func IsPalindrome(x int) bool {
	arr := strconv.Itoa(x)
	var fl = true
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			fl = false
			break
		}
	}
	return fl
}
