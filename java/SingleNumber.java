class Solution {
    // Run time: 99.082.%, Memory Usage: 99.11%
    public int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums)
            result ^= num;
        return result;
    }
}

public class SingleNumber {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] input = {4, 1, 2, 1, 2};
        System.out.println(s.singleNumber(input)); 
    }
}