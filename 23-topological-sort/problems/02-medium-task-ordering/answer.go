package main

import (
	"bufio"
	"fmt"
	"os"
)

// topologicalSort는 위상 정렬을 수행하여 작업 순서를 구한다.
//
// [매개변수]
//   - adj: 인접 리스트로 표현된 방향 그래프 (1-indexed)
//   - inDegree: 각 정점의 진입 차수 배열
//   - n: 정점(작업)의 수
//
// [반환값]
//   - []int: 위상 정렬 결과 (사이클이 있으면 nil)
//
// [알고리즘 힌트]
//
//	Kahn 알고리즘으로 위상 정렬을 수행한다.
//	진입 차수가 0인 정점을 큐에 넣고, 큐에서 꺼낸 정점을 결과에 추가한다.
//	인접 정점의 진입 차수를 감소시키고, 0이 되면 큐에 추가한다.
//	결과 배열의 길이가 N이 아니면 사이클이 존재한다.
func topologicalSort(adj [][]int, inDegree []int, n int) []int {
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0, n)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)

		for _, next := range adj[cur] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(result) != n {
		return nil
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 인접 리스트와 진입 차수 배열 초기화
	adj := make([][]int, n+1)
	inDegree := make([]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = []int{}
	}

	// 선행 관계 입력
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		adj[a] = append(adj[a], b)
		inDegree[b]++
	}

	// 핵심 함수 호출
	result := topologicalSort(adj, inDegree, n)

	if result == nil {
		fmt.Fprintln(writer, -1)
		return
	}

	// 위상 정렬 결과 출력
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
