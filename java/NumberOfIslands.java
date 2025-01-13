import java.util.*;

class Solution {

    public void seedFill(char[][] board, int h, int w, int i, int j) {
        if (i < 0 || i >= h || j < 0 || j >= w)
            return;
        if (board[i][j] == '1') {
            board[i][j] = '#';
            seedFill(board, h, w, i - 1, j);
            seedFill(board, h, w, i + 1, j);
            seedFill(board, h, w, i, j - 1);
            seedFill(board, h, w, i, j + 1);
        }
    }

    // Runtime: 100.00%, Memory Usage: 58.61%
    public int numIslands(char[][] grid) {
        if (grid.length > 0) {
            if (grid[0].length < 0)
                return 0;
        } else {
            return 0;
        }

        int h = grid.length, w = grid[0].length;
        
        int count = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == '1') {
                    count++;
                    seedFill(grid, h, w, i, j);
                }
            }
        }

        return count;
    }
}

public class NumberOfIslands {
    public static void main(String[] args) {
        Solution s = new Solution();

        char[][] testCase1 = { { '1', '1', '1', '1', '0' }, { '1', '1', '0', '1', '0' }, { '1', '1', '0', '0', '0' }, { '0', '0', '0', '0', '0' } };
        System.out.println(s.numIslands(testCase1));

        char[][] testCase2 = { { '1', '1', '0', '0', '0' }, { '1', '1', '0', '0', '0' }, { '0', '0', '1', '0', '0' }, { '0', '0', '0', '1', '1' } };
        System.out.println(s.numIslands(testCase2));

        char[][] testCase3 = { { '1', '1', '1' }, { '0', '1', '0' }, { '1', '1', '1' } };
        System.out.println(s.numIslands(testCase3));

        char[][] testCase4 = {{'1','0','1','1','1'},{'1','0','1','0','1'},{'1','1','1','0','1'}};
        System.out.println(s.numIslands(testCase4));

        char[][] testCase5 = { {'1','1','1','1','1','1','1'}, {'0','0','0','0','0','0','1'}, {'1','1','1','1','1','0','1'}, {'1','0','0','0','1','0','1'}, {'1','0','1','0','1','0','1'}, {'1','0','1','1','1','0','1'}, {'1','1','1','1','1','1','1'}};
        System.out.println(s.numIslands(testCase5));
    }
}
