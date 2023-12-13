package main

import (
	"fmt"
	"strings"
)

func main() {
	genereteArray("")
	fmt.Println()
}

func distinctSequences(n int) int {
	if n == 1 {
		return 6
	}
	const mod int = 1e9 + 7
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, 6)
		for j := range dp[i] {
			dp[i][j] = make([]int, 6)
		}
	}
	//1:23456  2:135  3:1245  4:135  5:12346  6:15
	m := [6]map[int]bool{
		{1: true, 2: true, 3: true, 4: true, 5: true},
		{0: true, 2: true, 4: true},
		{0: true, 1: true, 3: true, 4: true},
		{0: true, 2: true, 4: true},
		{1: true, 2: true, 3: true, 0: true, 5: true},
		{0: true, 4: true},
	}
	for i := 0; i < 6; i++ {
		for last := 0; last < 6; last++ {
			if m[last][i] {
				dp[1][last][i] = 1
			}
		}
	}
	for i := 2; i < n; i++ {
		for last1 := 0; last1 < 6; last1++ {
			for last2 := 0; last2 < 6; last2++ {
				for j := 0; j < 6; j++ {
					if j != last2 && m[j][last1] {
						dp[i][j][last1] += dp[i-1][last1][last2]
						dp[i][j][last1] %= mod
					}
				}
			}
		}
	}
	res := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			res += dp[n-1][i][j]
			res %= mod
		}
	}
	return res
}

func unique(nums []int) []int {
	res := make([]int, len(nums))
	k := -1
	for _, v := range nums {
		if (k < 0 || res[k] != v) && v >= 0 {
			k++
			res[k] = v
		}
	}
	return res[:k+1]
}

type pair struct {
	I int
	V int
}
type Heap []pair

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].V < h[j].V
}
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type Q struct {
	X int
	Y int
	I int
}

func getBit(v int64) int {
	if v == 0 {
		return 1
	}
	res := 0
	for v > 0 {
		res++
		v /= 2
	}
	return res
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

type Node struct {
	Index     int
	Val       int
	Neighbors []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
func reverse(num string) string {
	temp := []byte(num)
	left, right := 0, len(num)-1
	for left < right {
		temp[left], temp[right] = temp[right], temp[left]
		left++
		right--
	}
	return string(temp)
}

// 快速幂
func quickPow(a, n int, mod int) int {
	res := 1
	for n > 0 {
		if n&1 > 0 {
			res *= a
			res %= mod
		}
		a = a * a % mod
		n >>= 1
	}
	return res % mod
}

// 拓展欧几里得
func exgcd(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd(b, a%b)
	y -= a / b * x
	return
}

// 求逆元
func invM(a, m int) int {
	g, x, _ := exgcd(a, m)
	if g != 1 && g != -1 {
		return -1
	}
	return (x%m + m) % m
}

// 生成数组测试数据
func genereteArray(str string) {
	s := strings.Replace(str, "[", "{", len(str))
	s = strings.Replace(s, "]", "}", len(s))
	fmt.Println(s)
}
