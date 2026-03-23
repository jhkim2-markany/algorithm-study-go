package main

import (
	"bufio"
	"fmt"
	"os"
)

// findArticulationPoints는 무방향 그래프에서 단절점 목록을 오름차순으로 반환한다.
//
// [매개변수]
//   - adj: 인접 리스트 (1-indexed)
//   - n: 정점 수
//
// [반환값]
//   - []int: 단절점 번호의 오름차순 배열
func findArticulationPoints(adj [][]int, n int) []int {
	// 여기에 코드를 작성하세요
	return nil
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
