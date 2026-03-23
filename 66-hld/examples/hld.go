package main

import "fmt"

// Heavy-Light Decomposition (HLD)
// 트리의 경로를 Heavy Chain으로 분해하고, 세그먼트 트리를 결합하여
// 경로 합 질의를 O(log²N)에 처리하는 예시이다.
// 시간 복잡도: 전처리 O(N), 질의 O(log²N)
// 공간 복잡도: O(N)

const MAXN = 100005

var (
	adj    [][]int // 인접 리스트
	sz     []int   // 서브트리 크기
	depth  []int   // 깊이
	parent []int   // 부모 노드
	heavy  []int   // Heavy Child (-1이면 리프)
	top    []int   // 체인의 최상단 노드
	pos    []int   // 세그먼트 트리 상 위치 (DFS 번호)
	val    []int   // 노드 값
	seg    []int   // 세그먼트 트리
	curPos int     // 현재 DFS 번호
	n      int     // 노드 수
)

// dfs1은 서브트리 크기, 깊이, 부모, Heavy Child를 계산한다
func dfs1(v, par, d int) {
	parent[v] = par
	depth[v] = d
	sz[v] = 1
	heavy[v] = -1
	maxSz := 0

	for _, u := range adj[v] {
		if u == par {
			continue
		}
		dfs1(u, v, d+1)
		sz[v] += sz[u]
		if sz[u] > maxSz {
			maxSz = sz[u]
			heavy[v] = u // 서브트리가 가장 큰 자식이 Heavy Child
		}
	}
}

// dfs2는 Heavy Child를 우선 방문하여 DFS 번호와 체인 정보를 부여한다
func dfs2(v, chainTop int) {
	top[v] = chainTop
	pos[v] = curPos
	curPos++

	// Heavy Child를 먼저 방문하여 같은 체인이 연속 구간이 되게 한다
	if heavy[v] != -1 {
		dfs2(heavy[v], chainTop)
	}

	// Light Child는 새로운 체인을 시작한다
	for _, u := range adj[v] {
		if u == parent[v] || u == heavy[v] {
			continue
		}
		dfs2(u, u) // 새 체인의 top = 자기 자신
	}
}

// --- 세그먼트 트리 ---

// segBuild는 세그먼트 트리를 구축한다
func segBuild(node, s, e int, arr []int) {
	if s == e {
		seg[node] = arr[s]
		return
	}
	mid := (s + e) / 2
	segBuild(node*2, s, mid, arr)
	segBuild(node*2+1, mid+1, e, arr)
	seg[node] = seg[node*2] + seg[node*2+1]
}

// segQuery는 구간 [l, r]의 합을 반환한다
func segQuery(node, s, e, l, r int) int {
	if r < s || e < l {
		return 0
	}
	if l <= s && e <= r {
		return seg[node]
	}
	mid := (s + e) / 2
	return segQuery(node*2, s, mid, l, r) + segQuery(node*2+1, mid+1, e, l, r)
}

// segUpdate는 인덱스 idx의 값을 newVal로 변경한다
func segUpdate(node, s, e, idx, newVal int) {
	if s == e {
		seg[node] = newVal
		return
	}
	mid := (s + e) / 2
	if idx <= mid {
		segUpdate(node*2, s, mid, idx, newVal)
	} else {
		segUpdate(node*2+1, mid+1, e, idx, newVal)
	}
	seg[node] = seg[node*2] + seg[node*2+1]
}

// --- HLD 경로 질의 ---

// pathQuery는 노드 u에서 노드 v까지의 경로 합을 반환한다
func pathQuery(u, v int) int {
	result := 0
	for top[u] != top[v] {
		// 체인의 top이 더 깊은 쪽을 올린다
		if depth[top[u]] < depth[top[v]] {
			u, v = v, u
		}
		// 현재 체인 구간을 질의한다
		result += segQuery(1, 0, n-1, pos[top[u]], pos[u])
		u = parent[top[u]] // 체인을 넘어 부모로 이동
	}
	// 같은 체인 내에서 질의
	if depth[u] > depth[v] {
		u, v = v, u
	}
	result += segQuery(1, 0, n-1, pos[u], pos[v])
	return result
}

func main() {
	// 예시 트리:
	//        1 (val=2)
	//       / \
	//      2   3 (val=5)
	// (val=3) / \
	//    / \  4   5
	//   6   7 (val=1)(val=4)
	// (val=6)(val=8)
	n = 7
	val = []int{0, 2, 3, 5, 1, 4, 6, 8} // 1-indexed

	adj = make([][]int, n+1)
	sz = make([]int, n+1)
	depth = make([]int, n+1)
	parent = make([]int, n+1)
	heavy = make([]int, n+1)
	top = make([]int, n+1)
	pos = make([]int, n+1)
	seg = make([]int, 4*n)

	// 간선 추가 (양방향)
	edges := [][2]int{{1, 2}, {1, 3}, {2, 6}, {2, 7}, {3, 4}, {3, 5}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 1차 DFS: 서브트리 크기, Heavy Child 결정
	dfs1(1, 0, 0)

	// 2차 DFS: DFS 번호 부여, 체인 구성
	curPos = 0
	dfs2(1, 1)

	// DFS 번호 순서대로 값 배열 구성
	flat := make([]int, n)
	for i := 1; i <= n; i++ {
		flat[pos[i]] = val[i]
	}

	// 세그먼트 트리 구축
	segBuild(1, 0, n-1, flat)

	// HLD 결과 출력
	fmt.Println("=== HLD 전처리 결과 ===")
	fmt.Printf("노드:   ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Printf("sz:     ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", sz[i])
	}
	fmt.Println()
	fmt.Printf("heavy:  ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", heavy[i])
	}
	fmt.Println()
	fmt.Printf("pos:    ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", pos[i])
	}
	fmt.Println()
	fmt.Printf("top:    ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", top[i])
	}
	fmt.Println()
	fmt.Printf("depth:  ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", depth[i])
	}
	fmt.Println()

	// 경로 합 질의
	fmt.Println("\n=== 경로 합 질의 ===")

	// 질의 1: 노드 6 → 노드 5 (경로: 6→2→1→3→5, 합 = 6+3+2+5+4 = 20)
	fmt.Printf("pathQuery(6, 5) = %d\n", pathQuery(6, 5))

	// 질의 2: 노드 7 → 노드 4 (경로: 7→2→1→3→4, 합 = 8+3+2+5+1 = 19)
	fmt.Printf("pathQuery(7, 4) = %d\n", pathQuery(7, 4))

	// 질의 3: 노드 6 → 노드 7 (경로: 6→2→7, 합 = 6+3+8 = 17)
	fmt.Printf("pathQuery(6, 7) = %d\n", pathQuery(6, 7))

	// 질의 4: 노드 1 → 노드 1 (자기 자신, 합 = 2)
	fmt.Printf("pathQuery(1, 1) = %d\n", pathQuery(1, 1))

	// 노드 값 갱신 후 질의
	fmt.Println("\n=== 노드 값 갱신 ===")
	fmt.Println("노드 2의 값을 3 → 10으로 변경")
	segUpdate(1, 0, n-1, pos[2], 10)
	fmt.Printf("pathQuery(6, 5) = %d (기대값: 27)\n", pathQuery(6, 5))
}
