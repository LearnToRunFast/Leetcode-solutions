class Solution {
    public String convertToTitle(int n) {
        String temp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        StringBuilder ans = new StringBuilder();

        while (n > 0) {
            ans.append(temp.charAt(--n % 26) );
            n /= 26;
        }
        return ans.reverse().toString();
    }
}