/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 *
 * https://leetcode-cn.com/problems/design-linked-list/description/
 *
 * algorithms
 * Medium (32.12%)
 * Likes:    296
 * Dislikes: 0
 * Total Accepted:    65.2K
 * Total Submissions: 202.1K
 * Testcase Example:  '["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]\n' +
  '[[],[1],[3],[1,2],[1],[1],[1]]'
 *
 * 设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next
 * 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
 *
 * 在链表类中实现这些功能：
 *
 *
 * get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
 * addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
 * addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
 * addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index
 * 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
 * deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
 *
 *
 *
 *
 * 示例：
 *
 * MyLinkedList linkedList = new MyLinkedList();
 * linkedList.addAtHead(1);
 * linkedList.addAtTail(3);
 * linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
 * linkedList.get(1);            //返回2
 * linkedList.deleteAtIndex(1);  //现在链表是1-> 3
 * linkedList.get(1);            //返回3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 所有val值都在 [1, 1000] 之内。
 * 操作次数将在  [1, 1000] 之内。
 * 请不要使用内置的 LinkedList 库。
 *
 *
*/

// @lc code=start
type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}
type MyLinkedList struct {
	dummy *Node
}

func Constructor() MyLinkedList {
	node := &Node{
		Val: -1,
	}
	node.Next = node
	node.Pre = node
	return MyLinkedList{node}
}

func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	for index > 0 && head != this.dummy {
		head = head.Next
		index--
	}
	if index != 0 {
		return -1
	}
	return head.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val, Next: this.dummy.Next, Pre: this.dummy}
	this.dummy.Next.Pre = node
	this.dummy.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val, Next: this.dummy, Pre: this.dummy.Pre}
	this.dummy.Pre.Next = node
	this.dummy.Pre = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.Next
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}

	node := &Node{Val: val, Next: head, Pre: head.Pre}
	head.Pre.Next = node
	head.Pre = node

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.dummy.Next == this.dummy || index < 0 {
		return
	}
	head := this.dummy.Next
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index == 0 {
		head.Pre.Next = head.Next
		head.Next.Pre = head.Pre
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

