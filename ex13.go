package main

// https://leetcode.com/problems/reverse-nodes-in-k-group/submissions/880939484/

// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
//
// You may not alter the values in the list's nodes, only nodes themselves may be changed.

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	thisGroupCount := 1
	nextGroupHead := head.Next
	var thisGroupTail *ListNode
	for nextGroupHead != nil && thisGroupCount < k {
		thisGroupTail = nextGroupHead
		nextGroupHead = nextGroupHead.Next
		thisGroupCount++
	}

	if thisGroupCount < k {
		return head
	}

	h, t := reverse(head, thisGroupTail)
	t.Next = reverseKGroup(nextGroupHead, k)
	return h

}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	if tail == nil {
		head.Next = nil
		return head, head
	}

	tail.Next = nil

	ptr1 := head
	ptr2 := head.Next
	for ptr2 != nil {
		oldPtr := ptr1
		ptr1 = ptr2
		ptr2 = ptr2.Next

		ptr1.Next = oldPtr
	}

	head.Next = nil

	head, tail = tail, head
	return head, tail
}
