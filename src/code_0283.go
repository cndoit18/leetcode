package src

func moveZeroes(nums []int) {
	low, high := 0, 0
	for high < len(nums) {
		if nums[high] == 0 {
			high++
			continue
		}
		if nums[low] == 0 {
			nums[low], nums[high] = nums[high], nums[low]
		}
		high++
		low++
	}
}
