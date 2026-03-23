package main

import (
	"bufio"
	"fmt"
	"os"
)

// 간선 정보를 저장하는 구조체
type edge struct {
	to, weight int
}

var (
	adj      [][]edge // 인접 리스트 (가중치 포함)
	depth    []int    // 각 노드에서 서브트리 내 가장 먼 리프까지의 거리
	diameter int      // 트리의 지름
)

// dfs는 후위 순회로 각 노드의 최대 깊이를 계산하고, 지름 후보를 갱신한다
func dfs(v, parent int) {
	depth[v] = 0
	// 가장 긴 두 경로를 추적한다
	max1, max2 := 0, 0

	for _, e := range adj[v] {
		if e.to == parent {
			continue // 부모 방향 역행 방지
		}
		dfs(e.to, v)

		// 자식을 통한 경로 길이
		childDist := depth[e.to] + e.weight

		// 가장 긴 두 경로 갱신
		if childDist >= max1 {
			max2 = max1
			max1 = childDist
		} else if childDist > max2 {
			max2 = childDist
		}
	}

	depth[v] = max1

	// 현재 노드를 꺾는 점으로 하는 경로 = 가장 긴 두 경로의 합
	if max1+max2 > diameter {
		diameter = max1 + max2
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수
	var n int
	fmt.Fscan(reader, &n)

	adj = make([][]edge, n+1)
	depth = make([]int, n+1)
	diameter = 0

	// 입력: 간선 정보
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}

	// 루트(1번)에서 DFS 수행
	dfs(1, 0)

	// 출력: 트리의 지름
	fmt.Fprintln(writer, diameter)
}
