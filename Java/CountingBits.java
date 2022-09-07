class Solution {
    // Runtime: 99.65% Memory Usage: 5.05%
    public int[] countBits(int num) {
        int[] result = new int[num + 1];
        if (num < 1) {
            return result;
        }
        result[1] = 1;
        int level = 1;
        for (int i = 2; i <= num; i++) {
            if ((i & (i - 1)) == 0)
                level *= 2;
            //System.out.println(i + ";" + level);
            result[i] = result[i - level] + 1;
        }
        return result;
    }
}

public class CountingBits {
    public static void main(String[] args) {
        Solution s = new Solution();
        for (int i = 0; i < 20; i++) {
            int[] results = s.countBits(i);
            System.out.print("result: ");
            for (int result : results) {
                System.out.print(result + " ");
            }
            System.out.print("\n");
        }
    }
}