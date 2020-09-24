class Solution {
    private static boolean isAlnum(char c) {
        if ((c >= 'a' && c <= 'z') ||
            (c >= 'A' && c <= 'Z') ||
            (c >= '0' && c <= '9')) {
            return true;
        }
        return false;
    }
    public boolean isPalindrome(String s) {
        int i = 0, j = s.length() - 1;
        while (i < j) {
            while (i < j && !isAlnum(s.charAt(i))) i++;
            while (i < j && !isAlnum(s.charAt(j))) j--;
            if(Character.toLowerCase(s.charAt(i)) != Character.toLowerCase(s.charAt(j))) return false;
            i++;
            j--;
        }
        return true;
    }
}

