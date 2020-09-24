class Solution:
    def longestPalindrome(self, s: str) -> str:
        def searchAround(s, i, j):
            l, r = i, j
            while l >= 0 and r < len(s) and s[r] == s[l]:
                l -= 1
                r += 1
            
            return r - l - 1
        
        start, end = 0, 0
        
        for i in range(len(s)):
            length1 = searchAround(s, i, i)
            length2 = searchAround(s, i, i + 1)
            max_len = max(length1, length2)
            if max_len > end - start:
                start = i - int((max_len - 1) / 2)
                end = i + int(max_len / 2)
        return s[start:end + 1]
            
        
            
                