package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 그래프 탐색 (Graph Traversal) - BFS와 DFS 기본 패턴 예시
// 그래프는 정점(Vertex)과 간선(Edge)으로 이루어진 자료구조이다.
// 그래프 표현 방식:
//   - 인접 리스트 (Adjacency List): 각 정점에 연결된 정점 목록을 저장한다
//     → 공간 복잡도 O(V + E), 희소 그래프에 적합
//   - 인접 행렬 (Adjacency Matrix): V×V 크기의 2차원 배열로 연결 여부를 저장한다
//     → 공간 복잡도 O(V²), 밀집 그래프에 적합
//
// 예시 1: BFS (너비 우선 탐색, Breadth-First Search)
//   - 시간 복잡도: O(V + E) (V: 정점 수, E: 간선 수)
//   - 공간 복잡도: O(V) (방문 배열 + 큐)
//   - 특징: 시작 정점에서 가까운 정점부터 탐색, 최단 경로 보장 (가중치 없는 그래프)
//
// 예시 2: DFS (깊이 우선 탐색, Depth-First Search)
//   - 시간 복잡도: O(V + E)
//   - 공간 복잡도: O(V) (방문 배열 + 재귀 호출 스택)
//   - 특징: 한 경로를 끝까지 탐색한 뒤 되돌아오는 방식, 경로 탐색/사이클 검출에 유용
//
// 예시 3: BFS 최단 거리 (가중치 없는 그래프)
//   - 시간 복잡도: O(V + E)
//   - 공간 복잡도: O(V) (거리 배열 + 큐)
//
// 예시 4: 다익스트라 (Dijkstra's Algorithm)
//   - 시간 복잡도: O((V + E) log V) (우선순위 큐 사용 시)
//   - 공간 복잡도: O(V + E) (거리 배열 + 우선순위 큐 + 인접 리스트)
//   - 특징: 가중치가 양수인 그래프에서 단일 출발점 최단 경로를 구한다
//
// 예시 5: 유니온-파인드 (Union-Find / Disjoint Set Union)
//   - Find 시간 복잡도: O(α(N)) ≈ O(1) (경로 압축 적용 시)
//   - Union 시간 복잡도: O(α(N)) ≈ O(1) (랭크 기반 합치기 적용 시)
//   - 공간 복잡도: O(N)
//   - 특징: 서로소 집합을 효율적으로 관리, 그래프의 연결 요소/사이클 검출에 활용

// Graph 구조체는 인접 리스트 방식으로 그래프를 표현한다.
// adj[v]는 정점 v에 연결된 인접 정점 목록이다.
type Graph struct {
	V   int     // 정점 수
	adj [][]int // 인접 리스트
}

// NewGraph 함수는 v개의 정점을 가진 빈 그래프를 생성한다.
func NewGraph(v int) *Graph {
	adj := make([][]int, v)
	for i := range adj {
		adj[i] = []int{}
	}
	return &Graph{V: v, adj: adj}
}

// AddEdge 함수는 무방향 간선을 추가한다.
// u와 v 사이에 양방향 연결을 만든다.
func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// BFS 함수는 시작 정점에서 너비 우선 탐색을 수행한다.
// 큐를 사용하여 가까운 정점부터 차례로 방문한다.
// 반환값: 방문 순서대로 정렬된 정점 목록
func (g *Graph) BFS(start int) []int {
	// 방문 여부를 기록하는 배열
	visited := make([]bool, g.V)
	// BFS에 사용할 큐 (슬라이스로 구현)
	queue := []int{start}
	visited[start] = true
	// 방문 순서를 저장할 결과 배열
	order := []int{}

	for len(queue) > 0 {
		// 큐의 맨 앞 정점을 꺼낸다 (FIFO)
		curr := queue[0]
		queue = queue[1:]
		order = append(order, curr)

		// 현재 정점의 모든 인접 정점을 확인한다
		for _, next := range g.adj[curr] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return order
}

// DFS 함수는 시작 정점에서 깊이 우선 탐색을 수행한다.
// 재귀를 사용하여 한 경로를 끝까지 탐색한 뒤 되돌아온다.
// 반환값: 방문 순서대로 정렬된 정점 목록
func (g *Graph) DFS(start int) []int {
	visited := make([]bool, g.V)
	order := []int{}
	g.dfsHelper(start, visited, &order)
	return order
}

// dfsHelper 함수는 DFS의 재귀 보조 함수이다.
// 현재 정점을 방문 처리하고, 방문하지 않은 인접 정점으로 재귀 호출한다.
func (g *Graph) dfsHelper(v int, visited []bool, order *[]int) {
	// 현재 정점을 방문 처리
	visited[v] = true
	*order = append(*order, v)

	// 인접 정점 중 방문하지 않은 정점으로 재귀 탐색
	for _, next := range g.adj[v] {
		if !visited[next] {
			g.dfsHelper(next, visited, order)
		}
	}
}

// DFSIterative 함수는 스택을 사용한 반복적 DFS를 수행한다.
// 재귀 대신 명시적 스택을 사용하여 스택 오버플로를 방지할 수 있다.
func (g *Graph) DFSIterative(start int) []int {
	visited := make([]bool, g.V)
	// DFS에 사용할 스택 (슬라이스로 구현)
	stack := []int{start}
	order := []int{}

	for len(stack) > 0 {
		// 스택의 맨 위 정점을 꺼낸다 (LIFO)
		top := len(stack) - 1
		curr := stack[top]
		stack = stack[:top]

		// 이미 방문한 정점이면 건너뛴다
		if visited[curr] {
			continue
		}

		// 현재 정점을 방문 처리
		visited[curr] = true
		order = append(order, curr)

		// 인접 정점을 스택에 추가 (역순으로 넣어야 작은 번호부터 방문)
		for i := len(g.adj[curr]) - 1; i >= 0; i-- {
			next := g.adj[curr][i]
			if !visited[next] {
				stack = append(stack, next)
			}
		}
	}

	return order
}

// BFSShortestDist 함수는 BFS를 이용하여 시작 정점에서 모든 정점까지의 최단 거리를 구한다.
// 가중치가 없는 그래프에서 BFS는 최단 경로를 보장한다.
// 반환값: dist[v] = 시작 정점에서 정점 v까지의 최단 거리 (-1이면 도달 불가)
func (g *Graph) BFSShortestDist(start int) []int {
	// 거리 배열 초기화 (-1은 아직 방문하지 않음을 의미)
	dist := make([]int, g.V)
	for i := range dist {
		dist[i] = -1
	}

	// 시작 정점의 거리는 0
	dist[start] = 0
	queue := []int{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// 인접 정점 중 아직 방문하지 않은 정점의 거리를 갱신
		for _, next := range g.adj[curr] {
			if dist[next] == -1 {
				dist[next] = dist[curr] + 1
				queue = append(queue, next)
			}
		}
	}

	return dist
}

// Edge 구조체는 가중치가 있는 간선을 표현한다.
type Edge struct {
	To     int // 도착 정점
	Weight int // 간선 가중치
}

// Item 구조체는 우선순위 큐에 들어갈 원소이다.
// 다익스트라에서 (거리, 정점) 쌍을 관리한다.
type Item struct {
	node int // 정점 번호
	dist int // 시작 정점으로부터의 거리
}

// PriorityQueue 는 최소 힙 기반 우선순위 큐이다.
type PriorityQueue []Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// dijkstra 함수는 가중치 그래프에서 시작 정점으로부터 모든 정점까지의 최단 거리를 구한다.
// 우선순위 큐(최소 힙)를 사용하여 가장 가까운 정점부터 처리한다.
// graph[v]는 정점 v에서 나가는 간선(Edge) 목록이다.
// 반환값: dist[v] = 시작 정점에서 정점 v까지의 최단 거리 (도달 불가 시 math.MaxInt64)
func dijkstra(graph [][]Edge, start int) []int {
	n := len(graph)

	// 거리 배열 초기화 (무한대로 설정)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0

	// 우선순위 큐에 시작 정점을 넣는다
	pq := &PriorityQueue{{node: start, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		// 현재 가장 가까운 정점을 꺼낸다
		curr := heap.Pop(pq).(Item)

		// 이미 더 짧은 경로를 찾은 경우 건너뛴다
		if curr.dist > dist[curr.node] {
			continue
		}

		// 인접 간선을 확인하여 거리를 갱신한다
		for _, edge := range graph[curr.node] {
			newDist := dist[curr.node] + edge.Weight
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				heap.Push(pq, Item{node: edge.To, dist: newDist})
			}
		}
	}

	return dist
}

// UnionFind 구조체는 서로소 집합(Disjoint Set)을 관리한다.
// 경로 압축(Path Compression)과 랭크 기반 합치기(Union by Rank)를 적용하여
// 거의 O(1)에 가까운 연산 속도를 보장한다.
type UnionFind struct {
	parent []int // parent[i] = i의 부모 노드
	rank   []int // rank[i] = i를 루트로 하는 트리의 랭크 (높이의 상한)
}

// NewUnionFind 함수는 n개의 원소를 가진 유니온-파인드를 생성한다.
// 초기에는 각 원소가 자기 자신만을 포함하는 독립 집합이다.
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i // 자기 자신이 루트
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find 메서드는 원소 x가 속한 집합의 대표(루트)를 찾는다.
// 경로 압축: 탐색 과정에서 만나는 모든 노드를 루트에 직접 연결한다.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 경로 압축
	}
	return uf.parent[x]
}

// Union 메서드는 원소 x와 y가 속한 두 집합을 합친다.
// 랭크 기반 합치기: 랭크가 낮은 트리를 높은 트리 아래에 붙인다.
// 반환값: 합치기 성공 시 true, 이미 같은 집합이면 false
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// 이미 같은 집합에 속해 있으면 합칠 필요 없음
	if rootX == rootY {
		return false
	}

	// 랭크가 낮은 트리를 높은 트리 아래에 붙인다
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}

func main() {
	// === 그래프 생성 ===
	// 0 -- 1 -- 3
	// |    |    |
	// 2    4    5
	//      |
	//      6
	fmt.Println("=== 그래프 구조 ===")
	fmt.Println()
	fmt.Println("  0 -- 1 -- 3")
	fmt.Println("  |    |    |")
	fmt.Println("  2    4    5")
	fmt.Println("       |")
	fmt.Println("       6")
	fmt.Println()

	g := NewGraph(7)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 6)

	// === BFS (너비 우선 탐색) ===
	fmt.Println("=== BFS (너비 우선 탐색) ===")
	fmt.Println()

	bfsOrder := g.BFS(0)
	fmt.Printf("시작 정점: 0\n")
	fmt.Printf("방문 순서: %v\n", bfsOrder)
	fmt.Println("설명: 가까운 정점부터 차례로 방문한다")

	// === DFS (깊이 우선 탐색 - 재귀) ===
	fmt.Println("\n=== DFS (깊이 우선 탐색 - 재귀) ===")
	fmt.Println()

	dfsOrder := g.DFS(0)
	fmt.Printf("시작 정점: 0\n")
	fmt.Printf("방문 순서: %v\n", dfsOrder)
	fmt.Println("설명: 한 경로를 끝까지 탐색한 뒤 되돌아온다")

	// === DFS (깊이 우선 탐색 - 반복) ===
	fmt.Println("\n=== DFS (깊이 우선 탐색 - 반복/스택) ===")
	fmt.Println()

	dfsIterOrder := g.DFSIterative(0)
	fmt.Printf("시작 정점: 0\n")
	fmt.Printf("방문 순서: %v\n", dfsIterOrder)
	fmt.Println("설명: 스택을 사용하여 재귀 없이 DFS를 수행한다")

	// === BFS 최단 거리 ===
	fmt.Println("\n=== BFS 최단 거리 ===")
	fmt.Println()

	dist := g.BFSShortestDist(0)
	fmt.Printf("시작 정점: 0\n")
	fmt.Println("각 정점까지의 최단 거리:")
	for v, d := range dist {
		fmt.Printf("  정점 %d: 거리 %d\n", v, d)
	}

	// === 연결 요소가 여러 개인 그래프 ===
	fmt.Println("\n=== 비연결 그래프에서의 탐색 ===")
	fmt.Println()

	g2 := NewGraph(6)
	g2.AddEdge(0, 1)
	g2.AddEdge(0, 2)
	g2.AddEdge(3, 4)
	// 정점 5는 고립됨

	fmt.Println("그래프: {0-1, 0-2}, {3-4}, {5}")
	fmt.Printf("정점 0에서 BFS: %v\n", g2.BFS(0))
	fmt.Printf("정점 3에서 BFS: %v\n", g2.BFS(3))

	dist2 := g2.BFSShortestDist(0)
	fmt.Println("정점 0에서의 최단 거리:")
	for v, d := range dist2 {
		if d == -1 {
			fmt.Printf("  정점 %d: 도달 불가\n", v)
		} else {
			fmt.Printf("  정점 %d: 거리 %d\n", v, d)
		}
	}

	// === 다익스트라 (Dijkstra) ===
	fmt.Println("\n=== 다익스트라 (Dijkstra) ===")
	fmt.Println()

	// 가중치 그래프 생성 (인접 리스트)
	//   0 --1-- 1 --3-- 3
	//   |       |       |
	//   4       2       1
	//   |       |       |
	//   2 --5-- 4 --2-- 5
	fmt.Println("  0 --1-- 1 --3-- 3")
	fmt.Println("  |       |       |")
	fmt.Println("  4       2       1")
	fmt.Println("  |       |       |")
	fmt.Println("  2 --5-- 4 --2-- 5")
	fmt.Println()

	wg := make([][]Edge, 6)
	for i := range wg {
		wg[i] = []Edge{}
	}
	// 무방향 가중치 간선 추가
	addWeightedEdge := func(u, v, w int) {
		wg[u] = append(wg[u], Edge{To: v, Weight: w})
		wg[v] = append(wg[v], Edge{To: u, Weight: w})
	}
	addWeightedEdge(0, 1, 1)
	addWeightedEdge(0, 2, 4)
	addWeightedEdge(1, 3, 3)
	addWeightedEdge(1, 4, 2)
	addWeightedEdge(2, 4, 5)
	addWeightedEdge(3, 5, 1)
	addWeightedEdge(4, 5, 2)

	dijkDist := dijkstra(wg, 0)
	fmt.Println("시작 정점: 0")
	fmt.Println("각 정점까지의 최단 거리:")
	for v, d := range dijkDist {
		if d == math.MaxInt64 {
			fmt.Printf("  정점 %d: 도달 불가\n", v)
		} else {
			fmt.Printf("  정점 %d: 거리 %d\n", v, d)
		}
	}

	// === 유니온-파인드 (Union-Find) ===
	fmt.Println("\n=== 유니온-파인드 (Union-Find) ===")
	fmt.Println()

	uf := NewUnionFind(7)
	fmt.Println("초기 상태: 각 원소가 독립 집합")

	// 간선을 추가하며 집합을 합친다
	edges := [][2]int{{0, 1}, {1, 2}, {3, 4}, {5, 6}, {4, 5}}
	for _, e := range edges {
		merged := uf.Union(e[0], e[1])
		fmt.Printf("  Union(%d, %d) → 합침: %v\n", e[0], e[1], merged)
	}

	fmt.Println()
	fmt.Println("같은 집합 여부 확인:")
	queries := [][2]int{{0, 2}, {3, 6}, {0, 3}, {2, 5}}
	for _, q := range queries {
		same := uf.Find(q[0]) == uf.Find(q[1])
		fmt.Printf("  %d와 %d: 같은 집합 = %v\n", q[0], q[1], same)
	}

	// 이미 같은 집합인 원소를 다시 합치기 시도
	fmt.Println()
	merged := uf.Union(0, 2)
	fmt.Printf("  Union(0, 2) → 합침: %v (이미 같은 집합)\n", merged)
}
