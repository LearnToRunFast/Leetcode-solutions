## dp[l, r] = (s[l] == s[r] and (r - l <= 2 or dp[l + 1, r - 1]))
class Solution:
    def longestPalindrome(self, s: str) -> str:
    	size = len(s)
    	if size <= 1:
    		return s

    	# 2d array with default value false
    	dp = [[False for _ in range(size)] for _ in range(size)]

    	res, max_len = s[0], 1

    	for r in range(1, size):
    		for l in range(r):
    			if s[l] == s[r] and ((r - l) <=2 or dp[l + 1][r - 1]):
    				dp[l][r] = True
    				cur_len = r - l + 1
    				if cur_len > max_len:
    					max_len = cur_len
    					res = s[l: r + 1]

    	return res