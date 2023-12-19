package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
)

// 10 5 10 9 6 9 8 9 5
func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	bit, nums := getB(n)
	mask := 1<<bit - 1

	var dfs func(mask int) int
	dfs = func(mask int) int {
		p := 1
		sum := 0
		last := 0
		for i := 0; i < bit; i++ {
			if mask>>i&1 == 1 {
				last = nums[i]
				sum += p * nums[i]
				p *= 10
			}
		}
		if last == 0 {
			return 0
		}
		if int(math.Sqrt(float64(sum)))*int(math.Sqrt(float64(sum))) == sum {
			return bits.OnesCount(uint(mask))
		}
		res := 0
		for i := 0; i < bit; i++ {
			if mask>>i&1 == 1 {
				res = max(res, dfs(mask^(1<<i)))
			}
		}
		return res
	}
	res := dfs(mask)
	if res == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(bit - res)
	}
}
func getB(n int) (int, []int) {
	res := 0
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
		res++
	}
	return res, nums
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
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
