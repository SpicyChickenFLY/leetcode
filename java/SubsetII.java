import java.util.*;

class Solution {
    // Runtime: 12.97%, Memory Usage: 98.53%
    public List<List<Integer>> subsetsWithDup(int[] nums) {
        Arrays.sort(nums);
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
        HashSet h = new HashSet(results);
        List<List<Integer>> newResults = new ArrayList<>();
        newResults.addAll(h);
        return newResults;
    }
}