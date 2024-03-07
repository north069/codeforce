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
	for ; T > 0; T-- {
		solve(in, out)
	}
	out.Flush()
}

func solve(in *bufio.Reader, out *bufio.Writer) {

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
