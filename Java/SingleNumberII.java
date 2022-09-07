class Solution {
    // Run time: 100.00%, Memory Usage: 99.75% 
    public int singleNumber(int[] nums) {
        if (nums.length == 0)
            return 0;
        int result = 0, carry = 0;
        for (int i = 0; i < nums.length; i++) {
            // simulate summary
            carry |= result & nums[i];
            result ^= nums[i];
            // count if three
            int threeCounter = result & carry;
            // 3bit XOR
            result ^= threeCounter;
            carry ^= threeCounter;
        }
        return result;
    }
}

public class SingleNumberII {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] input = { 4, 1, 2, 1, 2 };
        System.out.println(s.singleNumber(input));
    }
}