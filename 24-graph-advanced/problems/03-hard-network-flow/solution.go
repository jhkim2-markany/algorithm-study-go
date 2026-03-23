package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1<<31 - 1

// maxFlow는 Edmonds-Karp 알고리즘으로 소스에서 싱크까지의 최대 유량을 구한다.
//
// [매개변수]
//   - capacity: 잔여 용량 행렬 (n+1 × n+1)
//   - adjList: 인접 리스트 (역방향 포함)
//   - n: 정점의 수
//   - s: 소스 정점
//   - t: 싱크 정점
//
// [반환값]
//   - int: 최대 유량
func maxFlow(capacity [][]int, adjList [][]int, n, s, t int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, s, t int
	fmt.Fscan(reader, &n, &m, &s, &t)

	// 잔여 용량 행렬 초기화
	capacity := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		capacity[i] = make([]int, n+1)
	}

	// 인접 리스트 초기화
	adjList := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adjList[i] = []int{}
	}

	// 간선 입력
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		capacity[u][v] += c
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	// 핵심 함수 호출
	result := maxFlow(capacity, adjList, n, s, t)

	fmt.Fprintln(writer, result)
}
