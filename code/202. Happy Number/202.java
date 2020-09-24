class Solution {
    static int getSum(int n) {
        int sum = 0;
        while (n > 0) {
            int r = n % 10;
            sum += r * r;
            n /= 10;
        }
        return sum;
    }
    public boolean isHappy(int n) {
        int slow = n, fast = getSum(n);
        
        while (slow != fast) {
            slow = getSum(slow);
            fast = getSum(fast);
            fast = getSum(fast);
        }
        return slow == 1;
    }
}