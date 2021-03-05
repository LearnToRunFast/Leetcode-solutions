/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (50.33%)
 * Likes:    466
 * Dislikes: 0
 * Total Accepted:    88.1K
 * Total Submissions: 175K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	l, m, r := dummy, head, head.Next
	for ; r != nil; r = r.Next {
		if r.Val > m.Val {
			if m.Next == r {
				l.Next = m
				l = l.Next
				m = m.Next
			} else {
				m = r
			}
		}
	}
	if m.Next == nil {
		l.Next = m
		l = l.Next
	}
	l.Next = nil
	return dummy.Next
}

// @lc code=end

