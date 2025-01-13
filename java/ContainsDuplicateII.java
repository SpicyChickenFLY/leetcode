import java.util.HashMap;

class Solution {
    // Run time: 5.02%, Memory Usage: 100.00%
    public boolean containsDuplicate(int[] nums, int k) {
        for (int i = 0; i < nums.length; i++) 
            for (int j = 1; j <= k; j++)
                if (i + j < nums.length && nums[i] == nums[i + j])
                    return true;
        return false;
    }
}

