

class Solution {
    public int[] plusOne(int[] digits) {
        int carry = 1;
        for (int i = digits.length - 1; i >= 0; i--) {
            if (digits[i] == 9 && carry == 1) {
                digits[i] = 0;
            }else if (carry == 1) {
                digits[i] += 1;
                carry = 0;
            }
        }
        if (carry == 1) {
            int result[] = new int[digits.length + 1];
            result[0] = 1;
            return result;
        }
        return digits;
    }
}

public class PlusOne {
    public static void main(String[] args) {
        Solution s = new Solution();
        int input[] = {9};
        int result[] = s.plusOne(input);
        for (int digit : result) {
            System.out.print(digit);
        }
    }
}