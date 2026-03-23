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
//
// [알고리즘 힌트]
//
//	상태를 (노드, 현재 색상)으로 확장하여 0-1 BFS를 수행한다.
//	dist[node][color]는 해당 상태로 도달하는 최소 비용을 저장한다.
//	같은 색 간선 이동(비용 0)은 덱 앞에, 다른 색 간선 이동(비용 1)은
//	덱 뒤에 추가한다. 시작 시 노드 1에서 R(0), B(1) 두 상태 모두
//	비용 0으로 초기화한다. 상태를 node*2+color로 인코딩하여 덱에 저장한다.
func minColorChangePath(n int, adj [][]Edge) int {
	dist := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i][0] = INF
		dist[i][1] = INF
	}

	dist[1][0] = 0
	dist[1][1] = 0

	deque := []int{1*2 + 0, 1*2 + 1}

	for len(deque) > 0 {
		cur := deque[0]
		deque = deque[1:]

		v, curColor := cur/2, cur%2

		for _, e := range adj[v] {
			var w int
			var nextColor int
			if e.color == curColor {
				w = 0
				nextColor = curColor
			} else {
				w = 1
				nextColor = e.color
			}

			newDist := dist[v][curColor] + w
			if newDist < dist[e.to][nextColor] {
				dist[e.to][nextColor] = newDist
				encoded := e.to*2 + nextColor
				if w == 0 {
					deque = append([]int{encoded}, deque...)
				} else {
					deque = append(deque, encoded)
				}
			}
		}
	}

	ans := dist[n][0]
	if dist[n][1] < ans {
		ans = dist[n][1]
	}
	if ans >= INF {
		return -1
	}
	return ans
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
