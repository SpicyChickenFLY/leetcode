class Solution {
    // Runtime: 88.88%, Memory Usage: 100.00%
    public int maxProfit(int[] prices) {
        if (prices.length <= 1)
            return 0;    
        for (int i = prices.length - 1; i > 0; i--) {
            prices[i] -= prices[i - 1];
        }
        int max = prices[1];
        if (max < 0)
            max = 0;
        for (int i = 2; i < prices.length; i++) {
            if (prices[i - 1] > 0)
                prices[i] += prices[i - 1];
            if (prices[i] > max)
                max = prices[i];
        }
        return max;
    }
}

public class BestTimeToBuyAndSellStock {
    public static void main(String[] args) {
        Solution s = new Solution();
        int [][]test_cases = {
            {7, 1, 5, 3, 6, 4},
            {7, 6, 4, 3, 1},
            {1, 2}
        };
        for (int[] test_case : test_cases) {
            System.out.println("Result: " + s.maxProfit(test_case));
        }
    }
}
