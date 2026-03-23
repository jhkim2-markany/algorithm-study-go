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
func topologicalSort(adj [][]int, inDegree []int, n int) []int {
	// 여기에 코드를 작성하세요
	return nil
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
