package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// head *ListNode: is a pointer to a struct -> used in param of function
// var v ListNode: v has struct type
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	fmt.Println(slow)
	return slow
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	run := head
	length := 0
	for run != nil {
		run = run.Next
		length += 1
	}
	idx, remove_idx := 0, length-n-1
	if remove_idx == -1 {
		return head.Next
	}
	run = head
	for idx != remove_idx && run != nil {
		run = run.Next
		idx += 1
	}
	if run.Next != nil {
		run.Next = run.Next.Next
	} else {
		run.Next = nil
	}
	return head
}
