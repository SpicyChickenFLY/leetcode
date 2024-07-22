package problems

var badVersion int = -1

func isBadVersion(version int) bool {
	return version >= badVersion
}

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	firstBad := n
	firstGood := 0

	for firstBad > firstGood+1 {
		mid := (firstBad + firstGood) / 2
		if isBadVersion(mid) {
			firstBad = mid
		} else {
			firstGood = mid
		}
	}

	return firstBad
}
