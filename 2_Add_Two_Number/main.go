package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l00, l01, l02, l10, l11, l12 ListNode
	l00.Val = 5
	l01.Val = 4
	l02.Val = 3
	l10.Val = 5
	l11.Val = 6
	l12.Val = 4
	l00.Next = &l01
	l01.Next = &l02
	l02.Next = nil
	l10.Next = &l11
	l11.Next = &l12
	l12.Next = nil

	l00.Next = nil
	l10.Next = nil
	res := addTwoNumbers(&l00, &l10)
	//res := &l00
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next

	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	car := 0
	result := &ListNode{}
	resultPtr := result
	for l1 != nil || l2 != nil || car > 0 {
		sum := car
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		car = sum / 10
		resultPtr.Next = &ListNode{}
		resultPtr.Next.Val = sum % 10
		resultPtr = resultPtr.Next
	}
	return result.Next
}
