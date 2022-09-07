class Solution {
    public int lengthOfLastWord(String s) {
        int length = 0;
        for (int i = s.length() - 1; i >= 0; i--) {
            if (s.charAt(i) == 32) {
                if (length != 0) {
                    return length;
                }else {
                    continue;
                }
            }else {
                length++;
            }
        }
        return length;
    }
}

public class LengthOfLastWord {
    public static void main(String[] args) {
        Solution s = new Solution();
        String inputs[] = { "hello world", "ni hao", "en ", "a", " ", "" };
        for (String str : inputs) {
            System.out.println("input: " + str);
            System.out.println("output: " + s.lengthOfLastWord(str));
        }
    }
}
