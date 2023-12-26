package main

//区间修改，区间查询线段树
//type segTree struct {
//	Tree map[int]bool
//	Lazy map[int]bool
//}
//
//func (t *segTree) Spread(idx int) {
//	if t.Lazy[idx] {
//		t.Lazy[idx] = false
//		t.Tree[idx*2] = t.Tree[idx]
//		t.Tree[idx*2+1] = t.Tree[idx]
//		t.Lazy[idx*2] = true
//		t.Lazy[idx*2+1] = true
//	}
//}
//func (t *segTree) Update(start, end int, l, r int, idx int, val bool) {
//	if start > r || end < l {
//		return
//	}
//	if start >= l && end <= r {
//		t.Tree[idx] = val
//		t.Lazy[idx] = true
//		return
//	}
//	t.Spread(idx)
//	mid := (start + end) >> 1
//	t.Update(start, mid, l, r, idx*2, val)
//	t.Update(mid+1, end, l, r, idx*2+1, val)
//	t.Tree[idx] = t.Tree[idx*2] && t.Tree[idx*2+1]
//}
//
//func (t *segTree) query(start, end int, l, r int, idx int) bool {
//	if start > r || end < l {
//		return true
//	}
//	if start >= l && end <= r {
//		return t.Tree[idx]
//	}
//	if t.Lazy[idx] {
//		return t.Tree[idx]
//	}
//	mid := (start + end) >> 1
//	left := t.query(start, mid, l, r, idx*2)
//	right := t.query(mid+1, end, l, r, idx*2+1)
//	return left && right
//}

//单点修改，区间查询线段树，不带懒标记
//type segTree struct {
//	Tree []int
//}
//
//func (s *segTree) add(l, r, start, end, idx, val int) {
//	if l > end || r < start {
//		return
//	}
//	if l >= start && r <= end {
//		s.Tree[idx] = max(s.Tree[idx], val)
//		return
//	}
//	mid := (l + r) >> 1
//	s.add(l, mid, start, end, idx*2, val)
//	s.add(mid+1, r, start, end, idx*2+1, val)
//	s.Tree[idx] = max(s.Tree[idx*2], s.Tree[idx*2+1])
//}
//func (s *segTree) query(l, r, start, end, idx int) int {
//	if l > end || r < start {
//		return 0
//	}
//	if l >= start && r <= end {
//		return s.Tree[idx]
//	}
//	mid := (l + r) >> 1
//	return max(s.query(l, mid, start, end, idx*2), s.query(mid+1, r, start, end, idx*2+1))
//}

//type fenwick []int
//
//func newFenwick(n int) fenwick {
//	return make([]int, n+1)
//}
//func (f fenwick) add(i, val int) {
//	for ; i < len(f); i += i & -i {
//		f[i] += val
//	}
//}
//func (f fenwick) pre(i int) int {
//	res := 0
//	for ; i > 0; i -= i & -i {
//		res += f[i]
//	}
//	return res
//}
