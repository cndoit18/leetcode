package src

func mySqrt(x int) int {
	low, high := 0, x
	middle := high
	for low <= high {
		middle = (high-low)/2 + low
		sum := middle * middle
		if x == sum {
			return middle
		} else if sum > x {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	if low > middle {
		return middle
	}
	return middle - 1
}
