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

/*
public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];
        if (map.containsKey(complement)) {
            return new int[] { map.get(complement), i };
        }
        map.put(nums[i], i);
    }
    throw new IllegalArgumentException("No two sum solution");
}
*/
