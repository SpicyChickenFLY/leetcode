class Solution {
    // Runtime: 100.00%, Memory Usage: 97.34%
    public int maxProfit(int[] prices) {
        if (prices.length <= 1)
            return 0;
        for (int i = prices.length - 1; i > 0; i--) {
            prices[i] -= prices[i - 1];
        }
        int max = 0;
        for (int i = 1; i < prices.length; i++) {
            if (prices[i] > 0)
                max += prices[i];
        }
        return max;
    }
}

public class BestTimeToBuyAndSellStockII {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[][] test_cases = { { 7, 1, 5, 3, 6, 4 }, { 7, 6, 4, 3, 1 }, { 1, 2, 3, 4, 5 } };
        for (int[] test_case : test_cases) {
            System.out.println("Result: " + s.maxProfit(test_case));
        }
    }
}
