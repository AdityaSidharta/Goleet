package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for key1, num1 := range nums {
		mymap[target-num1] = key1
	}
	fmt.Println(mymap)
	for key2, num2 := range nums {
		key1, ok := mymap[num2]
		if ok {
			if key1 != key2 {
				return []int{key2, key1}
			}
		}
	}
	return []int{0, 1}
}
