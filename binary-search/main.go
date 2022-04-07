package main

import (
	"fmt"
	"sort"
)

func Contains(a []int, target int) bool {
	if !sort.IntsAreSorted(a) {
		sort.Ints(a)
	}

	left := 0
	right := len(a) - 1
	mid := 0
	for {
		fmt.Printf("start process left %d, right %d, target=%d\n", left, right, target)
		if right < left {
			return false
		}

		mid = left + (right-left)/2
		if a[mid] == target {
			return true
		}

		if a[mid] > target {
			right = mid - 1
			continue
		}

		left = mid + 1
	}
}

// return first index equal target if notfound return -1
func FindFirstIndexEqualTarget(a []int, target int) int {
	if !sort.IntsAreSorted(a) {
		sort.Ints(a)
	}

	ans := -1
	l := 0
	r := len(a) - 1
	for {
		fmt.Printf("start process left %d, right %d, target=%d\n", l, r, target)
		if l > r {
			break
		}

		mid := l + (r-l)/2
		if a[mid] > target {
			r = mid - 1
			continue
		}

		if a[mid] < target {
			l = mid + 1
			continue
		}

		ans = mid
		r = mid - 1
	}
	return ans
}

func FindLastIndexEqualTarget(a []int, target int) int {
	ans := -1
	l := 0
	r := len(a) - 1

	for {
		fmt.Printf("start process left %d, right %d, target=%d\n", l, r, target)
		if l > r {
			break
		}

		mid := l + (r-l)/2
		if a[mid] > target {
			r = mid - 1
			continue
		}

		if a[mid] < target {
			l = mid + 1
			continue
		}

		ans = mid
		l = mid + 1
	}

	return ans
}
func main() {
	//fmt.Println(Contains([]int{3, 4, 6, 17, 28, 39, 40, 54, 55, 78, 92}, 1))
	//fmt.Println(FindFirstIndexEqualTarget([]int{3, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 17, 28, 39, 40, 54, 55, 78, 92}, 5))
	fmt.Println(FindLastIndexEqualTarget([]int{3, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 17, 28, 39, 40, 54, 55, 78, 92}, 5))
}
