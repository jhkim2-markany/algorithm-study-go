package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// findArticulationPoints는 무방향 그래프에서 단절점 목록을 오름차순으로 반환한다.
//
// [매개변수]
//   - adj: 인접 리스트 (1-indexed)
//   - n: 정점 수
//
// [반환값]
//   - []int: 단절점 번호의 오름차순 배열
//
// [알고리즘 힌트]
//
//	Tarjan 알고리즘을 사용한다.
//	DFS를 수행하며 각 정점에 방문 순서를 부여하고,
//	서브트리에서 도달 가능한 가장 빠른 방문 순서(low)를 추적한다.
//	루트가 아닌 정점: 자식의 low 값이 자신의 방문 순서 이상이면 단절점이다.
//	루트 정점: DFS 트리에서 자식이 2개 이상이면 단절점이다.
func findArticulationPoints(adj [][]int, n int) []int {
	discovered := make([]int, n+1)
	isAP := make([]bool, n+1)
	timer := 0

	var dfs func(cur, parent int) int
	dfs = func(cur, parent int) int {
		timer++
		discovered[cur] = timer
		minOrder := discovered[cur]
		childCount := 0

		for _, next := range adj[cur] {
			if discovered[next] == 0 {
				childCount++
				low := dfs(next, cur)
				if low < minOrder {
					minOrder = low
				}
				if parent != 0 && low >= discovered[cur] {
					isAP[cur] = true
				}
			} else if next != parent {
				if discovered[next] < minOrder {
					minOrder = discovered[next]
				}
			}
		}

		if parent == 0 && childCount >= 2 {
			isAP[cur] = true
		}

		return minOrder
	}

	dfs(1, 0)

	var result []int
	for i := 1; i <= n; i++ {
		if isAP[i] {
			result = append(result, i)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정점 수와 간선 수 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트 초기화
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = []int{}
	}

	// 간선 입력 (양방향)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 핵심 함수 호출
	result := findArticulationPoints(adj, n)

	// 결과 출력
	fmt.Fprintln(writer, len(result))
	if len(result) > 0 {
		for i, v := range result {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
