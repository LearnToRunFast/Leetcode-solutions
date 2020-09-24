class Solution {
    public String addBinary(String a, String b) {
    
    int len = Math.max(a.length(), b.length()) + 1;
    while (a.length() != len) a = '0' + a;
    while (b.length() != len) b = '0' + b;
    
    StringBuilder s = new StringBuilder();
    int carry = 0;
    for (int i = len - 1; i >= 0; i--) {
        int temp = (a.charAt(i) - '0') + (b.charAt(i) - '0') + carry;
        if (temp > 1) {
            s.append(temp - 2);
            carry = 1;
        } else {
            carry = 0;
            s.append(temp);
        }
    }
    s.reverse();
    return s.charAt(0) == '0' ? s.substring(1, len).toString() : s.toString();
        
    }
}