package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, res, next *ListNode
	next = head.Next
	res = next

	for head != nil && next != nil {

		if prev != nil {
			prev.Next = next
		}
		head.Next = next.Next
		next.Next = head

		prev = head
		head = head.Next
		if head != nil {
			next = head.Next
		} else {
			head = nil
		}
	}

	return res
}

func main() {
	head := &ListNode{}
	head.Val = 1
	head.Next = head
	head.Next.Val = 2
	head.Next.Next = head
	head.Next.Next.Val = 3
	head.Next.Next.Next = head
	head.Next.Next.Next.Val = 4
	swapPairs(head)

}
