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
	/**
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("]"))
	fmt.Println(isValid("["))
	fmt.Println(isValid(""))
	*/

	// 21. Merge Two Sorted Lists
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}
	h := mergeTwoLists(l1, l2)
	for {
		fmt.Println(h.Val)
		if h.Next == nil {
			break
		}
		h = h.Next
	}

	l1 = &ListNode{}
	l2 = &ListNode{}
	h = mergeTwoLists(l1, l2)
	for {
		fmt.Println(h.Val)
		if h.Next == nil {
			break
		}
		h = h.Next
	}
}
