import (
	"fmt"
	"sort"
)

func diffWaysToCompute(expression string) []int {
	return recursive(parser(expression))
}

func recursive(nums []int, ops []string) []int {
	if len(nums) == 1 {
		return nums
	}
	if len(nums) == 2 {
		return []int{calu(nums[0], nums[1], ops[0])}
	}

	ans := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		lN := nums[0:i]
		lOp := ops[0 : i-1]
		rN := nums[i:len(nums)]
		rOp := ops[i:len(ops)]

		lA := recursive(lN, lOp)
		rA := recursive(rN, rOp)
		ans = append(ans, caluMultiple(lA, rA, ops[i-1])...)
	}
	return ans
}

func caluMultiple(n1, n2 []int, op string) []int {
	ans := make([]int, 0, len(n1)*len(n2))
	for i := range n1 {
		for j := range n2 {
			ans = append(ans, calu(n1[i], n2[j], op))
		}
	}
	return ans
}

func calu(n1, n2 int, op string) int {
	switch op {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	}
	return 0
}

func parser(expression string) ([]int, []string) {
	ns := strings.FieldsFunc(expression, func(r rune) bool {
		return r == '+' || r == '*' || r == '-'
	})
	var nums = make([]int, len(ns))
	for i := range nums {
		nums[i], _ = strconv.Atoi(ns[i])
	}
	ops := strings.FieldsFunc(expression, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	return nums, ops
}