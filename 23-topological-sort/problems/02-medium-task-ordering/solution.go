package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	// Kahn 알고리즘으로 위상 정렬 수행
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0, n)
	for len(queue) > 0 {
		// 큐에서 정점을 꺼내 결과에 추가
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)

		// 인접 정점의 진입 차수 감소
		for _, next := range adj[cur] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 사이클 판별: 모든 작업이 처리되었는지 확인
	if len(result) != n {
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
