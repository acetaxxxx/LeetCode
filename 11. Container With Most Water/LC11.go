package LC11

import (
	"math"
)

func maxArea(height []int) int {
	length := len(height)

	maxIndex := 0
	result := 0
	for t := length - 1; t > 0; t-- {
		var areaL, areaR int
		areaL = 0
		areaR = 0

		leftIndex := maxIndex - t
		if leftIndex >= 0 && leftIndex < length {
			areaL = minInt(height[maxIndex], height[leftIndex]) * t
			if height[maxIndex] < height[leftIndex] {
				maxIndex = leftIndex
			}
		}

		rightIndex := maxIndex + t
		if rightIndex >= 0 && rightIndex < length {
			areaR = minInt(height[maxIndex], height[rightIndex]) * t
			if height[maxIndex] < height[rightIndex] {
				maxIndex = rightIndex
			}
		}

		area := maxInt(areaL, areaR)
		result = maxInt(area, result)
	}

	return result
}

func minInt(i, j int) int {
	r := math.Min(float64(i), float64(j))
	return int(r)
}

func maxInt(i, j int) int {
	r := math.Max(float64(i), float64(j))
	return int(r)
}
