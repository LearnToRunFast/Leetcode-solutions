class Solution {
    public int myAtoi(String str) {
        
        if (str == null) {
            return 0;
        }
        int ans = 0;
        int len = 0;
        int numOfSign = 0;
        int numOfZero = 0;
        boolean isPositive = true;
        
        for (int i = 0; i < str.length(); i++) {
            if (str.charAt(i) - '0' >= 0 && str.charAt(i) - '0' <= 9) {
                if (len == 0 && str.charAt(i) - '0' == 0) {
                    numOfZero++;
                    continue;
                }
                
                int pop = str.charAt(i) - '0';
                if (ans > Integer.MAX_VALUE / 10 || 
                    (ans == Integer.MAX_VALUE / 10 && pop > 7) ||
                    (ans == Integer.MIN_VALUE / 10 && pop > 8)) {
                    return isPositive ? Integer.MAX_VALUE : Integer.MIN_VALUE;
                }
                ans = ans * 10 + pop;
                len++;
                // make sure sign is before any 0
            } else if ((str.charAt(i) == '-' || str.charAt(i) == '+') && numOfSign == 0 && numOfZero == 0 && len == 0) {
                numOfSign++;
                if (str.charAt(i) == '-') isPositive = false;
            } else if (str.charAt(i) == ' ' && numOfSign == 0 && numOfZero == 0 && len == 0) {
                continue;
            } else {
                break;
            }
        }
        return isPositive ? ans : -ans;           
    }
}