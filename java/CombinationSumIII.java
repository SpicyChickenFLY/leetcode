import java.util.*;

class Solution {
    //sort solution - Run time: 100.00%, Memory Usage: 8.00%
    public List<List<Integer>> sub(int min, int k, int n) {
        List<List<Integer>> results = new ArrayList<>();
        
        if (min > n)
            return results;
        
        if (k == 1) {
            if (n > 9)
                return results;
            List<Integer> temp = new ArrayList<>();
            temp.add(n); 
            results.add(temp);
        }else {
            for (int i = min; i < 10; i++) {
                List<List<Integer>> sub_results = sub(i + 1, k - 1, n - i);
                if (sub_results.size() == 0)
                    continue;
                else {
                    for (List<Integer> sub_result :  sub_results) {
                        sub_result.add(0, i);
                        results.add(sub_result);
                    }
                }
            }
        }

        return results;
    }

    public List<List<Integer>> combinationSum3(int k, int n) {
        return sub(1, k, n);
    }
}

public class CombinationSumIII {
    public static void main(String[] args) {
        Solution s = new Solution();
        for (int k = 3; k < 5; k++) {
            for (int n = 3; n <= 30; n ++) {
                System.out.println(k + " - " + n);
                List<List<Integer>> results = s.combinationSum3(k, n);
                for (List<Integer> result : results) {
                    System.out.print("[");
                    for (int num : result) {
                        System.out.print(num + " ");
                    }
                    System.out.println("]");
                }
            }
        }    
    }
}