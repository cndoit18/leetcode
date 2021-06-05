package src

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	sum := 0
	result := 0
	for j < len(nums) {
		sum += nums[j]
		j++
		for sum >= target && i < len(nums) {
			if result > (j-i) || result == 0 {
				result = j - i
			}
			sum -= nums[i]
			i++
		}
	}
	return result
}
