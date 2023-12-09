package main

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

//是否存在的动态线段树
//type segTree struct {
//	Tree map[int]bool //整个区间有没有预定的
//	Lazy map[int]bool //预定了
//}
//
////idx从1开始， 左子树idx*2， 右子树idx*2+1
//
//func (t *segTree) query(start, end int, l, r int, idx int) bool {
//	if start > r || end < l { //没有重叠
//		return false
//	}
//	if t.Lazy[idx] { //访问过
//		return true
//	}
//	if l <= start && end <= r { //完全包含，直接返回
//		return t.Tree[idx]
//	}
//	mid := (start + end) >> 1
//	left := t.query(start, mid, l, r, idx*2)
//	right := t.query(mid+1, end, l, r, idx*2+1)
//	return left || right
//}
//func (t *segTree) update(start, end int, l, r int, idx int) bool {
//	if start > r || end < l { //没有重叠
//		return false
//	}
//	if l <= start && end <= r { //完全包含，直接返回
//		t.Lazy[idx] = true
//		t.Tree[idx] = true
//		return true
//	}
//	mid := (start + end) >> 1
//	left := t.update(start, mid, l, r, idx*2)
//	right := t.update(mid+1, end, l, r, idx*2+1)
//	t.Tree[idx] = left || right
//	return t.Tree[idx]
//}

//带延迟标记的线段树
//type MyCalendarTwo struct {
//	ST segTree
//}
//
//func Constructor() MyCalendarTwo {
//	return MyCalendarTwo{ST: segTree{map[int]int{}, map[int]int{}}}
//}
//
//func (this *MyCalendarTwo) Book(start int, end int) bool {
//	this.ST.update(1, 1e9, start, end-1, 1, 1)
//	if this.ST.Tree[1] > 2 {
//		this.ST.update(1, 1e9, start, end-1, 1, -1)
//		return false
//	}
//	return true
//}
//
//type segTree struct {
//	Tree map[int]int
//	Lazy map[int]int
//}
//
//func (t *segTree) update(start, end int, l, r int, idx int, val int) {
//	if start > r || end < l {
//		return
//	}
//	if l <= start && end <= r {
//		t.Lazy[idx] += val //这里lazy加了值，子节点不一定加上了
//		t.Tree[idx] += val
//		return
//	}
//	mid := (start + end) >> 1
//	t.update(start, mid, l, r, idx*2, val)
//	t.update(mid+1, end, l, r, idx*2+1, val)
//	t.Tree[idx] = t.Lazy[idx] + max(t.Tree[idx*2], t.Tree[idx*2+1])		//下面返回的+下面因为Lazy没加的
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
