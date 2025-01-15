package main

import "fmt"

type ExamRoom struct {
	capacity int
	student  []int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		capacity: N,
		student:  []int{},
	}
}

func (this *ExamRoom) Seat() int {
	if len(this.student) == 0 {
		this.student = append(this.student, 0)
		return 0
	}
	// 全切片遍历
	seat, distance, index := 0, this.student[0], 0
	for i := 0; i < len(this.student)-1; i++ {
		currentDistance := (this.student[i+1] - this.student[i]) / 2
		if currentDistance > distance {
			distance = currentDistance
			seat = this.student[i] + distance
			index = i + 1
		}
	}
	if this.capacity-1-distance > this.student[len(this.student)-1] {
		seat = this.capacity - 1
		index = len(this.student)
	}
	// 切片中index位置插入新元素
	this.student = append(this.student, 0)
	copy(this.student[index+1:], this.student[index:])
	this.student[index] = seat
	return seat
}

func (this *ExamRoom) Leave(p int) {
	l, r := 0, len(this.student)-1
	for l <= r {
		m := (l + r) >> 1
		if this.student[m] == p {
			// remove student[m]
			this.student = append(this.student[:m], this.student[m+1:]...)
			return
		}
		if this.student[m] > p {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return
}

func main() {
	// Your ExamRoom object will be instantiated and called as such:
	N := 10
	obj := Constructor(N)
	fmt.Println(obj.Seat())
	obj.Leave(4)

}
