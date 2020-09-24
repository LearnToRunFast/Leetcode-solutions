class Solution {
    public int strStr(String haystack, String needle) {
        int needle_len = needle.length();
        if (needle_len == 0) return needle_len;
        
        int hay_len = haystack.length();
        if (needle_len > hay_len) return -1;
        
        for (int i = 0; i < hay_len; i++) {
            if (hay_len - i < needle_len) {
                break;
            }
            
            if (haystack.substring(i, needle_len + i).equals(needle)) {
                return i;
            }
        }
        
        return -1;
    }
}