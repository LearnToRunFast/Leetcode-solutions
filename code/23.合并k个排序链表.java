/*
 * @lc app=leetcode.cn id=23 lang=java
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (47.59%)
 * Likes:    348
 * Dislikes: 0
 * Total Accepted:    45.5K
 * Total Submissions: 95.5K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 * 
 * 示例:
 * 
 * 输入:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
 * 
 */
/**
 * Definition for singly-linked list. public class ListNode { int val; ListNode
 * next; ListNode(int x) { val = x; } }
 */
class Solution {
    private static ListNode merge(ListNode l1, ListNode l2) {
        ListNode curr = new ListNode(0); // dummy node
        ListNode head = curr;

        while (l1 != null && l2 != null) {
            if (l1.val <= l2.val) {
                curr.next = l1;
                l1 = l1.next;
            } else {
                curr.next = l2;
                l2 = l1;
                l1 = curr.next.next;
            }
            curr = curr.next;
        }

        if (l1 == null) {
            curr.next = l2;
        } else {
            curr.next = l1;
        }
        return head.next;
    }

    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0)
            return null;

        int len = lists.length;
        // start at second eg. 0 1 2 3
        // increase double
        for (int i = 1; i < len; i *= 2) {
            // !!! increse by double of i not j
            for (int j = 0; j < len - i; j += 2 * i) {
                // remember to return to first list
                lists[j] = merge(lists[j], lists[j + i]);
            }
        }
        return lists[0];
    }
}
