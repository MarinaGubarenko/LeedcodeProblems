package main

import (
	"fmt"

	pld "./palindrome"
)

func main() {
	var x int
	fmt.Scanf("%d\n", &x)
	fmt.Print(pld.IsPalindrome(x))

}
