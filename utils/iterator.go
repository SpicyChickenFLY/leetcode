package utils

// Iterator is a data structure interface
type Iterator interface {
	HasNext() bool
	Next() int
	Peek() int
}

// MyIterator my inplemention for Iterator
type MyIterator struct {
	nums  []int
	index int
}

// NewMyIterator initialize MyIterator here
func NewMyIterator(nums []int) *MyIterator {
	return &MyIterator{nums, 0}
}

// HasNext check elements exists
func (mi *MyIterator) HasNext() bool {
	return !(mi.index == len(mi.nums))
}

// Next return current element value and jump to next
func (mi *MyIterator) Next() int {
	val := mi.nums[mi.index]
	mi.index++
	return val
}

// Peek return current element value
func (mi *MyIterator) Peek() int {
	val := mi.nums[mi.index]
	mi.index++
	return val
}
