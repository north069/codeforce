package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		fmt.Fscan(in, &n)
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscan(in, &nums[i])
		}
		solve(n, nums)
	}
}

// 31次操作，全是正数需要19次操作，剩12次，
func solve(n int, nums []int) {

}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

//a := make([]int, n)
//for i := range a {
//	fmt.Fscan(in, &a[i])
//}

type neighbor struct {
	to, wt int
}
