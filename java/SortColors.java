class Solution {
    public void sortColors(int[] nums) {
        int start = 0, end = nums.length;
        for (int var : nums) {
            switch (var) {
            case 0:
                nums[start++] = 0;
                break;
            case 2:
                end--;
            }
        }
        for (int i = start; i < end; i++) {
            nums[i] = 1;
        }
        for (int i = end; i < nums.length; i++) {
            nums[i] = 2;
        }
        for (int var : nums) {
            System.out.print(var);
        }
    }
}

public class SortColors {
    public static void main(String[] args) {
        Solution s = new Solution();
        int inputs[][] = { { 2, 0, 2, 1, 1, 0 }, };
        for (int input[] : inputs) {
            s.sortColors(input);
        }
    }
}