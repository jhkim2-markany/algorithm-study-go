package main

import (
	"bufio"
	"fmt"
	"os"
)

// roadsAndLibraries는 모든 도시에서 도서관에 접근 가능하게 하는 최소 비용을 반환한다.
//
// [매개변수]
//   - n: 도시 수
//   - cLib: 도서관 건설 비용
//   - cRoad: 도로 수리 비용
//   - edges: 도로 목록 (각 원소는 [2]int{u, v})
//
// [반환값]
//   - int64: 최소 비용
//
// [알고리즘 힌트]
//
//	도로 비용 >= 도서관 비용이면 모든 도시에 도서관을 건설한다.
//	그렇지 않으면 BFS로 연결 요소를 구하고,
//	각 연결 요소에 도서관 1개 + (크기-1)개 도로를 배치한다.
func roadsAndLibraries(n int, cLib int, cRoad int, edges [][2]int) int64 {
	// 도로 비용이 도서관 비용 이상이면 모든 도시에 도서관 건설
	if cRoad >= cLib {
		return int64(n) * int64(cLib)
	}

	// 인접 리스트 구성
	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// BFS로 연결 요소 크기 계산
	visited := make([]bool, n+1)
	var totalCost int64

	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}
		// BFS로 연결 요소 크기 계산
		size := 0
		queue := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			size++
			for _, next := range adj[cur] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		// 도서관 1개 + (크기-1)개 도로
		totalCost += int64(cLib) + int64(size-1)*int64(cRoad)
	}

	return totalCost
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	for ; q > 0; q-- {
		var n, m, cLib, cRoad int
		fmt.Fscan(reader, &n, &m, &cLib, &cRoad)

		edges := make([][2]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &edges[i][0], &edges[i][1])
		}

		result := roadsAndLibraries(n, cLib, cRoad, edges)
		fmt.Fprintln(writer, result)
	}
}
