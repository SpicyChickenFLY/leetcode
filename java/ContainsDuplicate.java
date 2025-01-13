import java.util.HashMap;

class Solution {
    // Run time: 57.36%, Memory Usage: 62.93%
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            if (map.containsKey(num))
                return true;
            else
                map.put(num, 0);
        }
        return false;
    }
}