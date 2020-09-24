class Solution {
    private static int getValue(char c) {
        switch (c) {
            case 'I':
                return 1;
            case 'V':
                return 5;
            case 'X':
                return 10;
            case 'L':
                return 50;
            case 'C':
                return 100;
            case 'D':
                return 500;
            case 'M':
                return 1000;
            default:
                throw new IllegalArgumentException("Illegal character");
        }

    }
    public int romanToInt(String s) {
        int res = 0, i = 0, n = s.length();
        while(i < n) {
            int curr = getValue(s.charAt(i));
            //if its last item or next item is larger than current item
            if (i == n - 1 || curr >= getValue(s.charAt(i + 1))) {
                res += curr;
            } else {
                res -= curr;
            }
            i++;
        }
        return res;
    }

}