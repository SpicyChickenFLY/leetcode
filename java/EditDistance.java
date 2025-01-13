class Solution {
    // Runtime: 37.53%, Memory Usage: 100.00%
    public int minDistance(String word1, String word2) {
        if (word1.length() < word2.length()) {
            String temp = word1;
            word1 = word2;
            word2 = temp;
        }
        int[][] table = new int[word1.length() + 1][word2.length() + 1];
        for (int i = 0; i < word1.length() + 1; i++)
            table[i][0] = i;
        for (int j = 0; j < word2.length() + 1; j++)
            table[0][j] = j;
        table[0][0] = 0;
        for (int i = 1; i < word1.length() + 1; i++) {
            for (int j = 1; j < word2.length() + 1; j++) {
                int temp = 1;
                if (word1.charAt(i - 1) == word2.charAt(j - 1))
                    temp = 0;
                int min = table[i - 1][j - 1] + temp;
                if (table[i - 1][j] < min)
                    min = table[i - 1][j] + 1;
                if (table[i][j - 1] < min)
                    min = table[i][j - 1] + 1;
                table[i][j] = min;
                // System.out.print(min + " ");
            }
            // System.out.println();
        }

        return table[word1.length()][word2.length()];
    }
}

public class EditDistance {

    public static void main(String[] args) {
        String[][] test_cases = { { "horse", "ros" }, { "intention", "execution" }, {"", "a"}, {"ab", "bc"} };
        Solution s = new Solution();
        for (String[] test_case : test_cases) {
            System.out.println(s.minDistance(test_case[0], test_case[1]));
        }

    }
}