package main

import "fmt"

// 센트로이드 분할 (Centroid Decomposition)
// 트리에서 센트로이드를 반복적으로 찾아 제거하면서 센트로이드 트리를 구축한다.
// 시간 복잡도: O(N log N)
// 공간 복잡도: O(N)

const MAXN = 100001

var (
	adj     [MAXN][]int // 원래 트리의 인접 리스트
	sz      [MAXN]int   // 서브트리 크기
	removed [MAXN]bool  // 센트로이드로 제거된 노드 표시
	cPar    [MAXN]int   // 센트로이드 트리에서의 부모
	n       int         // 노드 수
)

// calcSize는 현재 서브트리의 크기를 계산한다
func calcSize(v, p int) int {
	sz[v] = 1
	for _, u := range adj[v] {
		if u == p || removed[u] {
			continue // 부모 방향 또는 이미 제거된 노드 건너뛰기
		}
		sz[v] += calcSize(u, v)
	}
	return sz[v]
}

// findCentroid는 크기가 treeSize인 서브트리에서 센트로이드를 찾는다
func findCentroid(v, p, treeSize int) int {
	for _, u := range adj[v] {
		if u == p || removed[u] {
			continue
		}
		// 자식 서브트리 크기가 전체의 절반을 초과하면 그 방향으로 이동
		if sz[u] > treeSize/2 {
			return findCentroid(u, v, treeSize)
		}
	}
	return v
}

// decompose는 센트로이드 분할을 재귀적으로 수행한다
func decompose(v, parent int) {
	// 1. 현재 서브트리 크기 계산
	treeSize := calcSize(v, -1)

	// 2. 센트로이드 찾기
	centroid := findCentroid(v, -1, treeSize)

	// 3. 센트로이드 트리에서 부모 설정
	cPar[centroid] = parent

	// 4. 센트로이드 제거 표시
	removed[centroid] = true

	// 5. 각 서브트리에 대해 재귀적으로 분할
	for _, u := range adj[centroid] {
		if removed[u] {
			continue
		}
		decompose(u, centroid)
	}
}

func main() {
	// 예시 트리:
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6
	//       |
	//       7
	n = 7
	edges := [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {5, 7}}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 센트로이드 분할 수행 (루트의 부모는 0으로 표시)
	decompose(1, 0)

	// 센트로이드 트리 출력
	fmt.Println("센트로이드 트리 (각 노드의 부모):")
	for i := 1; i <= n; i++ {
		if cPar[i] == 0 {
			fmt.Printf("  노드 %d: 루트\n", i)
		} else {
			fmt.Printf("  노드 %d: 부모 = %d\n", i, cPar[i])
		}
	}
	// 출력 예시:
	// 센트로이드 트리 (각 노드의 부모):
	//   노드 1: 부모 = 3
	//   노드 2: 루트
	//   노드 3: 부모 = 2
	//   노드 4: 부모 = 2
	//   노드 5: 부모 = 2
	//   노드 6: 부모 = 3
	//   노드 7: 부모 = 5
}
