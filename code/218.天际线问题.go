/*
 * @lc app=leetcode.cn id=218 lang=golang
 *
 * [218] 天际线问题
 *
 * https://leetcode-cn.com/problems/the-skyline-problem/description/
 *
 * algorithms
 * Hard (46.38%)
 * Likes:    383
 * Dislikes: 0
 * Total Accepted:    14.2K
 * Total Submissions: 30.7K
 * Testcase Example:  '[[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]'
 *
 * 城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回由这些建筑物形成的 天际线 。
 *
 * 每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti]
 * 表示：
 *
 *
 * lefti 是第 i 座建筑物左边缘的 x 坐标。
 * righti 是第 i 座建筑物右边缘的 x 坐标。
 * heighti 是第 i 座建筑物的高度。
 *
 *
 * 天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序
 * 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0
 * ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。
 *
 * 注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...]
 * 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
 * 输出：[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
 * 解释：
 * 图 A 显示输入的所有建筑物的位置和高度，
 * 图 B 显示由这些建筑物形成的天际线。图 B 中的红点表示输出列表中的关键点。
 *
 * 示例 2：
 *
 *
 * 输入：buildings = [[0,2,3],[2,5,3]]
 * 输出：[[0,3],[5,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0 i < righti
 * 1 i
 * buildings 按 lefti 非递减排序
 *
 *
 */
package main

import (
	"container/heap"
	"sort"
)

// @lc code=start
type building struct{ x, h int }
type buildingHeap []*building

func (bh buildingHeap) Len() int            { return len(bh) }
func (bh buildingHeap) Less(i, j int) bool  { return bh[i].h > bh[j].h }
func (bh buildingHeap) Swap(i, j int)       { bh[i], bh[j] = bh[j], bh[i] }
func (bh *buildingHeap) Push(x interface{}) { *bh = append(*bh, x.(*building)) }
func (bh *buildingHeap) Pop() interface{} {
	old := *bh
	n := len(old)
	x := old[n-1]
	*bh = old[:n-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	points := make([]int, 0, n*2)
	for _, b := range buildings {
		points = append(points, b[0], b[1])
	}
	sort.Ints(points)
	bh := &buildingHeap{}
	heap.Push(bh, &building{1<<63 - 1, 0})
	idx := 0
	ans := [][]int{}
	for _, point := range points {
		// left
		for idx < n && buildings[idx][0] == point {
			heap.Push(bh, &building{buildings[idx][1], buildings[idx][2]})
			idx++
		}

		for (*bh)[0].x <= point {
			heap.Pop(bh)
		}
		m := len(ans)
		// if the height is diff , add to ans
		if m == 0 || ans[m-1][1] != (*bh)[0].h {
			ans = append(ans, []int{point, (*bh)[0].h})
		}
	}
	return ans
}

// @lc code=end
