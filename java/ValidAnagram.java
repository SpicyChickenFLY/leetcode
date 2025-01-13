import java.util.*;

class Solution {
    // Runtime: 39.57%, Memory Usage: 94.19%%
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length())
            return false;
        char []a = s.toCharArray();
        char []b = t.toCharArray();
        Arrays.sort(a);
        Arrays.sort(b);
        for (int index = 0; index < s.length(); index++) {
            if (a[index] != b[index])
                return false;
        }
        return true;
    }
}