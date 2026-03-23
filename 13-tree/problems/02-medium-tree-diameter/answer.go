package main

import (
	"bufio"
	"fmt"
	"os"
)

// Edge는 가중치가 있는 간선을 나타낸다.
type Edge struct {
	to, weight int
}

// treeDiameter는 가중치가 있는 트리의 지름(가장 먼 두 노드 사이의 거리)을 반환한다.
//
// [매개변수]
//   - adj: 가중치 간선의 인접 리스트 (1-indexed)
//   - n: 노드 수
//
// [반환값]
//   - int: 트리의 지름
//
// [알고리즘 힌트]
//
//	두 번의 BFS를 사용한다.
//	1단계: 임의의 노드(1번)에서 BFS로 가장 먼 노드를 찾는다.
//	2단계: 찾은 노드에서 다시 BFS로 가장 먼 노드를 찾으면 그 거리가 지름이다.
func treeDiameter(adj [][]Edge, n int) int {
	bfs := func(start int) (int, []int) {
		dist := make([]int, n+1)
		for i := 0; i <= n; i++ {
			dist[i] = -1
		}
		dist[start] = 0
		queue := []int{start}
		farthest := start

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, e := range adj[cur] {
				if dist[e.to] == -1 {
					dist[e.to] = dist[cur] + e.weight
					queue = append(queue, e.to)
					if dist[e.to] > dist[farthest] {
						farthest = e.to
					}
				}
			}
		}
		return farthest, dist
	}

	far1, _ := bfs(1)
	far2, dist := bfs(far1)
	return dist[far2]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 인접 리스트 초기화
	adj := make([][]Edge, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []Edge{}
	}

	// 간선 입력 (가중치 포함)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], Edge{v, w})
		adj[v] = append(adj[v], Edge{u, w})
	}

	// 핵심 함수 호출
	result := treeDiameter(adj, n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
