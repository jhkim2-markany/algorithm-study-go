package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1

// 간선 구조체: 도착 노드와 색상 (0: R, 1: B)
type Edge struct {
	to, color int
}

// minColorChangePath는 노드 1에서 노드 N까지 색상 변경 최소 비용 경로를 구한다.
// 같은 색 간선 이동은 비용 0, 다른 색 간선 이동은 비용 1이며,
// 상태 (노드, 현재 색상)에 대해 덱 기반 0-1 BFS를 사용한다.
//
// [매개변수]
//   - n: 노드 수
//   - adj: 인접 리스트 (adj[i]는 노드 i에서 나가는 간선 목록)
//
// [반환값]
//   - int: 노드 1에서 노드 N까지의 최소 색상 변경 비용 (도달 불가 시 -1)
func minColorChangePath(n int, adj [][]Edge) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수, 간선 수
	var n, m int
	fmt.Fscan(reader, &n, &m)

	adj := make([][]Edge, n+1)

	// 입력: 간선 정보
	for i := 0; i < m; i++ {
		var u, v int
		var c string
		fmt.Fscan(reader, &u, &v, &c)
		color := 0 // R
		if c == "B" {
			color = 1 // B
		}
		adj[u] = append(adj[u], Edge{v, color})
		adj[v] = append(adj[v], Edge{u, color})
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, minColorChangePath(n, adj))
}
