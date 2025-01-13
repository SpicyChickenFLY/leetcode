
class Solution {
    // Runtime: 100.00%, Memory Usage: 100.00%
    public void rotate(int[][] matrix) {
        int middleLength = matrix.length / 2 + 1, length = matrix.length;
        for (int cirleIndex = 0; cirleIndex < middleLength; cirleIndex++) {
            int reverseIndex = length - cirleIndex - 1;
            for (int i1 = cirleIndex; i1 < reverseIndex; i1++) {
                int temp = matrix[cirleIndex][i1];
                matrix[cirleIndex][i1] = matrix[length - i1 - 1][cirleIndex];
                matrix[length - i1 - 1][cirleIndex] = matrix[reverseIndex][length - i1 - 1];
                matrix[reverseIndex][length - i1 - 1] = matrix[i1][reverseIndex];
                matrix[i1][reverseIndex] = temp;
            }
        }
    }
}

public class RotateImage {
    public static void main(String[] args) {
        Solution s = new Solution();

        int[][] testCase1 = { { 1, 2, 3 }, { 4, 5, 6 }, { 7, 8, 9 } };
        s.rotate(testCase1);
        for (int[] line : testCase1) {
            for (int num : line) {
                System.out.print(num + " ");
            }
            System.out.print('\n');
        }

        int[][] testCase2 = { { 5, 1, 9, 11 }, { 2, 4, 8, 10 }, { 13, 3, 6, 7 }, { 15, 14, 12, 16 } };
        s.rotate(testCase2);
        for (int[] line : testCase2) {
            for (int num : line) {
                System.out.print(num + " ");
            }
            System.out.print('\n');
        }
    }
}