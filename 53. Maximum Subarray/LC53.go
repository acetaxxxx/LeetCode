package lc53

import (
	"math"
)

func maxSubArray(nums []int) int {
	length := len(nums)
	dpmap := make([]int, length)
	result := math.MinInt32
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j < i && i-1 >= 0 {
				dpmap[j] = dpmap[j] + nums[i]
			} else if i == j {
				dpmap[j] = nums[i]
			} else {
				dpmap[j] = math.MinInt32
			}
			if dpmap[j] > result {
				result = dpmap[j]
			}
		}
	}

	return result
}

func maxSubArray1(nums []int) int {
	length := len(nums)
	dpmap := make([][]int, length)
	result := math.MinInt32
	for i := 0; i < length; i++ {
		dpmap[i] = make([]int, length)
		for j := 0; j < length; j++ {
			if j < i && i-1 >= 0 {
				dpmap[i][j] = dpmap[i-1][j] + nums[i]
			} else if i == j {
				dpmap[i][j] = nums[i]
			} else {
				dpmap[i][j] = math.MinInt32
			}
			if dpmap[i][j] > result {
				result = dpmap[i][j]
			}
		}
	}

	return result
}
