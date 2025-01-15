package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 将数组调整为最大堆的层级遍历形式
func maxHeapify(heap []int, start, end int) {
	// 假设左子节点是最大
	child_max := start*2 + 1
	if child_max > end {
		// 没有子节点，不需要调整
		return
	}
	if child_max+1 <= end && heap[child_max+1] > heap[child_max] {
		// 有右子节点,且比左节点大
		child_max = child_max + 1
	}
	if heap[start] > heap[child_max] {
		// 父节点已经比子节点大，不需要调整
		return
	}
	// 父节点比最大的子节点小，交换两值
	heap[start], heap[child_max] = heap[child_max], heap[start]
	// 继续调整
	maxHeapify(heap, child_max, end)
}

func heapSort(arr []int) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		// 假设数组为堆的层级遍历形式，长度的一半就是最右侧的非叶子节点
		maxHeapify(arr, i, length-1)
	}
}

// TestMaxHeapify if parent node can be swap by max child node
func TestMaxHeapify(t *testing.T) {
	type testCase struct {
		heap   []int
		start  int
		end    int
		expect []int
	}
	testCases := []testCase{
		{
			heap:   []int{4, 1, 3, 4, 6, 7},
			start:  1,
			end:    5,
			expect: []int{4, 1, 3, 4, 6, 7},
		},
	}

	for index, testCase := range testCases {
		maxHeapify(testCase.heap, testCase.start, testCase.end)
		for i := len(testCase.expect); i >= 0; i-- {
			if testCase.heap[i] != testCase.expect[i] {
				assert.Fail(t, fmt.Sprintf("第%d个案例结果不一致", index+1))
			}
		}
	}
}
