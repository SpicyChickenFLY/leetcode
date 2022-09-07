class Solution {
    public int [][]dfa = {
        { 1, -1,  2,  9,  0},
        { 1,  5, -1,  3,  8},
        { 1, -1, -1, 10, -1},
        { 4,  5, -1, -1,  8},
        { 4,  5, -1, -1,  8},
        { 6, -1,  7, -1, -1},
        { 6, -1, -1, -1,  8},
        { 6, -1, -1, -1, -1},
        {-1, -1, -1, -1,  8},
        { 4, -1, -1, -1, -1},
        { 4, -1, -1, -1, -1}
    };

    // DFA solution - Run time: 77.74%, Memory Usage: 99.29%
    public boolean isNumber(String s) {
        int state = 0;
        char current;
        for (int i = 0; i < s.length(); i++) {
            
            current = s.charAt(i);
            switch (current) {
                case 'e':
                    state = dfa[state][1];
                    break;
                case '.':
                    state = dfa[state][3];
                    break;
                case ' ':
                    state = dfa[state][4];
                    break;
                default:
                    if (current >= '0' && current <= '9')
                        state = dfa[state][0];
                    else if (current == '+' || current == '-')
                        state = dfa[state][2];
                    else
                        return false;
            }
            System.out.print(state + "->");
            if (state == -1) 
                return false;
        }
        if (state == 0 || state == 2 || state == 5 || state == 7 || state == 9 || state == 10)   
            return false;
        return true;
    }
}

public class ValidNumber {
    public static void main(String[] args) {
        Solution s = new Solution();
        String []test_cases = {
            "0",             // => true
            " 0.1 ",         // => true
            " -90e3   ",     // => true
            "2e10",          // => true
            " 6e-1",         // => true
            "53.5e93",       // => true
            "3.",
            ".3",
            "+.8",
            "46.e3",

            ".",
            ". ",
            "abc",           // => false
            "1 a",           // => false
            " 1e",           // => false
            "e3",            // => false
            " 99e2.5 ",      // => false
            " --6 ",         // => false
            "-+3",           // => false
            "95a54e53",        // => false
            " -.",
            ".44.8",
            "6.3.0"
        };
        for (String test_case : test_cases) {
            System.out.println("\n" + test_case + ": " + s.isNumber(test_case));
        }
    }
}