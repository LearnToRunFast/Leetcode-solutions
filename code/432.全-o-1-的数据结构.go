/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 *
 * https://leetcode-cn.com/problems/all-oone-data-structure/description/
 *
 * algorithms
 * Hard (37.79%)
 * Likes:    91
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 16.1K
 * Testcase Example:  '["AllOne","inc","inc","getMaxKey","getMinKey","inc","getMaxKey","getMinKey"]\n' +
  '[[],["hello"],["hello"],[],[],["leet"],[],[]]'
 *
 * 请你实现一个数据结构支持以下操作：
 *
 *
 * Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
 * Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否则使一个存在的 key 值减一。如果这个 key
 * 不存在，这个函数不做任何事情。key 保证不为空字符串。
 * GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串"" 。
 * GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
 *
 *
 *
 *
 * 挑战：
 *
 * 你能够以 O(1) 的时间复杂度实现所有操作吗？
 *
*/

// @lc code=start
// link list related
type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}
type List struct {
	head *ListNode
	tail *ListNode
}

func (l *List) PushFront(node *ListNode) {
	node.Prev = nil
	if l.head == nil {
		l.head = node
		l.tail = node
		node.Next = nil
	} else {
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}
}
func (l *List) PushBack(node *ListNode) {
	node.Next = nil
	if l.tail == nil {
		l.head = node
		l.tail = node
		node.Prev = nil
	} else {
		l.tail.Next = node
		node.Prev = l.tail
		l.tail = node
	}
}

func (l *List) InsertBefore(before, node *ListNode) {
	if before == l.head {
		l.PushFront(node)
	} else {
		node.Prev, node.Next = before.Prev, before
		before.Prev.Next, before.Prev = node, node
	}
}
func (l *List) InsertAfter(after, node *ListNode) {
	if after == l.tail {
		l.PushBack(node)
	} else {
		node.Prev, node.Next = after, after.Next
		after.Next.Prev, after.Next = node, node
	}
}
func (l *List) PopFront() *ListNode {
	if l.head == nil {
		return nil
	}
	node := l.head
	if l.head.Next == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.Next
	}
	return node
}
func (l *List) PopBack() *ListNode {
	if l.tail == nil {
		return nil
	}
	node := l.tail
	if l.tail.Prev == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.Prev
	}
	return node
}
func (l *List) Remove(node *ListNode) {
	if node == l.head {
		l.PopFront()
	} else if node == l.tail {
		l.PopBack()
	} else {
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	}

}

type AllOne struct {
	keyCountMap  map[string]int
	countKeysMap map[int]map[string]bool
	countListMap map[int]*ListNode
	list         *List
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		keyCountMap:  make(map[string]int),
		countKeysMap: make(map[int]map[string]bool),
		countListMap: make(map[int]*ListNode),
		list:         &List{},
	}
}

func (this *AllOne) add(key string, oldCount, newCount int) {
	if keys, ok := this.countKeysMap[newCount]; ok {
		keys[key] = true
	} else {
		this.countKeysMap[newCount] = map[string]bool{key: true}
	}
	if len(this.countKeysMap[newCount]) == 1 {
		node := &ListNode{Val: newCount}
		this.countListMap[newCount] = node
		if newCount > oldCount {
			if newCount != 1 {
				this.list.InsertAfter(this.countListMap[oldCount], node)
			} else {
				this.list.PushFront(node)
			}
		} else {
			this.list.InsertBefore(this.countListMap[oldCount], node)
		}

	}
}
func (this *AllOne) remove(key string, oldCount, newCount int) {
	keys := this.countKeysMap[oldCount]
	if len(keys) != 1 {
		delete(keys, key)
	} else {
		delete(this.countKeysMap, oldCount)
		this.list.Remove(this.countListMap[oldCount])
		delete(this.countListMap, oldCount)
	}

}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	this.keyCountMap[key]++
	newCount := this.keyCountMap[key]
	this.add(key, newCount-1, newCount)
	if newCount != 1 {
		this.remove(key, newCount-1, newCount)
	}

}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if _, ok := this.keyCountMap[key]; !ok {
		return
	}
	this.keyCountMap[key]--
	newCount := this.keyCountMap[key]
	if newCount != 0 {
		this.add(key, newCount+1, newCount)
	}
	this.remove(key, newCount+1, newCount)
}
func (this *AllOne) GetOneKey(node *ListNode) string {
	if node == nil {
		return ""
	}
	keys := this.countKeysMap[node.Val]
	for k, _ := range keys {
		return k
	}
	return ""
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	return this.GetOneKey(this.list.tail)
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	return this.GetOneKey(this.list.head)
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end

