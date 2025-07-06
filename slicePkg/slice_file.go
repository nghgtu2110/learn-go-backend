package slicePkg

func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func FindMin(nums []int) int {
	min := nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min
}

func FindMax(nums []int) int {
	max := nums[0]

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}

func Avg(nums []int) int {
	sum := Sum(nums)
	avg := sum / len(nums)
	return avg
}

func Sort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				// Swap nums[j] and nums[j+1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
