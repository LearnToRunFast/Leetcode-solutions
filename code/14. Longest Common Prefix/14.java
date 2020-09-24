class Solution {
    public String longestCommonPrefix(String[] strs) {
        if (strs.length == 0 || strs[0] == "") {
            return "";
        }
        
        Arrays.sort(strs, String.CASE_INSENSITIVE_ORDER);
        String start = strs[0];
        String end = strs[strs.length - 1];
        int len = Math.min(start.length(), end.length());
        for(int i = 0; i < len; i++) {
            if (start.charAt(i) != end.charAt(i)) {
                return start.substring(0, i);
            }
        }
        return start;
        
    }
}