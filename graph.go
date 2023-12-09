package main

//dij贪心求单源最短路径 O(n2)
//func dijkstra(g [][]int, st int) []int {
//	n := len(g)
//
//	const inf int = 1e8
//
//	dis := make([]int, n)
//	vis := make([]bool, n)
//	for i := range dis {
//		dis[i] = inf
//	}
//	dis[st] = 0
//	for {
//		v := -1
//		//找到未标记节点中dis最小的
//		for w, b := range vis {
//			if !b && (v < 0 || dis[w] < dis[v]) {
//				v = w
//			}
//		}
//		if v < 0 {
//			return dis
//		}
//		vis[v] = true
//		for w, wt := range g[v] {
//			if newD := wt + dis[v]; newD < dis[w] {
//				dis[w] = newD
//			}
//		}
//	}
//}

//堆优化的dij  O(mlogn)
//type neighbor struct {
//	to, wt int
//}
//
//func dijkstra(g [][]neighbor, st int, n int) []int {
//
//	const inf int = 1e18
//	dis := make([]int, n)
//	for i := range dis {
//		dis[i] = inf
//	}
//	dis[st] = 0
//	h := &dijkstraHeap{{st, 0}}
//	for len(*h) > 0 {
//		top := heap.Pop(h).(dijkstraPair)
//		v := top.v
//		if top.dis > dis[v] {
//			continue
//		}
//		for _, e := range g[v] {
//			w := e.to
//			if newD := dis[v] + e.wt; newD < dis[w] {
//				dis[w] = newD
//				heap.Push(h, dijkstraPair{w, dis[w]})
//			}
//		}
//	}
//	return dis
//}
//
//type dijkstraPair struct{ v, dis int }
//type dijkstraHeap []dijkstraPair
//
//func (h dijkstraHeap) Len() int           { return len(h) }
//func (h dijkstraHeap) Less(i, j int) bool { return h[i].dis < h[j].dis }
//func (h dijkstraHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *dijkstraHeap) Push(v any)        { *h = append(*h, v.(dijkstraPair)) }
//func (h *dijkstraHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

//func (h *dijkstraHeap) push(v dijkstraPair) { heap.Push(h, v) }
//func (h *dijkstraHeap) pop() dijkstraPair   { return heap.Pop(h).(dijkstraPair) }
