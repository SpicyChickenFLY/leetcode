import java.util.*;

class Solution {
    private Map dict;

    public int recursion(int[] nums, int index) {
        // System.out.println(index);
        if (dict.containsKey(index))
            return (int) dict.get(index);
        if (index >= nums.length)
            return 0;
        int maxResult = 0;
        for (int i = index; i < nums.length; i++) {
            int temp = nums[i] + recursion(nums, i+2);
            if (temp > maxResult)
                maxResult = temp;
        }
        dict.put(index, maxResult);
        return maxResult;
    }

    // Runtime: 6.07%, Memory Usage: 98.95%
    public int rob(int[] nums) {
        dict = new HashMap();
        return recursion(nums, 0);
    }
}

public class HouseRobber {
    public static void main(String[] args) {
        Solution s = new Solution();
        int [][]test_cases = {
            {1,2,3,1},
            {2,7,9,3,1}
        };
        for (int []test_case : test_cases)
            System.out.println(s.rob(test_case));
    }
}