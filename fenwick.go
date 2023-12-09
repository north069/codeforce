package main

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
