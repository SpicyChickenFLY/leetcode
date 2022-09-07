import java.util.HashSet;
import java.util.Set;

class Solution {
    // My solution combine bit manipulation and sum - Run time: 31.56% , Memory Usage: 98.81%
    public int[] singleNumber1(int[] nums) {
        if (nums.length == 2)
            return nums;
        Set<Integer> set = new HashSet<>();
        int sum = 0, xor = 0;
        for (int num : nums) {
            sum += num;
            xor ^= num;
            set.add(num);
        }
        int new_sum = 0;
        for (int num : set)
            new_sum += num * 2;
        // System.out.print(new_sum + ":" + sum + "\n");
        sum = new_sum - sum;
        for (int num : set) {
            int otherNum = sum - num;
            if (set.contains(otherNum) && (num ^ (otherNum)) == xor) {
                return new int[] {num, otherNum};
            }
        }
        return new int[2];
    }

    // inspired by others only use bit manipulation - Run time: 99.95% , Memory Usage: 96.43%
    public int[] singleNumber(int[] nums) {
        int totalxor = 0;
        for (int num : nums)
            totalxor ^= num;
        // mask the final 1 bit of xor result
        int mask = totalxor & ~(totalxor - 1);
        int xor1 = 0, xor2 = 0;
        for (int num : nums) {
            if ((num & mask) > 0)
                xor1 ^= num;
            else
                xor2 ^= num;
        }
        return new int[] {xor1, xor2};
    }
}

public class SingleNumberIII {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] input = { 2,1,2,3,4,1 };
        int[] result = s.singleNumber(input);
        for (int var : result) {
            System.out.print(var + " ");
        }
    }
}