class Solution {
    //Linear Search, O(n)
    public int mySqrt1(int x) {
        int result = 1;
        while (result * result <= x) {
            result++;
        }
        if (result - 1 > 46339) {
            return 46340;
        }
        return result - 1;
    }
    //Binary Search,
    public int mySqrt(int x) {
        int result = 0;
        int start = 0, end = 46340;
        if (x > 2147395599) {
            return 46340;
        }
        while (true) {
            result = (start + end) / 2;
            if (result * result == x) {
                return result;
            } else if (result * result > x) {
                end = result;
            } else {
                if ((result + 1) * (result + 1) > x) {
                    return result;
                } else {
                    start = result;
                }
            }
        }
    }
}

public class Sqrtx {
    public static void main(String[] args) {
        Solution s = new Solution();

        for (int i = 2147395580; i < 2147395610; i++) {
            System.out.println(i + " " + s.mySqrt(i));
        }
    }
}