import java.util.HashMap;
import java.util.Map;

class Solution {
    private Map<Integer, Integer> record = new HashMap<>();
    public int climbStairs(int n) {
        if (record.containsKey(n)) {
            return record.get(n);
        }
        if (n == 0 || n == 1) {
            return 1;
        }
        int result = climbStairs(n - 1) + climbStairs(n - 2);
        record.put(n, result);
        return result;
    }
}

public class ClimbingStairs {
    public static void main(String[] args) {
        Solution s = new Solution();
        for (int i = 0; i < 10; i++) {
            System.out.println(s.climbStairs(i));
        }
    }
}