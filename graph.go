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

//func floyd(dist [][]int) [][]int {
//	for k := range dist {
//		for i := range dist {
//			for j := range dist {
//				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
//			}
//		}
//	}
//	return dist
//}

// tarjan求桥
//func findBridge(n int, edge [][]int) [][]int {
//	type neighbor struct{ to, eid int }
//	g := make([][]neighbor, n)
//	for i := range g {
//		g[i] = make([]neighbor, 0)
//	}
//	for i, v := range edge {
//		g[v[0]] = append(g[v[0]], neighbor{v[1], i})
//		g[v[1]] = append(g[v[1]], neighbor{v[0], i})
//	}
//	dfn := make([]int, n)
//	isBridge := make([]bool, len(edge))
//	time := 0
//	var tarjan func(v int, eid int) int
//	tarjan = func(v int, fid int) int { //使用fid而不是fa用来避免重边
//		time++
//		dfn[v] = time
//		lowV := time
//		for _, e := range g[v] {
//			w := e.to
//			if dfn[w] == 0 {
//				lowW := tarjan(w, e.eid)
//				lowV = min(lowW, lowV)
//				if lowW > dfn[v] { //以w为根的子树中没有返回来的边可以到v，所以v-w一定是桥
//					isBridge[e.eid] = true
//				}
//			} else if e.eid != fid { //
//				lowV = min(lowV, dfn[w])
//			}
//		}
//		return lowV
//	}
//	for v, timestamp := range dfn {
//		if timestamp == 0 {
//			tarjan(v, -1)
//		}
//	}
//
//	res := make([][]int, 0)
//	for i, v := range isBridge {
//		if v {
//			res = append(res, edge[i])
//		}
//	}
//	return res
//}
