package main

//type Heap []pair
//
//func (h Heap) Len() int            { return len(h) }
//func (h Heap) Less(i, j int) bool  { return h[i].V < h[j].V }
//func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }
//func (h *Heap) Pop() interface{} {
//	x := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return x
//}

//基础大堆
//type Heap []int
//
//func (h Heap) Len() int {
//	return len(h)
//}
//func (h Heap) Less(i, j int) bool {
//	return h[i] > h[j]
//}
//func (h Heap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//func (h *Heap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//func (h *Heap) Pop() interface{} {
//	n := len(*h)
//	x := (*h)[n-1]
//	*h = (*h)[:n-1]
//	return x
//}

//链表堆
//type Heap []*ListNode
//
//func (h Heap) Len() int {
//	return len(h)
//}
//func (h Heap) Less(i, j int) bool {
//	return h[i].Val < h[j].Val
//}
//func (h Heap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//func (h *Heap) Push(x interface{}) {
//	*h = append(*h, x.(*ListNode))
//}
//func (h *Heap) Pop() interface{} {
//	n := len(*h)
//	x := (*h)[n-1]
//	*h = (*h)[:n-1]
//	return x
//}

//Pair堆
//type Heap [][2]int
//
//func (h Heap) Len() int {
//	return len(h)
//}
//
//func (h Heap) Less(i, j int) bool {
//	return h[i][1] > h[j][1]
//}
//
//func (h Heap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *Heap) Push(x interface{}) {
//	*h = append(*h, x.([2]int))
//}
//
//func (h *Heap) Pop() interface{} {
//	l := len(*h)
//	x := (*h)[l-1]
//	*h = (*h)[:l-1]
//	return x
//}

//数组的第k大的子序列
//func kSum(nums []int, k int) int64 {
//	var res int64
//	for i, v := range nums {
//		if v > 0 {
//			res += int64(v)
//		} else {
//			nums[i] = -nums[i]
//		}
//	}
//
//	h := &Heap{}
//	heap.Push(h, pair{0, 0})
//	for k > 1 {
//		k--
//		top := heap.Pop(h).(pair)
//		if top.Index < len(nums) {
//			heap.Push(h, pair{top.Sum + int64(nums[top.Index]), top.Index + 1})
//			if top.Index > 0 {
//				heap.Push(h, pair{top.Sum + int64(nums[top.Index]-nums[top.Index-1]), top.Index + 1})
//			}
//		}
//	}
//	return res - heap.Pop(h).(pair).Sum
//}
//
//type pair struct {
//	Sum   int64
//	Index int
//}
//type Heap []pair
//
//func (h Heap) Len() int {
//	return len(h)
//}
//func (h Heap) Less(i, j int) bool {
//	return h[i].Sum < h[j].Sum
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
