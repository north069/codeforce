package main

import (
	"fmt"
	"strings"
)

func main() {
	genereteArray("")
	fmt.Println()
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
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

// 生成数组测试数据
func genereteArray(str string) {
	s := strings.Replace(str, "[", "{", len(str))
	s = strings.Replace(s, "]", "}", len(s))
	fmt.Println(s)
}
