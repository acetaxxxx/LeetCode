package LC11

import (
	"math"
)

func maxArea(height []int) int {
	length := len(height)

	result := math.MinInt32
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i < j {
				tmp := int(math.Min(float64(height[i]), float64(height[j])) * float64((j - i)))
				if tmp > result {
					result = tmp
				}
			}
		}
	}

	return result
}
