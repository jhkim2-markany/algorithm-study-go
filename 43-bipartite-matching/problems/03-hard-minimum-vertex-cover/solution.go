package main

import (
	"bufio"
	"fmt"
	"os"
)

// minVertexCover는 이분 그래프에서 쾨니그 정리를 활용하여 최소 버텍스 커버 크기를 구한다.
// 격자에서 장애물이 없는 칸을 행-열 이분 그래프로 모델링한다.
//
// [매개변수]
//   - n: 행의 수
//   - m: 열의 수
//   - blocked: 장애물 위치 맵 (0-indexed, [행][열] → true)
//
// [반환값]
//   - int: 최소 버텍스 커버 크기 (= 최대 매칭 수)
func minVertexCover(n, m int, blocked map[[2]int]bool) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	blocked := make(map[[2]int]bool)
	for i := 0; i < k; i++ {
		var r, c int
		fmt.Fscan(reader, &r, &c)
		blocked[[2]int{r - 1, c - 1}] = true
	}

	fmt.Fprintln(writer, minVertexCover(n, m, blocked))
}
