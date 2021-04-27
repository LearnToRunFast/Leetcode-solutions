/*
 * @lc app=leetcode.cn id=273 lang=golang
 *
 * [273] 整数转换英文表示
 *
 * https://leetcode-cn.com/problems/integer-to-english-words/description/
 *
 * algorithms
 * Hard (30.29%)
 * Likes:    144
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 33.5K
 * Testcase Example:  '123'
 *
 * 将非负整数 num 转换为其对应的英文表示。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = 123
 * 输出："One Hundred Twenty Three"
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = 12345
 * 输出："Twelve Thousand Three Hundred Forty Five"
 *
 *
 * 示例 3：
 *
 *
 * 输入：num = 1234567
 * 输出："One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
 *
 *
 * 示例 4：
 *
 *
 * 输入：num = 1234567891
 * 输出："One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven
 * Thousand Eight Hundred Ninety One"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 *
 *
 */

// @lc code=start
func toPartWord(num int) []string {
	result := []string{}
	switch num / 100 {
	case 1:
		result = append(result, "One Hundred")
	case 2:
		result = append(result, "Two Hundred")
	case 3:
		result = append(result, "Three Hundred")
	case 4:
		result = append(result, "Four Hundred")
	case 5:
		result = append(result, "Five Hundred")
	case 6:
		result = append(result, "Six Hundred")
	case 7:
		result = append(result, "Seven Hundred")
	case 8:
		result = append(result, "Eight Hundred")
	case 9:
		result = append(result, "Nine Hundred")
	}

	switch num % 100 {
	case 0:
		return result
	case 10:
		result = append(result, "Ten")
		return result
	case 11:
		result = append(result, "Eleven")
		return result
	case 12:
		result = append(result, "Twelve")
		return result
	case 13:
		result = append(result, "Thirteen")
		return result
	case 14:
		result = append(result, "Fourteen")
		return result
	case 15:
		result = append(result, "Fifteen")
		return result
	case 16:
		result = append(result, "Sixteen")
		return result
	case 17:
		result = append(result, "Seventeen")
		return result
	case 18:
		result = append(result, "Eighteen")
		return result
	case 19:
		result = append(result, "Nineteen")
		return result
	}
	num %= 100
	switch num / 10 {
	case 2:
		result = append(result, "Twenty")
	case 3:
		result = append(result, "Thirty")
	case 4:
		result = append(result, "Forty")
	case 5:
		result = append(result, "Fifty")
	case 6:
		result = append(result, "Sixty")
	case 7:
		result = append(result, "Seventy")
	case 8:
		result = append(result, "Eighty")
	case 9:
		result = append(result, "Ninety")
	}
	switch num % 10 {
	case 1:
		result = append(result, "One")
	case 2:
		result = append(result, "Two")
	case 3:
		result = append(result, "Three")
	case 4:
		result = append(result, "Four")
	case 5:
		result = append(result, "Five")
	case 6:
		result = append(result, "Six")
	case 7:
		result = append(result, "Seven")
	case 8:
		result = append(result, "Eight")
	case 9:
		result = append(result, "Nine")
	}
	return result
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	unit := []string{"", "Thousand", "Million", "Billion", "Trillion"}
	ans := []string{}
	unitOrder := 0
	for num != 0 {
		part := toPartWord(num % 1000)
		if len(part) > 0 && unitOrder > 0 {
			part = append(part, unit[unitOrder])
		}
		ans = append(part, ans...)
		num /= 1000
		unitOrder++
	}
	return strings.Join(ans, " ")

}

// @lc code=end

