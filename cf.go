package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	res := ""
	for i := 0; i < n; i++ {
		var s string
		var pos int
		fmt.Fscan(in, &s, &pos)
		now := len(s)
		cnt := 0
		for pos > now {
			pos -= now
			now--
			cnt++
		}
		st := make([]int, 0)
		for j, v := range s {
			for cnt > 0 && len(st) > 0 && v < rune(s[st[len(st)-1]]) {
				st = st[:len(st)-1]
				cnt--
			}
			st = append(st, j)
		}
		res += string(s[st[pos-1]])
	}
	fmt.Println(res)
}

//aadba

//func max(i, j int) int {
//	if i > j {
//		return i
//	} else {
//		return j
//	}
//}
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
