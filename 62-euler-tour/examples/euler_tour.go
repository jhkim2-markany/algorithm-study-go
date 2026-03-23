package main

import "fmt"

// 오일러 경로 테크닉 (Euler Tour Technique)
// 트리를 DFS로 순회하여 1차원 배열로 평탄화하고,
// 서브트리 합 질의를 누적 합(prefix sum)으로 처리하는 예시이다.
// 시간 복잡도: 전처리 O(N), 질의 O(1)
// 공간 복잡도: O(N)

var (
	adj   [][]int // 인접 리스트
	in    []int   // 방문 시작 시각
	out   []int   // 방문 종료 시각
	euler []int   // 오일러 투어 순서 (euler[i] = i번째 방문한 노드)
	timer int     // 타이머
)

// dfs는 오일러 투어를 수행하여 in/out 타임스탬프를 기록한다
func dfs(v, parent int) {
	in[v] = timer
	euler[timer] = v
	timer++

	for _, u := range adj[v] {
		if u == parent {
			continue // 부모 방향 역행 방지
		}
		dfs(u, v)
	}

	out[v] = timer - 1
}

func main() {
	// 예시 트리:
	//       1 (val=3)
	//      / \
	//     2   3 (val=1)
	// (val=5)   \
	//   / \      4 (val=2)
	//  5   6
	// (val=4)(val=7)
	n := 6
	val := []int{0, 3, 5, 1, 2, 4, 7} // 1-indexed 노드 값

	adj = make([][]int, n+1)
	in = make([]int, n+1)
	out = make([]int, n+1)
	euler = make([]int, n)

	// 간선 추가 (양방향)
	edges := [][2]int{{1, 2}, {1, 3}, {2, 5}, {2, 6}, {3, 4}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 오일러 투어 수행 (루트 = 1)
	timer = 0
	dfs(1, 0)

	// 타임스탬프 출력
	fmt.Println("=== 오일러 투어 결과 ===")
	fmt.Printf("노드:  ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Printf("in:    ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", in[i])
	}
	fmt.Println()
	fmt.Printf("out:   ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", out[i])
	}
	fmt.Println()

	// 평탄화 배열 구성
	flat := make([]int, n) // flat[i] = val[euler[i]]
	for i := 0; i < n; i++ {
		flat[i] = val[euler[i]]
	}

	fmt.Println("\n=== 평탄화 배열 ===")
	fmt.Printf("인덱스: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Printf("노드:   ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", euler[i])
	}
	fmt.Println()
	fmt.Printf("값:     ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", flat[i])
	}
	fmt.Println()

	// 누적 합 구성 (서브트리 합 질의용)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + flat[i]
	}

	// 서브트리 합 질의: sum(in[v]..out[v]) = prefix[out[v]+1] - prefix[in[v]]
	fmt.Println("\n=== 서브트리 합 질의 ===")
	for v := 1; v <= n; v++ {
		sum := prefix[out[v]+1] - prefix[in[v]]
		fmt.Printf("노드 %d의 서브트리 합: %d (구간 [%d, %d])\n", v, sum, in[v], out[v])
	}
	// 출력:
	// 노드 1의 서브트리 합: 22 (구간 [0, 5])
	// 노드 2의 서브트리 합: 16 (구간 [1, 3])
	// 노드 3의 서브트리 합: 3  (구간 [4, 5])
	// 노드 4의 서브트리 합: 2  (구간 [5, 5])
	// 노드 5의 서브트리 합: 4  (구간 [2, 2])
	// 노드 6의 서브트리 합: 7  (구간 [3, 3])

	// 서브트리 포함 판별
	fmt.Println("\n=== 서브트리 포함 판별 ===")
	// 노드 5가 노드 2의 서브트리에 속하는가?
	u, v := 5, 2
	if in[v] <= in[u] && in[u] <= out[v] {
		fmt.Printf("노드 %d는 노드 %d의 서브트리에 속한다\n", u, v)
	}
	// 노드 4가 노드 2의 서브트리에 속하는가?
	u, v = 4, 2
	if in[v] <= in[u] && in[u] <= out[v] {
		fmt.Printf("노드 %d는 노드 %d의 서브트리에 속한다\n", u, v)
	} else {
		fmt.Printf("노드 %d는 노드 %d의 서브트리에 속하지 않는다\n", u, v)
	}
}
