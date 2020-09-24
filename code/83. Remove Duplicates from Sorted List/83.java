/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        if (head == null) return null;
        
        ListNode ans = head;
        while (head != null) {
            if (head.next == null) break;
            
            while (head.val == head.next.val) {
                head.next = head.next.next;
                if (head.next == null) break;
            }

            head = head.next;
        }
        return ans;
    }
}