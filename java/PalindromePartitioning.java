import java.util.*;

class Solution {
    public boolean isPalindromeValid(String s, int end) {
        int start = 0;
        while (start < end) {
            if (s.charAt(start) == s.charAt(end)) {
                start++;
                end--;
            } else {
                return false;
            }
        }
        return true;
    }

    // Inspired by other solution - RunTime: 21.49%, MemoryUsage: 95.45%
    public List<List<String>> partition(String s) {
        List<List<String>> results = new ArrayList<>();
        for (int i = 0; i < s.length(); i++) {
            if (isPalindromeValid(s, i)) {
                List<List<String>> subResults = partition(s.substring(i+1));
                if (subResults.size() > 0) {
                    for (List<String> subResult : subResults) {
                        subResult.add(0, s.substring(0, i + 1));
                        results.add(subResult);
                    }
                } else {
                    List<String> firstResult = new ArrayList<>();
                    firstResult.add(s.substring(0, i + 1));
                    results.add(firstResult);
                }
            } 
        }
        return results;
    }
}

public class PalindromePartitioning {
    public static void main(String[] args) {
        Solution s = new Solution();
        String []testCases = {
            "aab",
            "aaba",
            "aabbaa"
        };
        for (String testCase : testCases) {
            List<List<String>> results =  s.partition(testCase);
            for (List<String> result : results) {
                for (String str : result) {
                    System.out.print(str + " ");
                }
                System.out.println("\n");
            }
            System.out.println("\n");
        }
    }
}
