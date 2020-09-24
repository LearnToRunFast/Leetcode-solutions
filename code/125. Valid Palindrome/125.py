class Solution:
    def isalnum(self, c):
        return 'A' <= c <= 'Z' or 'a' <= c <= 'z' or '0' <= c <= '9'
        
    def isPalindrome(self, s: str) -> bool:
        i = 0
        j = len(s) - 1
        while i < j:
            while i < j and not self.isalnum(s[i]):
                i += 1
            while i < j and not self.isalnum(s[j]):
                j -= 1
            if s[i].lower() != s[j].lower():
                return False
            i += 1
            j -= 1
        return True