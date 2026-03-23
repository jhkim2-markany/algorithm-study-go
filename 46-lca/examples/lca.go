package main

import (
	"fmt"
	"math"
)

// LCA (최소 공통 조상, Lowest Common Ancestor)
// 나이브 방법, Binary Lifting, 오일러 투어 + RMQ 세 가지 방법을 구현한다.

const MAXN = 100001
const LOG = 17 // log2(100000) ≈ 17

// 트리 인접 리스트
var adj [MAXN][]int

// 나이브 방법에 사용하는 배열
var parentNaive [MAXN]int
var depthNaive [MAXN]int

// Binary Lifting에 사용하는 배열
var up [MAXN][LOG]int
var depth [MAXN]int

// 오일러 투어 + RMQ에 사용하는 배열
var euler []int          // 오일러 투어 순서
var eulerDepth []int     // 오일러 투어에서 각 위치의 깊이
var firstVisit [MAXN]int // 각 노드가 오일러 투어에서 처음 등장하는 위치

// 나이브 방법: DFS로 부모와 깊이를 구한다
// 시간 복잡도: O(N)
func dfsNaive(v, par, d int) {
	parentNaive[v] = par
	depthNaive[v] = d
	for _, u := range adj[v] {
		if u != par {
			dfsNaive(u, v, d+1)
		}
	}
}

// 나이브 LCA: 깊이를 맞추고 한 칸씩 올라간다
// 시간 복잡도: O(N) (최악)
func lcaNaive(u, v int) int {
	// 깊이가 더 깊은 노드를 올린다
	for depthNaive[u] > depthNaive[v] {
		u = parentNaive[u]
	}
	for depthNaive[v] > depthNaive[u] {
		v = parentNaive[v]
	}
	// 같아질 때까지 동시에 올린다
	for u != v {
		u = parentNaive[u]
		v = parentNaive[v]
	}
	return u
}

// Binary Lifting: DFS로 깊이와 희소 배열을 구성한다
// 시간 복잡도: O(N log N)
func dfsBinaryLifting(v, par, d int) {
	depth[v] = d
	up[v][0] = par
	// 2^k번째 조상을 채운다
	for k := 1; k < LOG; k++ {
		up[v][k] = up[up[v][k-1]][k-1]
	}
	for _, u := range adj[v] {
		if u != par {
			dfsBinaryLifting(u, v, d+1)
		}
	}
}

// Binary Lifting LCA: 2의 거듭제곱 단위로 점프한다
// 시간 복잡도: O(log N)
func lcaBinaryLifting(u, v int) int {
	// u가 더 깊도록 보장한다
	if depth[u] < depth[v] {
		u, v = v, u
	}
	// 깊이 차이만큼 u를 올린다
	diff := depth[u] - depth[v]
	for k := 0; k < LOG; k++ {
		if (diff>>k)&1 == 1 {
			u = up[u][k]
		}
	}
	// 같으면 LCA를 찾은 것이다
	if u == v {
		return u
	}
	// 큰 점프부터 시도하며 LCA 바로 아래까지 올린다
	for k := LOG - 1; k >= 0; k-- {
		if up[u][k] != up[v][k] {
			u = up[u][k]
			v = up[v][k]
		}
	}
	// 한 칸 위가 LCA이다
	return up[u][0]
}

// 오일러 투어: DFS 순서대로 방문 노드를 기록한다
func dfsEuler(v, par, d int) {
	firstVisit[v] = len(euler)
	euler = append(euler, v)
	eulerDepth = append(eulerDepth, d)
	for _, u := range adj[v] {
		if u != par {
			dfsEuler(u, v, d+1)
			// 서브트리에서 돌아올 때 다시 기록한다
			euler = append(euler, v)
			eulerDepth = append(eulerDepth, d)
		}
	}
}

// 희소 테이블 (Sparse Table) - 구간 최솟값 쿼리용
var sparseTable [][]int // 인덱스 저장
var sparseLog []int

// 희소 테이블을 구성한다
// 시간 복잡도: O(N log N)
func buildSparseTable() {
	n := len(eulerDepth)
	// 로그 값 전처리
	sparseLog = make([]int, n+1)
	for i := 2; i <= n; i++ {
		sparseLog[i] = sparseLog[i/2] + 1
	}
	maxLog := sparseLog[n] + 1
	sparseTable = make([][]int, maxLog)
	for k := 0; k < maxLog; k++ {
		sparseTable[k] = make([]int, n)
	}
	// 길이 1인 구간 초기화
	for i := 0; i < n; i++ {
		sparseTable[0][i] = i
	}
	// 길이 2^k인 구간을 채운다
	for k := 1; k < maxLog; k++ {
		for i := 0; i+(1<<k)-1 < n; i++ {
			left := sparseTable[k-1][i]
			right := sparseTable[k-1][i+(1<<(k-1))]
			if eulerDepth[left] <= eulerDepth[right] {
				sparseTable[k][i] = left
			} else {
				sparseTable[k][i] = right
			}
		}
	}
}

// 오일러 투어 + RMQ로 LCA를 구한다
// 시간 복잡도: O(1) (전처리 후)
func lcaEulerRMQ(u, v int) int {
	l := firstVisit[u]
	r := firstVisit[v]
	if l > r {
		l, r = r, l
	}
	length := r - l + 1
	k := int(math.Log2(float64(length)))
	left := sparseTable[k][l]
	right := sparseTable[k][r-(1<<k)+1]
	if eulerDepth[left] <= eulerDepth[right] {
		return euler[left]
	}
	return euler[right]
}

func main() {
	// 예제 트리 구성 (1-indexed, 루트: 1)
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6
	//  /
	// 7
	n := 7
	edges := [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {4, 7}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 1. 나이브 방법
	fmt.Println("=== 나이브 방법 ===")
	dfsNaive(1, 0, 0)
	fmt.Printf("LCA(4, 5) = %d\n", lcaNaive(4, 5)) // 2
	fmt.Printf("LCA(4, 6) = %d\n", lcaNaive(4, 6)) // 1
	fmt.Printf("LCA(7, 5) = %d\n", lcaNaive(7, 5)) // 2
	fmt.Printf("LCA(7, 3) = %d\n", lcaNaive(7, 3)) // 1

	// 2. Binary Lifting
	fmt.Println("\n=== Binary Lifting ===")
	dfsBinaryLifting(1, 0, 0)
	fmt.Printf("LCA(4, 5) = %d\n", lcaBinaryLifting(4, 5)) // 2
	fmt.Printf("LCA(4, 6) = %d\n", lcaBinaryLifting(4, 6)) // 1
	fmt.Printf("LCA(7, 5) = %d\n", lcaBinaryLifting(7, 5)) // 2
	fmt.Printf("LCA(7, 3) = %d\n", lcaBinaryLifting(7, 3)) // 1

	// K번째 조상 구하기 (Binary Lifting 활용)
	fmt.Println("\n=== K번째 조상 ===")
	// 노드 7의 1번째 조상 = 4, 2번째 조상 = 2, 3번째 조상 = 1
	node := 7
	for k := 1; k <= 3; k++ {
		ancestor := node
		remaining := k
		for bit := 0; bit < LOG && remaining > 0; bit++ {
			if (remaining>>bit)&1 == 1 {
				ancestor = up[ancestor][bit]
			}
		}
		fmt.Printf("노드 %d의 %d번째 조상 = %d\n", node, k, ancestor)
	}

	// 3. 오일러 투어 + RMQ
	fmt.Println("\n=== 오일러 투어 + RMQ ===")
	euler = nil
	eulerDepth = nil
	dfsEuler(1, 0, 0)
	buildSparseTable()
	fmt.Printf("LCA(4, 5) = %d\n", lcaEulerRMQ(4, 5)) // 2
	fmt.Printf("LCA(4, 6) = %d\n", lcaEulerRMQ(4, 6)) // 1
	fmt.Printf("LCA(7, 5) = %d\n", lcaEulerRMQ(7, 5)) // 2
	fmt.Printf("LCA(7, 3) = %d\n", lcaEulerRMQ(7, 3)) // 1

	// 두 노드 사이의 거리 (LCA 활용)
	fmt.Println("\n=== 두 노드 사이의 거리 ===")
	pairs := [][2]int{{4, 5}, {7, 6}, {7, 3}}
	for _, p := range pairs {
		u, v := p[0], p[1]
		lca := lcaBinaryLifting(u, v)
		dist := depth[u] + depth[v] - 2*depth[lca]
		fmt.Printf("dist(%d, %d) = %d (LCA = %d)\n", u, v, dist, lca)
	}

	_ = n
}
