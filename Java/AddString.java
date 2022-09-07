class Solution {
    public String addStrings(String num1, String num2) {
        String result = "";
        int length = num1.length() < num2.length() ? num1.length() : num2.length();
        int carry = 0;
        for (int l = 1; l < length + 1; l++) {
            int sum = num1.charAt(num1.length() - l) + num2.charAt(num2.length() - l) - 96 + carry;
            carry = sum > 9 ? 1 : 0;
            result = sum % 10 + result;
        }

        if (num1.length() < num2.length())
            num1 = num2;

        for (int l = num1.length() - length - 1; l >= 0; l--) {
            if (num1.charAt(l) == '9' && carry == 1) {
                result = "0" + result;
                carry = 1;
            } else {
                result = (char) (num1.charAt(l) + carry) + result;
                carry = 0;
            }
        }
        if (carry == 1)
            result = "1" + result;
        return result;
    }
}

public class AddString {
    public static void main(String[] args) {
        Solution s = new Solution();
        String inputs[][] = { 
            { "11", "1" }, 
            { "1009", "1" }, 
            { "990", "110" } 
        };
        for (String input[] : inputs) {
            System.out.println(s.addString(input[0], input[1]));
        }
    }
}