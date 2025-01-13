package main

type Iterator struct {
	nums  []int
	index int
}

func NewIterator(nums []int) *Iterator {
	return &Iterator{nums, 0}
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return !(this.index == len(this.nums))
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	val := this.nums[this.index]
	this.index++
	return val
}

type PeekingIterator struct {
	nums  []int
	index int
}

func Constructor(iter *Iterator) *PeekingIterator {
	nums := []int{}
	for iter.hasNext() {
		nums = append(nums, iter.next())
	}
	return &PeekingIterator{nums, 0}
}

func (this *PeekingIterator) hasNext() bool {
	return !(this.index == len(this.nums))
}

func (this *PeekingIterator) next() int {
	val := this.nums[this.index]
	this.index++
	return val
}

func (this *PeekingIterator) peek() int {
	return this.nums[this.index]
}
