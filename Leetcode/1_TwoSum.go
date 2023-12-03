package Leetcode

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for currentPosition, currentValue := range nums {
		if requiredIndex, isPresent := sumMap[target-currentValue]; isPresent {
			return []int{requiredIndex, currentPosition}
		}
		sumMap[currentValue] = currentPosition
	}
	return []int{}
}
