package main

// 01dfs-堆   https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/
//func minCost(grid [][]int) int {
//	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
//	n, m := len(grid), len(grid[0])
//	dp := make([][]int, n)
//	for i := range dp {
//		dp[i] = make([]int, m)
//		for j := range dp[i] {
//			dp[i][j] = m * n
//		}
//	}
//	dp[0][0] = 0
//	h := &Heap{}
//	heap.Init(h)
//	heap.Push(h, pair{0, 0, 0})
//	for h.Len() > 0 {
//		top := heap.Pop(h).(pair)
//		for i, v := range dir {
//			x, y, move := top.I+v[0], top.J+v[1], top.M
//			if i+1 != grid[top.I][top.J] {
//				move++
//			}
//			if x >= 0 && x < n && y >= 0 && y < m && move < dp[x][y] {
//				dp[x][y] = move
//				heap.Push(h, pair{x, y, move})
//			}
//		}
//	}
//	return dp[n-1][m-1]
//}
//type pair struct {
//	I int
//	J int
//	M int
//}
//type Heap []pair
//
//func (h Heap) Len() int {
//	return len(h)
//}
//func (h Heap) Less(i, j int) bool {
//	return h[i].M < h[j].M
//}
//func (h Heap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//func (h *Heap) Push(x interface{}) {
//	*h = append(*h, x.(pair))
//}
//func (h *Heap) Pop() interface{} {
//	n := len(*h)
//	x := (*h)[n-1]
//	*h = (*h)[:n-1]
//	return x
//}
