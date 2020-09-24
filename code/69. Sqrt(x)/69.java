class Solution {
    public int mySqrt(int x) {
        // 模版 binary search
        if (x <= 1) return x;
        
        long l = 1, h = x >> 1;
        while (l < h) {
            long mid = (l + h + 1) >>> 1;
            long temp = mid * mid;
            if (temp > x) {
                h = mid - 1;
            } else {
                l = mid;
            }
        }
        return (int)l;
    }
}

/*
class Solution {
    public int mySqrt(int x) {
        if (x <= 1) return x;
        
        long l = 1, h = x >> 1;
        while (l <= h) {
            long mid = (l + h + 1) >>> 1;
            long temp = mid * mid;
            if (temp > x) h = mid - 1;
            if (temp < x) l = mid + 1;
            if (temp == x) return (int)mid;
        }
        return (int)h;
    }
}
*/