package src

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	low, high := 0, len(nums)-1
	middle := 0
	for low <= high {
		middle = (high-low)/2 + low
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			high = middle - 1
		} else if nums[middle] < target {
			low = middle + 1
		}
	}
	if middle-low < 0 {
		return middle + 1
	}
	return middle
}
