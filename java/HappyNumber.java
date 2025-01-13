import java.lang.Math;
import java.util.HashMap;

class Solution {
    // Runtime: 66.86%, Memory Usage: 10.60%
    private HashMap<Integer, Boolean> record; 

    public boolean isHappy(int n) {
        if (record.containsKey(n))
            return record.get(n);
        record.put(n, false);
        int sum = 0;
        while (n > 0) {
            sum += Math.pow(n % 10, 2);
            n /= 10;
        }
        if (sum == 0)
            return false;
        else if (sum == 1) {
            record.put(sum, true);
            return true;            
        }
        else {
            return isHappy(sum);
        }
    }

    Solution() {
        record = new HashMap<>();
    }
}

public class HappyNumber {
    public static void main(String[] args) {
        Solution s = new Solution();
        int []testCases = {
            19,
            2,
            10,
            0,
            1
        };
        for (int testCase : testCases) {
            System.out.println(testCase + ":" + s.isHappy(testCase));
        }
    }
}