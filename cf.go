package main

import (
	"bufio"
	"fmt"
	"os"
)

// 10 5 10 9 6 9 8 9 5
func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	a, b := make([]int, n), make([]int, m)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	check := func(span int) bool {
		j := 0
		for _, v := range a {
			for j < m && !(v <= b[j]+span && v >= b[j]-span) {
				j++
			}
			if j == m {
				return false
			}
		}
		return true
	}

	l, r := 0, 2000000000
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
}

//aadba

//	func max(i, j int) int {
//		if i > j {
//			return i
//		} else {
//			return j
//		}
//	}
//func min(i, j int) int {
//	if i > j {
//		return j
//	}
//	return i
//}

//a := make([]int, n)
//for i := range a {
//	fmt.Fscan(in, &a[i])
//}

type neighbor struct {
	to, wt int
}
