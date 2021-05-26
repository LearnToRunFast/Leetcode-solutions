/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 *
 * https://leetcode-cn.com/problems/insert-delete-getrandom-o1/description/
 *
 * algorithms
 * Medium (49.64%)
 * Likes:    321
 * Dislikes: 0
 * Total Accepted:    27.2K
 * Total Submissions: 54.7K
 * Testcase Example:  '["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]\n' +
  '[[],[1],[2],[2],[],[1],[2],[]]'
 *
 * 设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
 *
 *
 * insert(val)：当元素 val 不存在时，向集合中插入该项。
 * remove(val)：元素 val 存在时，从集合中移除该项。
 * getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
 *
 *
 * 示例 :
 *
 *
 * // 初始化一个空的集合。
 * RandomizedSet randomSet = new RandomizedSet();
 *
 * // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
 * randomSet.insert(1);
 *
 * // 返回 false ，表示集合中不存在 2 。
 * randomSet.remove(2);
 *
 * // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
 * randomSet.insert(2);
 *
 * // getRandom 应随机返回 1 或 2 。
 * randomSet.getRandom();
 *
 * // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
 * randomSet.remove(1);
 *
 * // 2 已在集合中，所以返回 false 。
 * randomSet.insert(2);
 *
 * // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
 * randomSet.getRandom();
 *
 *
*/

// @lc code=start
type RandomizedSet struct {
	set  map[int]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.set[val]; !ok {
		s.list = append(s.list, val)
		s.set[val] = len(s.list) - 1
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (s *RandomizedSet) Remove(val int) bool {
	if v, ok := s.set[val]; ok {
		// swap and remove val
		last := len(s.list) - 1
		lastVal := s.list[last]
		s.list[v] = lastVal
		s.list = s.list[:last]

		// update idx in map and remove val from map
		s.set[lastVal] = s.set[val]
		delete(s.set, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (s *RandomizedSet) GetRandom() int {
	// intn panic if n <= 0
	n := len(s.list)
	if n == 0 {
		return -1
	}
	idx := rand.Intn(n)
	return s.list[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

