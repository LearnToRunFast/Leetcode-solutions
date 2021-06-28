/*
 * @lc app=leetcode.cn id=423 lang=golang
 *
 * [423] 从英文中重建数字
 *
 * https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/description/
 *
 * algorithms
 * Medium (56.71%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    7.4K
 * Total Submissions: 13K
 * Testcase Example:  '"owoztneoer"'
 *
 * 给定一个非空字符串，其中包含字母顺序打乱的英文单词表示的数字0-9。按升序输出原始的数字。
 *
 * 注意:
 *
 *
 * 输入只包含小写英文字母。
 * 输入保证合法并可以转换为原始的数字，这意味着像 "abc" 或 "zerone" 的输入是不允许的。
 * 输入字符串的长度小于 50,000。
 *
 *
 * 示例 1:
 *
 *
 * 输入: "owoztneoer"
 *
 * 输出: "012" (zeroonetwo)
 *
 *
 * 示例 2:
 *
 *
 * 输入: "fviefuro"
 *
 * 输出: "45" (fourfive)
 *
 *
 */

// @lc code=start
func originalDigits(s string) string {
	count := make([]int, 'z'+1)
	for _, c := range s {
		count[c]++
	}
	dict := make([]int, 10)
	// letter "z" is present only in "zero"
	dict[0] = count['z']
	// letter "w" is present only in "two"
	dict[2] = count['w']
	// letter "u" is present only in "four"
	dict[4] = count['u']
	// letter "x" is present only in "six"
	dict[6] = count['x']
	// letter "g" is present only in "eight"
	dict[8] = count['g']
	// letter "h" is present only in "three" and "eight"
	dict[3] = count['h'] - dict[8]
	// letter "f" is present only in "five" and "four"
	dict[5] = count['f'] - dict[4]
	// letter "s" is present only in "seven" and "six"
	dict[7] = count['s'] - dict[6]
	// letter "i" is present in "nine", "five", "six", and "eight"
	dict[9] = count['i'] - dict[5] - dict[6] - dict[8]
	// letter "n" is present in "one", "nine", and "seven"
	dict[1] = count['n'] - dict[7] - 2*dict[9]
	var b strings.Builder
	for i := 0; i < 10; i++ {
		for j := 0; j < dict[i]; j++ {
			b.WriteString(strconv.Itoa(i))
		}
	}
	return b.String()
}

// @lc code=end

