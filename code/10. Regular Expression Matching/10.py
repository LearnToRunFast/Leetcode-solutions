class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        if not p:
            return not s
        
        match = bool(s) and p[0] in [s[0], '.']
        # if my current p is least 2 length long
        # I will check whether * is in second place.
        if len(p) >= 2 and p[1] == '*':
            # self.isMatch(s, p[2:]) is letter* appear 0 times
            return self.isMatch(s, p[2:]) or match and self.isMatch(s[1:], p)
        else:
            return match and self.isMatch(s[1:], p[1:])
