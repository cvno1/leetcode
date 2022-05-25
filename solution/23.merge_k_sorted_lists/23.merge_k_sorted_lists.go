/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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
func mergeKLists(lists []*ListNode) *ListNode {
	lenth := len(lists)
	if lenth == 0 {
		return nil
	}

	if lenth == 1 {
		return lists[0]
	}
	lenth = lenth >> 1
	left := mergeKLists(lists[:lenth])
	right := mergeKLists(lists[lenth:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := dummyHead

	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
			continue
		}

		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}
	}
	return dummyHead.Next
}

// @lc code=end
