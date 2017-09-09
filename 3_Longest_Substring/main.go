package main

import (
	"fmt"
)

func main() {
	fmt.Println(3 == lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(1 == lengthOfLongestSubstring("bbbb"))
	fmt.Println(1 == lengthOfLongestSubstring("b"))
	fmt.Println(1 == lengthOfLongestSubstring("bbbbb"))
	fmt.Println(2 == lengthOfLongestSubstring("aab"))
	fmt.Println(3 == lengthOfLongestSubstring("dvdf"))

}

func lengthOfLongestSubstring(s string) int {
	a := []byte(s)
	m := make([]bool, 255)
	l, res := 0, 0
	for h, curr := range a {
		if m[curr] {
			for a[l] != curr {
				m[a[l]] = false
				l++
			}
			l++
		} else {
			m[curr] = true
		}
		if h-l+1 > res {
			res = h - l + 1
		}
	}
	return res
}
