package src

func sortedSquares(nums []int) []int {
	index := 0
	for index < len(nums) {
		if nums[index] >= 0 {
			break
		}
		index++
	}

	sub := make([]int, 0, index)
	sup := make([]int, 0, len(nums)-index)
	for i := index - 1; i >= 0; i-- {
		sub = append(sub, nums[i])
	}
	for i := index; i < len(nums); i++ {
		sup = append(sup, nums[i])
	}
	result := make([]int, 0, len(nums))
	for len(sub) != 0 || len(sup) != 0 {
		if len(sub) == 0 {
			result = append(result, sup...)
			sup = nil
			break
		}
		if len(sup) == 0 {
			result = append(result, sub...)
			sub = nil
			break
		}

		if abs(sub[0]) > abs(sup[0]) {
			result = append(result, sup[0])
			sup = sup[1:]
		} else {
			result = append(result, sub[0])
			sub = sub[1:]
		}
	}
	for i := range result {
		result[i] = result[i] * result[i]
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
