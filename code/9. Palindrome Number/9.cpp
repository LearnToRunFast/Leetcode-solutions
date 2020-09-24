class Solution {
public:

    bool isPalindrome(int x) {
        if(x < 0 || (x % 10 == 0 && x != 0)) return false;
        if (x == 0) return true;
        
        int revertedNumber = 0;
        while(x > revertedNumber) {
            revertedNumber = revertedNumber * 10 + x % 10;
            x /= 10;
        }
        // if the length is even,       if the length is odd remove the middle number
        return x == revertedNumber || x == revertedNumber/10;
    }
};
