package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	genereteArray("")
	fmt.Println()
}

type pair struct {
	N int
	C int
}

// 数组原地去重
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

// 返回整个数组到v的距离
// 求前缀和
func getArrDis(nums []int, v int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + nums[i-1]
	}
	getDis := func(v int) int {
		i := sort.SearchInts(nums, v)
		res := i*v - pre[i]
		res += pre[n] - pre[i] - v*(n-i)
		return res
	}
	return getDis(v)
}

// 求数的二进制位数
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

// 判断是否是质数
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

// 让数组相等的最小代价
func LetNumsEqual(nums []int) int {
	sort.Ints(nums)
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		res += abs(nums[i] - nums[n/2])
	}
	return res
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

// 字符串反转
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

// 按顺序生成1到1e9+1的回文数
func getPalArr() []int {
	pal := make([]int, 0, 109999)
	for base := 1; base <= 10000; base *= 10 {
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			pal = append(pal, x)
		}
		if base <= 1000 {
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				pal = append(pal, x)
			}
		}
	}
	pal = append(pal, 1_000_000_001)
	return pal
}

// 生成数组测试数据
func genereteArray(str string) {
	s := strings.Replace(str, "[", "{", len(str))
	s = strings.Replace(s, "]", "}", len(s))
	fmt.Println(s)
}
