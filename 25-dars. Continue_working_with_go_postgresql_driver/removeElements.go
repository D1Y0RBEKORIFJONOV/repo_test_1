package main

import "fmt"

/*
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.



Example 1:


Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
Example 2:

Input: head = [], val = 1
Output: []
Example 3:

Input: head = [7,7,7,7], val = 7
Output: []


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head != nil && head.Val == val {
		return head.Next
	}

	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	head = removeElements(head, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
