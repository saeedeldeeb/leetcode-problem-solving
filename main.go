package main

import "fmt"

func main() {
	// 1. Two Sum
	/**
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
	fmt.Println(twoSum(nums, target))
	*/

	// 9. Palindrome Number
	/**
	fmt.Println(isPalindrome(-121))
	*/

	//13. Roman to Integer
	/**
	fmt.Println(romanToInt("MCMXCIV"))
	*/

	// 14. Longest Common Prefix
	/**
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
	*/

	// 20. Valid Parentheses
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("]"))
	fmt.Println(isValid("["))
	fmt.Println(isValid(""))
}
