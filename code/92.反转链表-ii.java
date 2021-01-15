/*
 * @lc app=leetcode.cn id=92 lang=java
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (51.95%)
 * Likes:    638
 * Dislikes: 0
 * Total Accepted:    95.9K
 * Total Submissions: 184.3K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 * 
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 * 
 * 示例:
 * 
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
// class Solution {
//     ListNode successor;
//     private ListNode reverseN(ListNode head, int n) {
//         if (n == 1) {
//             successor = head.next;
//             return head;
//         }
//         ListNode reversed = reverseN(head.next, n - 1);

//         head.next.next = head;
//         head.next = successor; // if no n, this will be null
//         return reversed;
//     }
//     public ListNode reverseBetween(ListNode head, int m, int n) {
//         if (m == 1) {
//             return reverseN(head, n);
//         }
//         head.next = reverseBetween(head.next, m - 1, n - 1);
//         return head;
//     }
// }
class Solution {

    public ListNode reverseBetween(ListNode head, int m, int n) {
        ListNode dummy = new ListNode();
        dummy.next = head;

        ListNode pre = dummy;
        // stop at 1 will be curr.
        while (m-- > 1) {
            n--;
            pre = pre.next;
        }
        n--; // we need extra -1 because every time we exchange two item
        // so there will be n - 1 exchange instead of n
        ListNode curr = pre.next;
        for (int i = 0; i < n; i++) {
            ListNode temp = curr.next;
            curr.next = curr.next.next;

            temp.next = pre.next;
            pre.next = temp;
        }
        return dummy.next;
    }
}
// @lc code=end

