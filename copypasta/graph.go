package copypasta

import (
	. "container/heap"
	"sort"
)

type neighbor struct {
	to     int
	weight int64
}

type graph struct {
	size     int
	edgeSize int
	edges    [][]neighbor
	visited  []bool
}

func newGraph(size, edgeSize int) *graph {
	return &graph{
		size:     size,
		edgeSize: edgeSize,
		edges:    make([][]neighbor, size+1),
		visited:  make([]bool, size+1),
	}
}

func (g *graph) add(from, to int, weight int64) {
	g.edges[from] = append(g.edges[from], neighbor{to, weight})
}

func (g *graph) addBoth(from, to int, weight int64) {
	g.add(from, to, weight)
	if from != to {
		g.add(to, from, weight)
	}
}

func (g *graph) resetStates() {
	g.visited = make([]bool, g.size+1)
}

// 遍历所有边
func (g *graph) dfs(v int, do func(from, to int, weight int64)) {
	g.visited[v] = true
	for _, e := range g.edges[v] {
		w, weight := e.to, e.weight
		if !g.visited[w] {
			do(v, w, weight)
			g.dfs(w, do)
		}
	}
}

// 遍历所有边
func (g *graph) bfs(v int, do func(from, to int, weight int64)) {
	g.visited[v] = true
	for queue := []int{v}; len(queue) > 0; {
		v, queue = queue[0], queue[1:]
		for _, e := range g.edges[v] {
			w := e.to
			if !g.visited[w] {
				do(v, w, e.weight)
				g.visited[w] = true
				queue = append(queue, w)
			}
		}
	}
}

// 遍历所有点
func (*graph) bfsWithDepth(g [][]int, st int, do func(v, dep int)) {
	visited := make([]bool, len(g))
	visited[st] = true
	type pair struct{ v, dep int }
	queue := []pair{{st, 0}}
	for len(queue) > 0 {
		var p pair
		p, queue = queue[0], queue[1:]
		do(p.v, p.dep)
		for _, w := range g[p.v] {
			if !visited[w] {
				visited[w] = true
				queue = append(queue, pair{w, p.dep + 1})
			}
		}
	}
}

// https://oi-wiki.org/graph/bridge/
// https://codeforces.com/blog/entry/68138
// 题目推荐 https://cp-algorithms.com/graph/bridge-searching.html#toc-tgt-2
func (*graph) findBridges(n, m int) {
	// TODO: implement
}

// 传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/shortest-path/#floyd
// 题目推荐 https://cp-algorithms.com/graph/all-pair-shortest-path-floyd-warshall.html#toc-tgt-5
func (*graph) shortestPathFloyd(dist [][]int64) [][]int64 {
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	for k := range dist { // 阶段
		for i := range dist { // 状态
			for j := range dist { // 决策
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return dist
}

// 传入邻接矩阵 weights
// weights[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/min-circle/#floyd
func (*graph) shortestCycleFloyd(weights [][]int64) int64 {
	min := func(a, b int64) int64 {
		if a <= b {
			return a
		}
		return b
	}
	const inf int64 = 1e18
	//const inf int = 1e8 // *NOTE*
	n := len(weights)
	dist := make([][]int64, n)
	for i := range dist {
		dist[i] = make([]int64, n)
		copy(dist[i], weights[i])
	}
	ans := inf
	for k := range dist { // 阶段
		for i := 0; i < k; i++ { // 状态
			for j := 0; j < i; j++ { // 决策
				ans = min(ans, dist[i][j]+weights[i][k]+weights[k][j])
			}
		}
		for i := range dist { // 状态
			for j := range dist { // 决策
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return ans
}

// 适用于稀疏图 O((|E|+|V|)⋅log|V|)
// https://oi-wiki.org/graph/shortest-path/#dijkstra
// 题目推荐 https://cp-algorithms.com/graph/dijkstra.html#toc-tgt-5
func (*graph) shortestPathDijkstra(n, m, start int) (dist []int64, parents []int) {
	type neighbor struct {
		to     int
		weight int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w int
		var weight int64
		//v, w, weight := read()-1, read()-1, read()
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}

	const inf int64 = 1e18 // 1e9+1
	dist = make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0
	parents = make([]int, n)
	for i := range parents {
		parents[i] = -1
	}

	h := &pairHeap{}
	Push(h, hPair{0, start})
	for h.Len() > 0 {
		p := Pop(h).(hPair)
		d, v := p.x, p.y
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.weight; newD < dist[w] {
				dist[w] = newD
				parents[w] = v
				Push(h, hPair{newD, w})
			}
		}
	}

	// path from end to start
	var end = n - 1
	path := make([]int, 0, n)
	for v := end; v != -1; v = parents[v] {
		path = append(path, v)
	}

	return
}

// https://oi-wiki.org/graph/shortest-path/#bellman-ford
// https://cp-algorithms.com/graph/bellman_ford.html
// https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
func (*graph) shortestPathBellmanFord(n, m, start int) (dist []int64) {
	type neighbor struct {
		to     int
		weight int64
	}
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		var v, w int
		var weight int64
		//v, w, weight := read()-1, read()-1, read()
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}

	const inf int64 = 1e18 // 1e9+1
	dist = make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0
	onQ := make([]bool, n)
	onQ[start] = true
	relaxedCnt := make([]int, n)

	queue := []int{start}
	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		onQ[v] = false
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.weight; newD < dist[w] {
				dist[w] = newD
				// there is no reason to put a vertex in the queue if it is already in.
				if !onQ[w] {
					queue = append(queue, w)
					onQ[w] = true
					relaxedCnt[w]++
					if relaxedCnt[w] > n {
						// found negative cycle
						return nil
					}
				}
			}
		}
	}
	return
}

// https://cp-algorithms.com/graph/finding-negative-cycle-in-graph.html
func (*graph) hasNegativeCycleBellmanFord() bool {
	// TODO implement
	return false
}

// 适用于稀疏图 O(|E|⋅log|E|)
// https://oi-wiki.org/graph/mst/#kruskal
// 题目推荐 https://cp-algorithms.com/graph/mst_kruskal.html#toc-tgt-5
func (*graph) mstKruskal(n, m int) (sum int64) {
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	type edge struct {
		v, w   int
		weight int
	}
	edges := make([]edge, m)
	for i := range edges {
		var v, w, weight int
		//v, w, weight := read()-1, read()-1, read()
		edges[i] = edge{v, w, weight}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })
	initFa(n)
	for _, e := range edges {
		if from, to := find(e.v), find(e.w); from != to {
			sum += int64(e.weight)
			fa[from] = to
		}
	}
	return
}

// 适用于稠密图 O(|V|^2)，传入邻接矩阵 dist
// dist[v][w] == inf 表示没有 v-w 边
// https://oi-wiki.org/graph/mst/#prim
func (*graph) mstPrim(dist [][]int) (sum int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const inf = 1.1e9
	n := len(dist)
	minWeights := make([]int, n)
	for i := range minWeights {
		minWeights[i] = inf
	}
	minWeights[0] = 0
	used := make([]bool, n)
	for {
		v := -1
		for i, use := range used {
			if !use && (v == -1 || minWeights[i] < minWeights[v]) {
				v = i
			}
		}
		if v == -1 {
			break
		}
		used[v] = true
		sum += minWeights[v]
		for w := range minWeights {
			minWeights[w] = min(minWeights[w], dist[v][w])
		}
	}
	return
}

// Second best Minimum Spanning Tree
// Using Kruskal and Lowest Common Ancestor
// https://cp-algorithms.com/graph/second_best_mst.html
func (*graph) secondMST(n, m int) (sum int64) {
	// TODO implement
	return
}

// 反图的连通分量 O(n+m)
// https://www.luogu.com.cn/blog/endlesscheng/solution-cf1242b
func (*graph) inverseGraphComponents(n, m int) (components [][]int) {
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	same := func(x, y int) bool { return find(x) == find(y) }

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var v, w int
		//v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	maxDegInv, maxDegInvV := 0, 0
	for v, edges := range g {
		if degInv := n - 1 - len(edges); degInv > maxDegInv {
			maxDegInv = degInv
			maxDegInvV = v
		}
	}
	if maxDegInv == 0 {
		components = make([][]int, n)
		for i := range components {
			components[i] = []int{i + 1}
		}
		return
	}

	mergeInv := func(v int, edges []int) {
		vs := map[int]bool{v: true}
		for _, w := range edges {
			vs[w] = true
		}
		for i := 0; i < n; i++ {
			if !vs[i] {
				merge(i, v)
			}
		}
	}
	initFa(n)
	mergeInv(maxDegInvV, g[maxDegInvV])
	for v, edges := range g {
		if !same(v, maxDegInvV) {
			mergeInv(v, edges)
		}
	}

	componentsMap := map[int][]int{}
	for i := range fa {
		rootV := find(i)
		componentsMap[rootV] = append(componentsMap[rootV], i+1)
	}
	components = make([][]int, 0, len(componentsMap))
	for _, vs := range componentsMap {
		components = append(components, vs)
	}
	return
}

// 二分图判定
// https://oi-wiki.org/graph/bi-graph/#_3
// https://cp-algorithms.com/graph/bipartite-check.html
// TODO: https://codeforces.com/contest/1144/problem/F
func (*graph) isBipartite(n, m int) bool {
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var v, w int
		//v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	colors := make([]int8, n)
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[v] {
			if colors[w] == c {
				return false
			}
			if colors[w] == 0 {
				if !f(w, 3-c) {
					return false
				}
			}
		}
		return true
	}
	for i, c := range colors {
		if c == 0 && !f(i, 1) {
			return false
		}
	}
	return true
}

// Kahn's algorithm
// isDAG = len(order)==n
// https://oi-wiki.org/graph/topo/
// https://cp-algorithms.com/graph/topological-sort.html
func (*graph) topSort(n, m int) (order []int, parents []int) {
	g := make([][]int, n)
	inDeg := make([]int, n)
	for i := 0; i < m; i++ {
		var v, w int
		//v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		inDeg[w]++
	}

	parents = make([]int, n)
	for i := range parents {
		parents[i] = -1
	}

	order = make([]int, 0, n)
	queue := []int{}
	levels := make([]int, n)
	for i, deg := range inDeg {
		if deg == 0 {
			queue = append(queue, i)
			levels[i] = 1
		}
	}
	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		order = append(order, v)
		for _, w := range g[v] {
			inDeg[w]--
			if inDeg[w] == 0 {
				queue = append(queue, w)
				levels[w] = levels[v] + 1
				parents[w] = v
			}
		}
	}

	// path from end to start
	var end = n - 1
	path := make([]int, 0, n)
	for v := end; v != -1; v = parents[v] {
		path = append(path, v)
	}

	return
}

// 强连通分量分解
// https://oi-wiki.org/graph/scc/#kosaraju
func (*graph) scc(n, m int) (components [][]int, sccIDs []int) {
	g := make([][]int, n)
	rg := make([][]int, n)
	for i := 0; i < m; i++ {
		var v, w int
		//v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}
	used := make([]bool, n)

	vs := make([]int, 0, n)
	var dfs func(int)
	dfs = func(v int) {
		used[v] = true
		for _, w := range g[v] {
			if !used[w] {
				dfs(w)
			}
		}
		// 后序遍历
		vs = append(vs, v)
	}
	for v := range g {
		if !used[v] {
			dfs(v)
		}
	}

	used = make([]bool, n)
	var comp []int
	var rdfs func(int)
	rdfs = func(v int) {
		used[v] = true
		comp = append(comp, v)
		for _, w := range rg[v] {
			if !used[w] {
				rdfs(w)
			}
		}
	}
	for i := len(vs) - 1; i >= 0; i-- {
		if !used[vs[i]] {
			comp = []int{}
			rdfs(vs[i])
			components = append(components, comp)
		}
	}

	sccIDs = make([]int, n)
	for i, cp := range components {
		for _, v := range cp {
			sccIDs[v] = i
		}
	}

	//lastComp := components[len(components)-1]
	//numCanBeVisitedFromAll := len(lastComp)
	//_ = numCanBeVisitedFromAll
	//used = make([]bool, n)
	//rdfs(lastComp[0])
	//for _, use := range used {
	//	if !use {
	//		numCanBeVisitedFromAll = 0
	//	}
	//}

	return
}

// https://oi-wiki.org/graph/2-sat/
// https://cp-algorithms.com/graph/2SAT.html
func (*graph) solve2SAT(n, m int) bool {
	// TODO implement
	return false
}
