package src

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	low, high := 0, 0
	for high < len(nums) {
		if nums[high] == val {
			high++
			continue
		}
		if nums[low] != nums[high] {
			nums[low] = nums[high]
		}
		low, high = low+1, high+1
	}
	return low
}
