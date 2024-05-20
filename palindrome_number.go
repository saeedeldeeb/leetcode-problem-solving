package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return reverseInt(x) == x
}

func reverseInt(x int) int {
	var rev int
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return rev
}
