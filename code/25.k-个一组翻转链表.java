/*
 * @lc app=leetcode.cn id=25 lang=java
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (54.57%)
 * Likes:    269
 * Dislikes: 0
 * Total Accepted:    22.4K
 * Total Submissions: 41K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * 
 * k 是一个正整数，它的值小于或等于链表的长度。
 * 
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 * 
 * 示例 :
 * 
 * 给定这个链表：1->2->3->4->5
 * 
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 * 
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 * 
 * 说明 :
 * 
 * 
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * 
 * 
 */
/**
 * Definition for singly-linked list. public class ListNode { int val; ListNode
 * next; ListNode(int x) { val = x; } }
 */
// class Solution {

//     // public ListNode reverseKGroup(ListNode head, int k) {
//     //     ListNode curr = head;
//     //     int count = 0;
//     //     while (curr != null && count < k) {
//     //         curr = curr.next;
//     //         count++;
//     //     }
//     //     // if curr is not null curr will be point to next set
//     //     if (count == k) {
//     //         // assume it works
//     //         curr = reverseKGroup(curr, k);

//     //         while (count-- > 0) {
//     //             ListNode backUp = head.next;
//     //             head.next = curr;
//     //             curr = head;
//     //             head = backUp;
//     //         }
//     //         head = curr;
//     //     }
//     //     return head;
//     // }
    
//     private ListNode reverse(ListNode head) {
//         ListNode tail = null;
//         while (head != null) {
//             ListNode next = head.next;
//             head.next = tail;
//             tail = head;
//             head = next;
//         }
//         return tail;
//     }
//     public ListNode reverseKGroup(ListNode head, int k) {
//         ListNode pre = new ListNode(0);
//         ListNode end = pre;
//         ListNode ans = pre;
//         pre.next = head;

//         while (end.next != null) {
//             for (int i = 0; i < k && end != null; i++) {
//                 end = end.next;
//             }

//             if (end == null) break; // less than k

//             ListNode start = pre.next;
//             ListNode nextStart = end.next;

//             end.next = null; // set terminator

//             pre.next = reverse(start);
//             pre = start;
//             start.next = nextStart;
            
//             end = pre;

//         }
//         return ans.next;

//     }
// }
class Solution {
    
    private ListNode reverse(ListNode head, ListNode end) {
        ListNode pre = null;
        while (head != end) {
            ListNode next = head.next;
            head.next = pre;
            pre = head;
            head = next;
        }
        return pre;
    }
    public ListNode reverseKGroup(ListNode head, int k) {
        if (head == null) return null;
        ListNode curr = head;
        for (int i = 0; i < k; i++) {
            if (curr == null) return head;
            curr = curr.next;
        }
        ListNode newHead = reverse(head, curr);

        head.next = reverseKGroup(curr, k);
        return newHead;

    }
}
