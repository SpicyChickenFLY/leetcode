class Solution {
    public int hammingWeight(int n) {
        int count = 0;
        while (n != 0) { // judge if there is a one bit left
            n = n & (n - 1); // remove the las one bit in n
            count ++;
        }
        return count;
    }
}

public class NumberOfOneBits {
    public static void main(String[] args) {
        Solution s = new Solution();
        for (int i = 0; i < 10; i++) {
            System.out.println(i +  ";" + s.hammingWeight(i));
        }
    }
}