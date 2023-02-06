package main

// https://leetcode.com/problems/swap-nodes-in-pairs/

// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		first, second := head.Next, head
		first.Next, second.Next = second, swapPairs(first.Next)
		return first
	} else if head != nil && head.Next == nil {
		return head
	} else {
		return nil
	}
}
