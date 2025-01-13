class Solution {
    private int record[][];
    private int paths[][];

    public int mininumPathsRecursion(int m, int n) {
        if (m < 1 || n < 1) {
            return Integer.MAX_VALUE;
        }
        if (record[m - 1][n - 1] != -1)
            return record[m - 1][n - 1];
        int path_up = mininumPathsRecursion(m - 1, n);
        int path_left = mininumPathsRecursion(m, n - 1);
        int result =  paths[m - 1][n - 1] + (path_up > path_left ? path_left : path_up) ;
        record[m - 1][n - 1] = result;
        return result;

    }

    public int minPathSum(int[][] grid) {
        paths = grid;
        int m = paths.length;
        int n = paths[0].length;
        if (m == 1) {
            int result = 0;
            for (int x = 0; x < n; x++) {
                result += paths[0][x];
            }
            return result;
        }
        if (n == 1) {
            int result = 0;
            for (int y = 0; y < m; y++) {
                result += paths[y][0];
            }
            return result;
        }
        record = new int[m][n];
        for (int y = 0; y < m; y++) {
            for (int x = 0; x < n; x++) {
                record[y][x] = -1;
            }
        }
        record[0][0] = paths[0][0];
        return mininumPathsRecursion(m, n);
    }
}

public class MinimumPathSum {
    public static void main(String[] args) {
        Solution s = new Solution();
        int input[][] = {
            {1, 3, 1},
            {1, 5, 1},
            {4, 2, 1}
        };
        /* int input[][] = {
            {0, 0}, {0, 0}
        }; */

        int result = s.minPathSum(input);
        System.out.println(result);
    }
}