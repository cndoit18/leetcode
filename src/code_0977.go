package src

func sortedSquares(nums []int) []int {
	low, high := 0, len(nums)-1
	result := make([]int, len(nums))
	index := len(result) - 1
	for low <= high {
		ln, hn := nums[low]*nums[low], nums[high]*nums[high]
		if ln > hn {
			low++
			result[index] = ln
		} else {
			high--
			result[index] = hn
		}
		index--
	}
	return result
}
