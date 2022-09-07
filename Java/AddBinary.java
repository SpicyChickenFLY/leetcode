class Solution {
    public String addBinary(String a, String b) {
        String result = "";
        int length = a.length() < b.length() ? a.length() : b.length();
        boolean carry = false;
        for (int l = 1; l < length + 1; l++) {
            if (a.charAt(a.length() - l) == b.charAt(b.length() - l)) {
                result = (carry ? "1" : "0") + result;
                carry = (a.charAt(a.length() - l) == '1');
            } else {
                result = (carry ? "0" : "1") + result;
            }
        }
        if (a.length() < b.length())
            a = b;

        for (int l = a.length() - length - 1; l >= 0; l--) {
            if (a.charAt(l) == '1') {
                result = (carry ? "0" : "1") + result;
            } else {
                result = (carry ? "1" : "0") + result;
                carry = false;
            }
        }
        if (carry)
            result = "1" + result;
        return result;
    }
}

public class AddBinary {
    public static void main(String[] args) {
        Solution s = new Solution();
        String inputs[][] = { 
            { "11", "1" }, 
            { "1010", "1011" }, 
            { "100", "110010" } 
        };
        for (String input[] : inputs) {
            System.out.println(s.addBinary(input[0], input[1]));
        }
    }
}