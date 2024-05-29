package main

func removeDuplicates(nums []int) int {
	f := nums[0]
	for i := 1; i < len(nums); {
		if f == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
		f = nums[i-1]
	}
	return len(nums)
}
