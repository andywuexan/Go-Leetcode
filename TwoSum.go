package Go_Leetcode

func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int, len(nums))

	for index, value := range nums {
		if j, isPresent := tempMap[target-value]; isPresent {
			return []int{j, index}
		}
		tempMap[value] = index
	}

	return nil
}
