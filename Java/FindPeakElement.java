
class Solution {
    // RunTime: 100.00%, MemoryUsage: 100.00%
    public int findPeakElement(int[] nums) {
        switch (nums.length) {
            case 0:
                return -1;
            case 1:
                return 0;
            default:
                int result = 0;
                boolean rise_flag = true;
                for (int i = 0; i < nums.length - 1; i++)
                    if (rise_flag && nums[i + 1] < nums[i])
                        return i;
                if (rise_flag)
                    return nums.length - 1;
                else
                    return -1;
        }
    }
}

public class FindPeakElement{
    public static void main(String[] args) {
        Solution s = new Solution();
        int [][]test_cases = {
            {1, 2, 3, 1},
            {1, 2, 1, 3, 5, 6, 4}
        };
        for (int []test_case : test_cases)
            System.out.println(s.findPeakElement(test_case));
    }
}