package main

import (
	"bufio"
	"fmt"
	"os"
)

// maxMatching은 이분 그래프에서 헝가리안 알고리즘으로 최대 매칭 수를 구한다.
//
// [매개변수]
//   - n: 왼쪽 정점(학생)의 수
//   - m: 오른쪽 정점(동아리)의 수
//   - adj: 왼쪽 정점별 연결된 오른쪽 정점 목록 (0-indexed)
//
// [반환값]
//   - int: 최대 매칭 수
func maxMatching(n, m int, adj [][]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(reader, &k)
		adj[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &adj[i][j])
			adj[i][j]--
		}
	}

	fmt.Fprintln(writer, maxMatching(n, m, adj))
}
