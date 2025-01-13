import java.util.*;

class Solution {
    // Runtime: 79.87%, Memory Usage: 98.00%
    public int evalRPN(String[] tokens) {
        Stack<Integer> operands = new Stack<Integer>();
        for (String token : tokens) {
            if (token.equals("+")) {
                int temp = operands.pop();
                operands.push(operands.pop() + temp);
            }
            else if (token.equals("-")) {
                int temp = operands.pop();
                operands.push(operands.pop() - temp);
            }
            else if (token.equals("*")) {
                int temp = operands.pop();
                operands.push(operands.pop() * temp);
            }
            else if (token.equals("/")) {
                int temp = operands.pop();
                operands.push(operands.pop() / temp);
            }
            else {
                operands.push(Integer.parseInt(token));
            }
        }
        return operands.pop();
    }
}

public class EvaluateReversePolishNotation {
    public static void main(String[] args) {
        Solution s = new Solution();
        String [][]testCases = {
            {"2", "1", "+", "3", "*"},
            {"4", "13", "5", "/", "+"},
            {"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
        };
        for (String []testCase : testCases) {
            System.out.println(testCase);
            System.out.println(s.evalRPN(testCase));
        }
    }
}