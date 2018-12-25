package lc473

import (
	"sort"
)

func makesquare(nums []int) bool {
	if len(nums) <= 0 {
		return false
	}
	sum := intSum(nums)
	if sum%4 != 0 {
		return false
	}
	length := sum / 4
	// sort
	sort.Ints(nums)
	if nums[0] > length {
		return false
	}
	result := make([]int, 4)
	for i := 0; i < len(result); i++ {

	}
	return true
}

func makesquare1(nums []int) bool {
	if len(nums) <= 0 {
		return false
	}
	sum := intSum(nums)
	if sum%4 != 0 {
		return false
	}
	length := sum / 4
	// sort
	sort.Ints(nums)
	if nums[0] > length {
		return false
	}
	result := make([]int, 4)
	minIndex := 0

	for i := len(nums) - 1; i >= 0; i-- {
		result[minIndex] += nums[i]
		if result[minIndex] > length {
			return false
		}
		minIndex = findMinIndex(result)
	}
	print(result)

	return true
}

func intSum(array []int) int {
	sum := 0
	for _, val := range array {
		sum += val
	}
	return sum
}

func findMinIndex(arr []int) int {

	minIndex := 0
	for i, val := range arr {
		if val < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func fillLateral(arr []int, target int) {
	tmpArray := arr[:]
	sum := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i]+tmp < target {
			tmp = tmp + arr[i]
		}
	}
}

func deleteIndexAt(arr []int, index int) []int {
	if index >= len(arr) || index < 0 {
		return nil
	}

	copy(arr[index:], arr[index+1:])
	arr = arr[:len(arr)-1]
	return arr
}
