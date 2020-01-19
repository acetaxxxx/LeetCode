/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	iter := result
	for !(l1 == nil || l2 == nil) {

		if l1.Val < l2.Val {
			iter.Val = l1.Val
			iter.Next = new(ListNode)
			iter = iter.Next
			l1 = l1.Next

		} else {
			iter.Val = l2.Val
			iter.Next = new(ListNode)
			iter = iter.Next
			l2 = l2.Next
		}

	}
	if l1 == nil {
		iter.Val = l2.Val
		iter.Next = l2.Next
	}
	if l2 == nil {
		iter.Val = l1.Val
		iter.Next = l1.Next
	}
	return result
}

