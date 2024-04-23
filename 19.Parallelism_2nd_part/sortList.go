package main

import (
	"fmt"
	"slices"
)

/*
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	curr := head
	a := []*ListNode{}

	for curr != nil {
		a = append(a, curr)
		curr = curr.Next
	}

	slices.SortFunc(a, func(l1, l2 *ListNode) int {
		return l1.Val - l2.Val
	})

	newNode := &ListNode{}
	curr = newNode

	for _, node := range a {
		curr.Next = node
		curr = node
	}

	curr.Next = nil

	return newNode.Next
}

func main() {
	headNums := []int{-1, 5, 3, 4, 0}
	head := &ListNode{}
	curr := head
	for _, headNum := range headNums {
		curr.Next = &ListNode{Val: headNum}
		curr = curr.Next
	}

	sorted := sortList(head)
	for sorted != nil {
		fmt.Print(sorted.Val)
		sorted = sorted.Next
	}
	fmt.Println()
}
