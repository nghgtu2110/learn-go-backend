package two_sum

//https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	// key: number, value: index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := numMap[complement]; found {
			return []int{index, i}
		}

		// Save the current number and its index to map
		numMap[num] = i
	}

	return nil // if no solution found
}
