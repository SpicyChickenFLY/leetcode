package golang

func jump(nums []int) int {
	last := len(nums) - 1 // 记录最后一个位置
	if last == 0 {
		return 0
	}
	var jumps int // 记录总跳数
	var goal int // 记从上一跳终点到目前位置, 可以达到的最远位置
	var prevGoal int // 记录上一跳最远到达位置

	for pos, step := range nums {
		if pos + step > goal { // 可达最远位置可以刷新
			goal = pos + pos
		}

		if pos == prevGoal { // 现在就在上一跳能最远到达位置

			jumps++ // 记录上一次跳的次数
			if goal >= last { // 已经到达终点
				break
			}
			prevGoal = goal // 记录上次跳完后, 这次能达到的最远位置
		}
	}

	return jumps
}
