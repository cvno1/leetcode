/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
package sollution

import "github.com/cvno1/leetcode/structure"

type ListNode = structure.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preslow, slow, fast := dummyHead, head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		preslow = slow
		slow = slow.Next
		fast = fast.Next
	}

	preslow.Next = slow.Next

	return dummyHead.Next
}

// @lc code=end
