/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (51.49%)
 * Likes:    406
 * Dislikes: 0
 * Total Accepted:    36.1K
 * Total Submissions: 69.5K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
*/

// @lc code=start
type maxHeap []int
type minHeap []int

// min heap
func (m minHeap) Len() int {
	return len(m)
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Top() int {
	return h[0]
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

// max heap
func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (m *maxHeap) Top() int {
	return (*m)[0]
}
func (m *maxHeap) Pop() interface{} {
	h := *m
	n := len(h)
	old := h[n-1]
	*m = h[:n-1]
	return old
}
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

type MedianFinder struct {
	left  *maxHeap
	right *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minheap := &minHeap{}
	maxheap := &maxHeap{}
	heap.Init(minheap)
	heap.Init(maxheap)
	return MedianFinder{maxheap, minheap}
}

func (f *MedianFinder) AddNum(num int) {
	heap.Push(f.left, num)
	heap.Push(f.right, heap.Pop(f.left))

	if f.left.Len() < f.right.Len() {
		heap.Push(f.left, heap.Pop(f.right))
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.left.Len() > f.right.Len() {
		return float64(f.left.Top())
	}
	return float64(f.left.Top()+f.right.Top()) / 2
}

// func (f *MedianFinder) AddNum(num int) {
// 	lLen, rLen := f.left.Len(), f.right.Len()
// 	if lLen == 0 {
// 		heap.Push(f.left, num)
// 		return
// 	}
// 	x := num - f.left.Top()
// 	if x < 0 || x == 0 && lLen <= rLen {
// 		heap.Push(f.left, num)
// 		lLen++
// 	} else if x > 0 || x == 0 && lLen > rLen {
// 		heap.Push(f.right, num)
// 		rLen++
// 	}
// 	diff := lLen - rLen
// 	if diff > 1 {
// 		heap.Push(f.right, heap.Pop(f.left))
// 	} else if diff < -1 {
// 		heap.Push(f.left, heap.Pop(f.right))

// 	}
// }

// func (f *MedianFinder) FindMedian() float64 {
// 	l, lLen := f.left, f.left.Len()
// 	r, rLen := f.right, f.right.Len()

// 	if lLen > rLen {
// 		return float64(l.Top())
// 	} else if lLen < rLen {
// 		return float64(r.Top())
// 	}
// 	return float64(l.Top()+r.Top()) / 2
// }
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

