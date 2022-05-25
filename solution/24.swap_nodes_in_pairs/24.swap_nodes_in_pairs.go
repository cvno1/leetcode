/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
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
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	for p := dummyHead; p.Next != nil && p.Next.Next != nil; {
		n1 := p.Next
		n2 := p.Next.Next
		p.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		p = n1
	}

	return dummyHead.Next
}

// @lc code=end
