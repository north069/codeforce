package main

import (
	"fmt"
	"strings"
)

func main() {
	genereteArray("")
	fmt.Println()
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
