package main
import (
	"fmt"
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0;
	}

	result := nums[0] + nums[1] + nums[2]
	if result == target {
		return target
	}
	min_diff := math.Abs(float64(result - target))

	sort.Ints(nums)

	for first := 0; first < len(nums) - 2; first++ {
		second := first + 1
		third := len(nums) - 1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return target
			}
			diff := math.Abs(float64(sum - target))
			if diff < min_diff {
				min_diff = diff
				result = sum
			}
			if sum < target {
				second += 1
			} else if sum > target {
				third -= 1
			}
		}
	}
	return result
}

func main() {
	testCases := [][][]int {
		{{-1 ,2, 1, -4}, {1}},
		{{1,8,6,2,5,4,8,3,7}, {4}},
	}
	
	for _, testCase := range testCases {
		nums := testCase[0]
		target := testCase[1][0]
		result := threeSumClosest(nums, target)
		fmt.Printf("%v,%d,%d\n", nums, target, result)
	}
}