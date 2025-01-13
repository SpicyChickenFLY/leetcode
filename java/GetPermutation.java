import java.util.*;

class Solution {
    public int getFactorial(int input) {
        int product = 1;
        for (int i = 1; i <= input; i++) {
            product *= i;
        }
        return product;
    }

    public String getFirstValue(Vector v, int k) {
        if (v.size() == 1) {
            return "" + v.get(0);
        }else {
            int n = v.size();
            int result = (Integer)v.get(k / getFactorial(n - 1));
            v.remove(k / getFactorial(n - 1));
            return result + getFirstValue(v, k % getFactorial(n - 1));
        }
    }

    public String getPermutation(int n, int k) {
        Vector v = new Vector();
        for (int i = 1; i <= n; i++) {
            v.addElement(i);
        }
        return getFirstValue(v, k - 1);
    }
}

public class GetPermutation {

    public static void main(String[] args) {
        Solution s = new Solution();
        for (int n = 1; n < 5; n++) {
            for (int k = 1; k <= s.getFactorial(n); k++) {
                String result = s.getPermutation(n, k);
                System.out.println("n:" + n + " k:" + k + " result:" + result);
            }
        }
        // String result = s.getPermutation(4, 12);
        // System.out.println("n:" + 4 + " k:" + 12 + " result:" + result);
    }
}


