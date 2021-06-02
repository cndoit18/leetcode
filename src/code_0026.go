package src

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	low, high := 0, 0
	for high < len(nums) {
		if nums[high] == nums[low] {
			high++
			continue
		}
		low++
		if high != low {
			nums[low] = nums[high]
		}
	}
	return low + 1
}
