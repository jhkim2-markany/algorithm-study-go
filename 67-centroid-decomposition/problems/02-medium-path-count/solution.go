package main

import (
	"bufio"
	"fmt"
	"os"
)

// countExactDistancePairs는 센트로이드 분할과 해시맵을 이용하여
// 트리에서 정확히 거리 K인 노드 쌍의 수를 구한다.
//
// [매개변수]
//   - n: 노드 수
//   - k: 목표 거리
//   - edges: 가중치 간선 목록 (u, v, w)
//
// [반환값]
//   - int64: 정확히 거리 K인 노드 쌍의 수
func countExactDistancePairs(n, k int, edges [][3]int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	edges := make([][3]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	fmt.Fprintln(writer, countExactDistancePairs(n, k, edges))
}
