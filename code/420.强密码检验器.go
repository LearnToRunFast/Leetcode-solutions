/*
 * @lc app=leetcode.cn id=420 lang=golang
 *
 * [420] 强密码检验器
 *
 * https://leetcode-cn.com/problems/strong-password-checker/description/
 *
 * algorithms
 * Hard (20.46%)
 * Likes:    65
 * Dislikes: 0
 * Total Accepted:    2.5K
 * Total Submissions: 12.4K
 * Testcase Example:  '"a"'
 *
 * 一个强密码应满足以下所有条件：
 *
 *
 * 由至少6个，至多20个字符组成。
 * 至少包含一个小写字母，一个大写字母，和一个数字。
 * 同一字符不能连续出现三次 (比如 "...aaa..." 是不允许的, 但是 "...aa...a..." 是可以的)。
 *
 *
 * 编写函数 strongPasswordChecker(s)，s 代表输入字符串，如果 s 已经符合强密码条件，则返回0；否则返回要将 s
 * 修改为满足强密码条件的字符串所需要进行修改的最小步数。
 *
 * 插入、删除、替换任一字符都算作一次修改。
 *
 */

// @lc code=start
func find(arr []int, findMaxReminder bool) int {
	a, b, c := -1, -1, -1
	for i := range arr {
		if arr[i] <= 0 {
			continue
		}
		if arr[i]%3 == 2 {
			if a != -1 && arr[a] > arr[i] {
				continue
			}
			a = i
			continue
		}
		if arr[i]%3 == 1 {
			if b != -1 && arr[b] > arr[i] {
				continue
			}
			b = i
			continue
		}
		if arr[i]%3 == 0 {
			if c != -1 && arr[c] > arr[i] {
				continue
			}
			c = i
		}
	}
	if findMaxReminder {
		if a != -1 {
			return a
		} else if b != -1 {
			return b
		} else {
			return c
		}
	}
	if c != -1 {
		return c
	} else if b != -1 {
		return b
	} else {
		return a
	}
}
func strongPasswordChecker(password string) int {
	n := len(password)

	count, step, missTypes := 0, 0, 3
	var curr byte
	lower, upper, num := false, false, false
	sames := []int{}
	for i := range password {
		if !num && password[i] >= '0' && password[i] <= '9' {
			num = true
			missTypes--
		} else if !lower && password[i] >= 'a' && password[i] <= 'z' {
			lower = true
			missTypes--
		} else if !upper && password[i] >= 'A' && password[i] <= 'Z' {
			upper = true
			missTypes--
		}
		if curr != password[i] {
			curr = password[i]
			if count >= 3 {
				sames = append(sames, count)
			}
			count = 1
			continue
		}
		count++
		if count%3 == 0 {
			step++
		}
	}
	if count >= 3 {
		sames = append(sames, count)
	}
	ans := 0
	// if it's less than 6 take max of missing type
	if n < 6 {
		ans = 6 - n
		if ans < missTypes {
			ans = missTypes
		}
		return ans
	}
	if n >= 6 && n <= 20 {
		// update only
		ans = step
		if ans < missTypes {
			ans = missTypes
		}
		return ans
	}
	extra := n - 20
	// 1. replace continuous with missingType if any
	// order %3 == 2 first, %3 == 1 second %3 == 0 last
	idx := find(sames, true)
	for idx != -1 && missTypes > 0 {
		sames[idx] -= 3
		ans++
		missTypes--
		idx = find(sames, true)
	}
	// 2. remove extra
	idx = find(sames, false)
	for idx != -1 && extra > 0 {
		extra--
		ans++
		sames[idx]--
		idx = find(sames, false)
	}
	m := 0
	for i := range sames {
		if sames[i] <= 0 {
			continue
		}
		m += sames[i] / 3
	}
	ans += extra + missTypes + m
	return ans
}

// @lc code=end

