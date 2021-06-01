package src

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	low, high := 0, len(nums)-1

	middle := 0
	for low <= high {
		middle = (high-low)/2 + low
		if nums[middle] == target {
			break
		} else if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	if nums[middle] != target {
		return []int{-1, -1}
	}

	low, high = middle-1, middle+1
	for low >= 0 && nums[low] == target {
		low--
	}

	for high < len(nums) && nums[high] == target {
		high++
	}
	return []int{
		low + 1, high - 1,
	}
}
