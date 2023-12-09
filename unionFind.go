package main

//type unionFind struct {
//	Parent, Rank []int
//}
//
//func newUnionFind(n int) unionFind {
//	parent := make([]int, n)
//	for i := range parent {
//		parent[i] = i
//	}
//	rank := make([]int, n)
//	return unionFind{parent, rank}
//}
//
//func (uf unionFind) merge(x, y int) {
//	x, y = uf.find(x), uf.find(y)
//	if x == y {
//		return
//	}
//	if uf.Rank[x] > uf.Rank[y] {
//		uf.Parent[y] = x
//	} else if uf.Rank[y] > uf.Rank[x] {
//		uf.Parent[x] = y
//	} else {
//		uf.Parent[y] = x
//		uf.Rank[x]++
//	}
//}
//
//func (uf unionFind) find(n int) int {
//	if n == uf.Parent[n] {
//		return n
//	}
//	uf.Parent[n] = uf.find(uf.Parent[n])
//	return uf.Parent[n]
//}
