import java.util.*;

class Other {
    public void printL(List<Integer> result) {
        for (int num : result) {
            System.out.print(num);
            System.out.print(" ");
        }
    }

    public void printLL(List<List<Integer>> results) {
        for (List<Integer> result : results) {
            printL(result);
            System.out.println("");
        }
    }
}

class Solution {
    // Runtime: 100.00%, Memory Usage: 99.18%
    public List<List<Integer>> subsets(int[] nums) {
        // Other o = new Other();
        List<List<Integer>> results = new ArrayList<>();
        List<Integer> blank = new ArrayList<Integer>();
        results.add(blank);
        for (int num : nums) {
            int size = results.size();
            for (int i = 0; i < size; i++) {
                List<Integer> result = new ArrayList<>(results.get(i));
                result.add(num);
                results.add(result);
            }
            // o.printLL(results);
        }
        return results;
    }
}

public class Subsets {
    

    public static void main(String[] args) {
        Solution s = new Solution();
        Other o = new Other();
        int[] nums = {1, 2, 3};
        List<List<Integer>> results = s.subsets(nums);
        o.printLL(results);
    }
}