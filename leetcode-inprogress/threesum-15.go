package leetcode_inprogress

import "fmt"

func twoSumDistinct(nums []int, target int) [][]int {
	result := make([][]int, 0)
	maps := make(map[int]int)
	for idx1, num1 := range nums{
		if _, ok := maps[target -num1]; !ok {
			maps[target -num1] = idx1
		}
	}
	for idx2, num2 := range nums{
		if num2 < target {
			if idx1, ok := maps[num2]; ok {
				if idx1 != idx2 {
					result = append(result, []int {nums[idx1], nums[idx2]})
					delete(maps, num2)
				}
			}
		} else if (num2 == target) && (target == 0) {
			if idx1, ok := maps[num2]; ok {
				if idx1 < idx2 {
					result = append(result, []int {nums[idx1], nums[idx2]})
					delete(maps, num2)
				}
			}
		}
	}
	fmt.Println("nums: ", nums)
	fmt.Println("target: ", target)
	fmt.Println("result: ", result)
	return result
}

func concat(frontNums []int, backNums []int) []int {
	result := make([]int, 0)
	for _, value := range frontNums{
		result = append(result, value)
	}
	for _, value := range backNums{
		result = append(result, value)
	}
	return result
}

// TODO : How to make threeSum Distinct?
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	visited := make(map[int]bool)
	for idx, num := range nums{
		if num <= 0 {
			if _, ok := visited[num]; !ok {
				visited[num] = true
				target := -num
				frontNums := nums[:idx]
				backNums := nums[idx + 1:]
				restNums := concat(frontNums, backNums)
				tmpResults := twoSumDistinct(restNums, target)
				for _, tmpResult := range tmpResults {
					tmpResult = append(tmpResult, num)
					result = append(result, tmpResult)
				}
			}
		}
	}
	fmt.Println(result)
	return result
}