package src

func isPerfectSquare(num int) bool {
	low, high := 0, num
	for low <= high {
		middle := (high-low)/2 + low
		sum := middle * middle
		if sum == num {
			return true
		}
		if sum > num {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return false
}
