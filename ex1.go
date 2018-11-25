/**
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list. You may assume the two numbers do
 * not contain any leading zero, except the number 0 itself.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return add(l1, l2, 0)
}

func add(l1, l2 *ListNode, r int) *ListNode {
    if l1 == nil && l2 == nil {
        if r != 0 {
            return &ListNode{Val: r}
        } else {
            return nil
        }
    }
    
    var l1Val, l2Val int
    var l1Next, l2Next *ListNode 
    
    if l1 != nil {
        l1Val = l1.Val
        l1Next = l1.Next
    }
    
    if l2 != nil {
        l2Val = l2.Val
        l2Next = l2.Next
    }
    
    v := l1Val + l2Val + r
    return &ListNode{
        Val: v % 10,
        Next: add(l1Next, l2Next, v / 10),
    }
}