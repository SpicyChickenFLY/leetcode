class Solution {
    private int record[][];

    public int uniquePathsRecursion(int m, int n, int[][] obstacleGrid) {
        if (record[m - 1][n - 1] != 0)
            return record[m - 1][n - 1];
        if (obstacleGrid[m - 1][n - 1] == 1) {
            record[m - 1][n - 1] = 0;
            return 0;
        }
        if (m == 1) {
            int result = uniquePathsRecursion(m, n - 1, obstacleGrid);
            record[m - 1][n - 1] = result;
            return result;
        }
        if (n == 1) {
            int result = uniquePathsRecursion(m - 1, n, obstacleGrid);
            record[m - 1][n - 1] = result;
            return result;
        }
        int result = uniquePathsRecursion(m - 1, n, obstacleGrid)
            + uniquePathsRecursion(m, n - 1, obstacleGrid);
        record[m - 1][n - 1] = result;
        return result;
    }

    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        if (obstacleGrid[0][0] == 1) {
            return 0;
        }
        int m = obstacleGrid.length;
        int n = obstacleGrid[0].length;
        if (m == 1 || n == 1) {
            for (int y = 0; y < m; y++) {
                for (int x = 0; x < n; x++) {
                    if (obstacleGrid[y][x] == 1) {
                        return 0;
                    }
                }
            }
            return 1;
        }
        record = new int[m][n];
        record[0][0] = 1;
        record[1][0] = obstacleGrid[1][0] == 1 ?  0 : 1;
        record[0][1] = obstacleGrid[0][1] == 1 ?  0 : 1;
        return uniquePathsRecursion(m, n, obstacleGrid);
    }
}

public class UniquePathsII {
    public static void main(String[] args) {
        Solution s = new Solution();
        int input[][] = { { 0, 0, 0 }, { 0, 1, 0 }, { 0, 0, 0 } };
        // int input[][] = { { 0, 0}, { 1, 0 } };
        // int input[][] = { { 1, 0 }, { 0, 0 } };
        // int input[][] = { { 0, 0 }, { 1, 1 }, { 0, 0 } };
        int result = s.uniquePathsWithObstacles(input);
        System.out.println(result);
    }
}