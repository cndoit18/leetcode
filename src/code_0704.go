package src

// code: https://leetcode-cn.com/problems/binary-search/
func search(nums []int, target int) int {
	left, right := len(nums)-1, 0

	for right <= left {
		middle := (left-right)/2 + right
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			left = middle - 1
		}
		if nums[middle] < target {
			right = middle + 1
		}
	}
	return -1
}
