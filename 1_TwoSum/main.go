package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 7, 11, 13}
	target := 14
	fmt.Printf("%v", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		pos, ok := m[target-num]
		if ok && pos != i {
			return []int{pos, i}
		}
		m[num] = i
	}
	return nil
}
