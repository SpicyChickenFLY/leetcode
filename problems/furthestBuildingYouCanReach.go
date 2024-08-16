package problems

import "container/heap"
import "github.com/SpicyChickenFLY/go-leetcode/utils"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	ladderUsers := &utils.IntHeap{}
	heap.Init(ladderUsers)

	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]

		// 下楼免费, 否则上楼重新计算成本
		if diff < 0 {
			continue
		}

		// 梯子有空余肯定先用梯子
		heap.Push(ladderUsers, diff)

        // 梯子不够用把不配用梯子的踢出来
		if ladderUsers.Len() > ladders {
			replaced := heap.Pop(ladderUsers).(int)

			// 剩余砖块不足以跨越这栋楼,就结束了
			if bricks < replaced {
				return i
			}

			// 被踢出的就用砖块替代
			bricks -= replaced
		}

	}
	return len(heights) - 1
}
