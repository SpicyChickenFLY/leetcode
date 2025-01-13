package golang

func findMaxLength(nums []int) int {
	sumRec := map[int]int{0: -1}
	maxLen := 0
	sum := 0

	for i, num := range nums {

		// 计算从头到这的相对和值
		if num == 0 {
			sum--
		} else {
			sum++
		}

		// 如果找到map中有相同相对值
		// 说明从这个位置到这有相同数目的0 和 1
		if _, find := sumRec[sum]; find {
			maxLen = max(maxLen, i-sumRec[sum])
            // 不记录map, 否则就不是最长的子序列了

			// 找不到说明目前为止不存在一个位置到这, 记录map
		} else {
			sumRec[sum] = i
		}
	}

	return maxLen
}
