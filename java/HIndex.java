import java.util.*;

class Solution {
    public int hIndex(int[] citations) {
        Arrays.sort(citations);
        for(int i = 0; i < citations.length; i++)
            if (citations[i] >= citations.length - i)
                return citations.length - i;
        return 0;
    }
}

public class HIndex {
    public static void main(String[] args) {
        Solution s = new Solution();
        int inputs[][] = {
            {3,0,6,1,5}
        };
        for (int input[] : inputs) {
            System.out.println(s.hIndex(input));
        }
    }
}