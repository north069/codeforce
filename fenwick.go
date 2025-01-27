package main

// 区间和树状数组，单点修改，区间查询
//type fenwick []int
//
//func newFenwickTree(n int) fenwick {
//	return make([]int, n+1)
//}
//func (f fenwick) add(i, val int) {
//	for ; i < len(f); i += i & -i {
//		//f[i] = max(val, f[i])
//		f[i] += val
//	}
//}
//func (f fenwick) pre(i int) (res int) {
//	for ; i > 0; i -= i & -i {
//		res += f[i]
//		//res = max(res, f[i])
//	}
//	return
//}
//func (f fenwick) query(l, r int) int {
//	return f.pre(r) - f.pre(l-1)
//}

// 单点修改，区间查询最值树状数组
// type fenwick struct {
// 	P []int
// 	T []int
// }

// func newFenwick(n int) fenwick {
// 	return fenwick{make([]int, n+1), make([]int, n+1)}
// }
// func (f fenwick) add(x, v int) {
// 	f.P[x] = max(f.P[x], v)
// 	for ; x < len(f.T); x += x & -x {
// 		f.T[x] = max(f.T[x], v)
// 	}
// }
// func (f fenwick) query(l, r int) int {
// 	if r < l {
// 		return 0
// 	}
// 	lowBit := r & -r
// 	if r-lowBit+1 >= l {
// 		return max(f.T[r], f.query(l, r-lowBit))
// 	} else {
// 		return max(f.P[r], f.query(l, r-1))
// 	}
// }

//离散化数组，最大的是1
//func resultArray(nums []int) []int {
//	n := len(nums)
//	p := make([]pair, n)
//	for i := range nums {
//		p[i] = pair{nums[i], i}
//	}
//	sort.Slice(p, func(i, j int) bool {
//		return p[i].V > p[j].V
//	})
//	id := make([]int, n)
//	id[p[0].I] = 1
//	for i := 1; i < n; i++ {
//		if p[i].V == p[i-1].V {
//			id[p[i].I] = id[p[i-1].I]
//		} else {
//			id[p[i].I] = i + 1
//		}
//	}
//}
//type pair struct {
//	V int
//	I int
//}
