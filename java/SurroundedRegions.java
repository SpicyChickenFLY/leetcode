import java.util.Vector;

class Solution {
    public void seedFill(char[][] board, int h, int w, int i, int j) {
        if (i < 0 || i >= h || j < 0 || j >= w)
            return;
        if (board[i][j] == 'O') {
            board[i][j] = '#';
            seedFill(board, h, w, i-1, j);
            seedFill(board, h, w, i+1, j);
            seedFill(board, h, w, i, j-1);
            seedFill(board, h, w, i, j+1);
        }
    }

    // Runtime: 51.21%, Memory Usage: 85.71%
    public void solve(char[][] board) {
        if (board.length <= 1 || board[0].length <= 1) {
            return;
        }
        int height = board.length, width = board[0].length;
        for (int i = 0; i < height; i++) {
            for (int j = 0; j < width; j++) {
                if (i != 0 && i != height - 1 && j != 0 && j != width - 1) {
                    continue;
                }
                if (board[i][j] == 'O') {
                    seedFill(board, height, width, i, j);
                }
            }
        }

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (board[i][j] == '#') {
                    board[i][j] = 'O';
                } else if(board[i][j] == 'O') {
                    board[i][j] = 'X';
                }
            }
        }
    }
}

public class SurroundedRegions {
    public static void main(String[] args) {
        Solution s = new Solution();
        char[][] testCase = {
            {'X', 'X', 'X', 'X'},
            {'X', 'O', 'O', 'X'},
            {'X', 'X', 'O', 'X'},
            {'X', 'O', 'X', 'X'}
        };
        s.solve(testCase);
        for (int i = 0; i < testCase.length; i++) {
            for (int j = 0; j < testCase[0].length; j++) {
                System.out.print(testCase[i][j] + " ");
            }
            System.out.print("\n");
        }
    }
}