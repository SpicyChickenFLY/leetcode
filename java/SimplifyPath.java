import java.util.*;

class Solution {
    public String simplifyPath(String path) {
        String result = "";
        Stack<String> stack = new Stack<String>();
        String sequence[] = path.split("/");
        for (String  element : sequence) {
            if (element.equals(".") || element.equals("")) 
                continue;
            if (element.equals("..")) {
                if (!stack.empty()) 
                    stack.pop();
                continue;
            }
            stack.push(element);
        }
        if (stack.empty())
            return "/";
        while(!stack.empty())
            result = stack.pop() + "/" + result;
        return "/" + result.substring(0, result.length() - 1);
    }
}

public class SimplifyPath {
    public static void main(String[] args) {
        Solution s = new Solution();
        String inputs[] = {
            "/home/",
            "/../",
            "/home//foo/",
            "/a/./b/../../c/",
            "/a/../../b/../c//.//",
            "/a//b////c/d//././/.."
        };
        for (String  input : inputs) {
            System.out.println(s.simplifyPath(input));
        }
    }
}