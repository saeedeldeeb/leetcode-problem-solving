package main

func removeDuplicates(nums []int) int {
	f := 0
	for i := 1; i < len(nums); {
		if nums[f] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
		f = i - 1
	}
	return len(nums)
}
