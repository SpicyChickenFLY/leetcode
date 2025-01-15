package main

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(intervals); i++ {
		// 无交集且比新的小, 直接加
		if intervals[i][1] < newInterval[0] {
			result = append(result, intervals[i])

        // 第一次无交集且比新的大
        // 新的还没加, 加一下
        // 后边intervals都比新的大, 直接加了返回
		} else if intervals[i][0] > newInterval[1] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)

        // 有交集, 合并进新的
		} else {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
		}
	}

	return append(result, newInterval)
}
