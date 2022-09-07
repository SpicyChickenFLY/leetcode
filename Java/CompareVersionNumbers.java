

class Solution {
    // Run time: 10.11%, Memory Usage: 100.00%
    public int compareVersion(String version1, String version2) {
        String[] level1 = version1.split("\\.");
        String[] level2 = version2.split("\\.");
        
        System.out.println(level1.length + ":" + level2.length);

        int minLength = level1.length < level2.length ? level1.length: level2.length;
        for (int i = 0; i < minLength; i++) {
            if (Integer.parseInt(level1[i]) > Integer.parseInt(level2[i]))
                return 1;
            if (Integer.parseInt(level1[i]) < Integer.parseInt(level2[i]))
                return -1;
        }
        if (level1.length > level2.length){
            for (int i = minLength; i < level1.length; i++)
                if (Integer.parseInt(level1[i]) > 0)
                    return 1;
        }
        if (level1.length < level2.length) {
            for (int i = minLength; i < level2.length; i++)
                if (Integer.parseInt(level2[i]) > 0)
                    return -1;
        }
        return 0;
    }
}

public class CompareVersionNumbers {
    public static void main(String[] args) {
        Solution s = new Solution();
        String [][] test_cases = {
            {"0.1", "1.1"},
            {"1.0.1", "1"},
            {"7.5.2.4", "7.5.3"},
            {"1.01", "1.001"},
            {"1", "1.1"}
        };
        for (String [] test_case : test_cases) {
            System.out.println(s.compareVersion(test_case[0], test_case[1]));
        }
    }
}
