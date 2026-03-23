package main

import (
	"bufio"
	"fmt"
	"os"
)

// 색상 변경 최소 비용 경로
// 상태: (노드, 현재 색상) → 0-1 BFS로 최단 거리 계산
// 같은 색 간선 이동: 비용 0, 다른 색 간선 이동: 비용 1 (색상 변경)

const INF = 1<<31 - 1

// 간선 구조체: 도착 노드와 색상 (0: R, 1: B)
type Edge struct {
	to, color int
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

	// dist[node][color]: 노드 node에 색상 color 상태로 도달하는 최소 비용
	// color: 0 = R, 1 = B
	dist := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i][0] = INF
		dist[i][1] = INF
	}

	// 시작: 노드 1에서 R 선택(비용 0) 또는 B 선택(비용 0)
	dist[1][0] = 0
	dist[1][1] = 0

	// 덱: (노드*2 + 색상)으로 인코딩
	deque := []int{1*2 + 0, 1*2 + 1}

	for len(deque) > 0 {
		// 덱 앞에서 꺼내기
		cur := deque[0]
		deque = deque[1:]

		v, curColor := cur/2, cur%2

		for _, e := range adj[v] {
			var w int
			var nextColor int
			if e.color == curColor {
				// 같은 색 간선: 비용 0, 색상 유지
				w = 0
				nextColor = curColor
			} else {
				// 다른 색 간선: 비용 1, 색상 변경
				w = 1
				nextColor = e.color
			}

			newDist := dist[v][curColor] + w
			if newDist < dist[e.to][nextColor] {
				dist[e.to][nextColor] = newDist
				encoded := e.to*2 + nextColor
				if w == 0 {
					// 비용 0: 덱 앞에 추가
					deque = append([]int{encoded}, deque...)
				} else {
					// 비용 1: 덱 뒤에 추가
					deque = append(deque, encoded)
				}
			}
		}
	}

	// 결과: 노드 N에 도달하는 최소 비용 (R 또는 B 중 작은 값)
	ans := dist[n][0]
	if dist[n][1] < ans {
		ans = dist[n][1]
	}
	if ans >= INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}
