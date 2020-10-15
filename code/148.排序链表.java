/*
 * @lc app=leetcode.cn id=148 lang=java
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (66.80%)
 * Likes:    777
 * Dislikes: 0
 * Total Accepted:    96.7K
 * Total Submissions: 144K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 * 
 * 示例 1:
 * 
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 * 
 * 
 * 示例 2:
 * 
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
// // recursion version
// class Solution {
//     public ListNode sortList(ListNode head) {
//         if (head == null || head.next == null) return head;

//         // find mid and split linklist into two
//         ListNode slow = head, fast = head;
//         while (fast.next != null && fast.next.next != null) {
//             slow = slow.next;
//             fast = fast.next.next;
//         }
//         ListNode tail = slow.next;
//         slow.next = null;

//         ListNode left = sortList(head);
//         ListNode right = sortList(tail);
        
//         // merge part
//         // dummy head node
//         ListNode dummyHead = new ListNode(0);
//         ListNode ans = dummyHead; // keep head will need to return it as answer

//         while(left != null && right != null) {
//             if (left.val > right.val) {
//                 dummyHead.next = right;
//                 right = right.next;
//             } else {
//                 dummyHead.next = left;
//                 left = left.next;
//             }
//             dummyHead = dummyHead.next;
//         }

//         dummyHead.next = left != null ? left : right;
//         return ans.next;
//     }
// }

// iterative version
class Solution {
    private int getLength(ListNode node) {
        int len = 0;
        while(node != null) {
            len++;
            node = node.next;
        }
        return len;
    }
    private ListNode nextHead(int subLen, ListNode head) {
        ListNode tmp = head;
        while (tmp != null && subLen-- > 1) {
            tmp = tmp.next;
        }

        if (tmp == null) return null;
        // terminate the previous linklist with null
        ListNode ans = tmp.next;
        tmp.next = null;
        return ans; 
    }

    private ListNode merge(ListNode left, ListNode right) {
        ListNode dummyHead = new ListNode(0);
        ListNode ans = dummyHead;

        while (left != null && right != null) {
            if (left.val > right.val) {
                dummyHead.next = right;
                right = right.next;
            } else {
                dummyHead.next = left;
                left = left.next;
            }
            dummyHead = dummyHead.next;
        }
        dummyHead.next = left != null ? left : right;

        return ans.next;
    }
    public ListNode sortList(ListNode head) {
        int len = getLength(head);
        ListNode dummyHead = new ListNode(0);
        dummyHead.next = head;

        for (int subLen = 1; subLen < len; subLen *= 2) {
            ListNode prev = dummyHead;
            ListNode curr = dummyHead.next;
            while (curr != null) {
                ListNode leftHead = curr;
                ListNode rightHead = nextHead(subLen, leftHead);
                // need to be infront of merge
                curr = nextHead(subLen, rightHead); // next segment

                ListNode merged = merge(leftHead, rightHead);

                prev.next = merged;
                // move to tail of merged
                while (prev.next != null) {
                    prev = prev.next;
                }
            }
        }

        return dummyHead.next;
    }
}
// @lc code=end

