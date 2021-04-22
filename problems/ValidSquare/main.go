package main

import "sort"

func calPowerDistance(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	powerDistances := []int{
		calPowerDistance(p1, p2),
		calPowerDistance(p1, p3),
		calPowerDistance(p1, p4),
		calPowerDistance(p2, p3),
		calPowerDistance(p2, p4),
		calPowerDistance(p3, p4),
	}
	sort.Ints(powerDistances)
	return powerDistances[0] > 0 &&
		powerDistances[0] == powerDistances[3] &&
		powerDistances[4] == powerDistances[5]
}

func main() {

}
