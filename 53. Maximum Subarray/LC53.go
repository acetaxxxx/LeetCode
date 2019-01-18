package lc53

func maxSubArray(nums []int) int {
	result := nums[0]
	sum := 0
	for _, val := range nums {
		if sum > 0 {
			sum += val
		} else {
			sum = val
		}
		if sum > result {
			result = sum
		}
	}
	return result
}
