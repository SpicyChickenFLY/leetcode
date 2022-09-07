class Solution {
    private int record[][];

    public int uniquePathsRecursion(int m, int n) {
        if (record[m - 1][n - 1] != 0)
            return record[m - 1][n - 1];
        if (m == 1 || n == 1) {
            record[m - 1][n - 1] = 1;
            return 1;
        }
        else {
            int result = uniquePathsRecursion(m - 1, n) + 
                uniquePathsRecursion(m, n - 1); 
            record[m - 1][n - 1] = result;
            return result;
        }
    }

    public int uniquePaths(int m, int n) {
        if (m == 1 || n == 1) {
            return 1;
        }
        record = new int[m][n];
        record[0][0] = 1;
        record[1][0] = 1;
        record[0][1] = 1;
        return uniquePathsRecursion(m, n);
    }
} 

public class UniquePaths {
    public static void main(String[] args) {
        Solution s = new Solution();
        for (int m = 1; m < 10; m++) {
            for (int n = 1; n < 10; n++) {
                System.out.printf(
                    "m: %2d, n: %2d, result: %d\n", 
                    m, n, s.uniquePaths(m, n)
                );
            }
        }
    }
}