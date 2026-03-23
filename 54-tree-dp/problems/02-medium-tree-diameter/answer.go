package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to, weight int
}

// treeDiameter는 가중치 트리의 지름(가장 먼 두 노드 사이의 거리)을 반환한다.
//
// [매개변수]
//   - n: 노드의 수
//   - adj: 가중치 간선의 인접 리스트
//
// [반환값]
//   - int: 트리의 지름
//
// [알고리즘 힌트]
//   - 후위 순회(DFS)로 각 노드에서 서브트리 내 가장 먼 리프까지의 거리를 계산한다
//   - 각 노드에서 가장 긴 두 경로(max1, max2)를 추적한다
//   - 현재 노드를 꺾는 점으로 하는 경로 = max1 + max2가 지름 후보이다
//   - 모든 노드에서의 후보 중 최댓값이 트리의 지름이다
func treeDiameter(n int, adj [][]edge) int {
	depth := make([]int, n+1)
	diameter := 0

	var dfs func(v, parent int)
	dfs = func(v, parent int) {
		depth[v] = 0
		max1, max2 := 0, 0

		for _, e := range adj[v] {
			if e.to == parent {
				continue
			}
			dfs(e.to, v)

			childDist := depth[e.to] + e.weight
			if childDist >= max1 {
				max2 = max1
				max1 = childDist
			} else if childDist > max2 {
				max2 = childDist
			}
		}

		depth[v] = max1
		if max1+max2 > diameter {
			diameter = max1 + max2
		}
	}

	dfs(1, 0)
	return diameter
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	adj := make([][]edge, n+1)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}

	fmt.Fprintln(writer, treeDiameter(n, adj))
}
