
class Solution {
    public int hIndex_traverse(int[] citations) {
        for (int i = 0; i < citations.length; i++) 
            if (citations[i] >= citations.length - i) 
                return citations.length - i;
        return 0;
    }

    public int hIndex_binary(int[] citations) {
        int high = citations.length - 1;
        int low = 0;
        int mid;
        while (low <= high) {
            mid = (high + low) / 2;
            if (citations[mid] >= citations.length - mid && (mid == 0 || citations[mid - 1] < citations.length - mid))
                return citations.length - mid;
            else if (citations[mid] < citations.length - mid)
                low = mid + 1;
            else
                high = mid - 1;
        }
        return citations.length - low;
    }
}

public class HIndexII {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[][] inputs = {
            {0, 1, 3, 5, 6}
        };
        for (int[] input : inputs) {
            System.out.println(s.hIndex_binary(input));
        }
    }
}