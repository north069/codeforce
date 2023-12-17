package main

import (
	"bufio"
	"fmt"
	"os"
)

// 10 5 10 9 6 9 8 9 5
func main() {
	var T int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &T)
	res := make([]int, T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Fscan(in, &n)
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscan(in, &nums[i])
		}
		inc := make([]int, n)
		for i := n - 2; i >= 0; i-- {
			if nums[i] < nums[i+1] {
				inc[i] = inc[i+1]
			} else {
				inc[i] = inc[i+1] + 1
			}
		}
		res[i] = inc[0]
		pre := 1
		for j := 1; j < n; j++ {
			res[i] = min(res[i], pre+inc[j])
			if nums[j] >= nums[j-1] {
				pre++
			}
		}
		res[i] = min(res[i], pre)
	}
	for _, v := range res {
		fmt.Println(v)
	}
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
