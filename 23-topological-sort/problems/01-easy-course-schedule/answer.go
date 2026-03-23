package main

import (
	"bufio"
	"fmt"
	"os"
)

// canFinishAll은 위상 정렬을 수행하여 모든 과목을 수강할 수 있는지 판별한다.
//
// [매개변수]
//   - adj: 인접 리스트로 표현된 방향 그래프 (1-indexed)
//   - inDegree: 각 정점의 진입 차수 배열
//   - n: 정점(과목)의 수
//
// [반환값]
//   - bool: 모든 과목을 수강할 수 있으면 true, 사이클이 있으면 false
//
// [알고리즘 힌트]
//
//	Kahn 알고리즘으로 위상 정렬을 수행한다.
//	진입 차수가 0인 정점을 큐에 넣고, 큐에서 꺼낸 정점의 인접 정점의
//	진입 차수를 감소시킨다. 진입 차수가 0이 되면 큐에 추가한다.
//	처리된 정점 수가 N이면 사이클이 없다.
func canFinishAll(adj [][]int, inDegree []int, n int) bool {
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++

		for _, next := range adj[cur] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return count == n
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

	// 선수 과목 관계 입력
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		adj[a] = append(adj[a], b)
		inDegree[b]++
	}

	// 핵심 함수 호출
	if canFinishAll(adj, inDegree, n) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
