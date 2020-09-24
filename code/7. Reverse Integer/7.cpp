class Solution {
public:
    int reverse(int x) {
        
        int pop = 0, temp = 0;
        int max = INT_MAX / 10;
        int min = INT_MIN / 10;
        while(x != 0) {
            pop = x % 10;
            if (temp > max || (temp == max && pop > 7)) return 0;
            if (temp < min || (temp == min && pop < -8)) return 0;
            temp = temp * 10 + pop;
            x /= 10;
        }
        return temp;
    }
};
