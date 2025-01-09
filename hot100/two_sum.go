package hot100

import "fmt"

func SolutionTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		index, exist := numMap[complement]
		if exist {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}
