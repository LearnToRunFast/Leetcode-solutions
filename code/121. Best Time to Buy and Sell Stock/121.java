class Solution {
    public int maxProfit(int[] prices) {
        int maxProfit = 0, minPrice = Integer.MAX_VALUE;
        for (int i = 0; i < prices.length; i++) {
            if (minPrice > prices[i]) {
                minPrice = prices[i];
            } else {
                int diff = 
                maxProfit = Math.max(prices[i] - minPrice, maxProfit);
            }
        }

        return maxProfit;
    }
}