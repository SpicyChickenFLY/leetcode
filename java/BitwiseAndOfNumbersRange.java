class Solution {
    // Runtime: 100.00%, Memory Usage: 10.00%
    public int rangeBitwiseAnd(int m, int n) {
        while (n > m) {
            n &= (n - 1);
        }
        return n;
    }   
}

public class BitwiseAndOfNumbersRange{
    public static void main(String[] args) {
        Solution s = new Solution();
        int [][]testCases = {
            {0, 1},
            {5, 7},
            {5, 8},
            {5, 9},
            {2147483647, 20000}
        };
        for (int[] testCase : testCases) {
            System.out.println(testCase[0] +"-"+ testCase[1] +": "+ s.rangeBitwiseAnd(testCase[0], testCase[1])); 
        }
    }
}