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

	// 행렬 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 행렬 입력
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	// K 입력
	var k int
	fmt.Fscan(reader, &k)

	// 이진 탐색으로 K번째 수를 찾는다
	// 탐색 범위: 행렬의 최솟값(좌상단) ~ 최댓값(우하단)
	lo, hi := matrix[0][0], matrix[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo)/2

		// mid 이하인 원소의 개수를 센다
		count := countLessOrEqual(matrix, n, mid)

		if count < k {
			// mid 이하인 원소가 K개 미만이면 답은 mid보다 크다
			lo = mid + 1
		} else {
			// mid 이하인 원소가 K개 이상이면 답은 mid 이하이다
			hi = mid
		}
	}

	fmt.Fprintln(writer, lo)
}

// countLessOrEqual은 행렬에서 val 이하인 원소의 개수를 반환한다.
// 각 행이 정렬되어 있으므로 행마다 upper_bound를 이진 탐색으로 구한다.
// 시간 복잡도: O(N × log N)
func countLessOrEqual(matrix [][]int, n int, val int) int {
	count := 0
	for i := 0; i < n; i++ {
		// i번째 행에서 val 이하인 원소의 개수
		lo, hi := 0, n
		for lo < hi {
			mid := (lo + hi) / 2
			if matrix[i][mid] <= val {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		count += lo
	}
	return count
}
