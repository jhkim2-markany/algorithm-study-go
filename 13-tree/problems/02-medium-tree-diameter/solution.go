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

var adj [][]edge
var dist []int

// bfs 함수는 시작 노드에서 모든 노드까지의 거리를 계산한다
func bfs(start, n int) int {
	dist = make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = -1
	}
	dist[start] = 0

	// 큐를 사용한 BFS 탐색
	queue := []int{start}
	farthest := start

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, e := range adj[cur] {
			if dist[e.to] == -1 {
				dist[e.to] = dist[cur] + e.weight
				queue = append(queue, e.to)
				// 가장 먼 노드를 갱신
				if dist[e.to] > dist[farthest] {
					farthest = e.to
				}
			}
		}
	}
	return farthest
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 인접 리스트 초기화
	adj = make([][]edge, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []edge{}
	}

	// 간선 입력 (가중치 포함)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], edge{v, w})
		adj[v] = append(adj[v], edge{u, w})
	}

	// 트리의 지름 구하기: 두 번의 BFS 사용
	// 1단계: 임의의 노드(1번)에서 가장 먼 노드를 찾는다
	far1 := bfs(1, n)

	// 2단계: 찾은 노드에서 다시 가장 먼 노드를 찾으면 그 거리가 지름이다
	far2 := bfs(far1, n)

	// 지름 출력
	fmt.Fprintln(writer, dist[far2])
}
