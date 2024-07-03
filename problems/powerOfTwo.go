package problems

func isPowerOfTwo(n int) bool {
    if (n <= 0) {
        return false
    }
    // 2的幂拥有特点是,二进制形式均为1开头,若干0结尾
    for (n & 1 == 0) {
        n = n >> 1
    }
    return n == 1
}
