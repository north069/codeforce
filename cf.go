package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var T int
	fmt.Fscan(in, &T)
	res := 0
	type pair struct {
		//L int
		Not  int
		Mask int
	}
	m := map[pair]int{}
	for ; T > 0; T-- {
		var s string
		fmt.Fscan(in, &s)
		now := 0
		app := 0
		for _, v := range s {
			app |= 1 << int(v-'a')
			now = now ^ (1 << int(v-'a'))
		}
		for i := 0; i < 26; i++ {
			if app>>i&1 == 0 {
				res += m[pair{i, now ^ ((1 << 26) - 1) ^ (1 << i)}]
				m[pair{i, now}]++
			}
		}
	}
	fmt.Fprintln(out, res)

	out.Flush()
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func solve(n int, nums [][]int, out *bufio.Writer) {

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
