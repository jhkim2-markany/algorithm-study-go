package main

import (
	"bufio"
	"fmt"
	"os"
)

// kthSmallest는 각 행이 정렬된 N×N 행렬에서 K번째로 작은 수를 반환한다.
//
// [매개변수]
//   - matrix: 각 행이 오름차순 정렬된 N×N 정수 행렬
//   - k: 찾을 순위 (1-indexed)
//
// [반환값]
//   - int: K번째로 작은 수
//
// [알고리즘 힌트]
//
//	값 범위에 대해 이진 탐색을 수행한다.
//	탐색 범위: 행렬 최솟값(좌상단) ~ 최댓값(우하단)
//	mid 값에 대해 행렬에서 mid 이하인 원소의 개수를 센다.
//	각 행이 정렬되어 있으므로 행마다 upper_bound로 O(log N)에 계산.
//	개수 < k이면 lo = mid + 1, 아니면 hi = mid.
//
//	시간복잡도: O(N log N × log(max-min))
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	lo, hi := matrix[0][0], matrix[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo)/2
		count := countLessOrEqual(matrix, n, mid)
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// countLessOrEqual은 행렬에서 val 이하인 원소의 개수를 반환한다.
func countLessOrEqual(matrix [][]int, n int, val int) int {
	count := 0
	for i := 0; i < n; i++ {
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

	// 핵심 함수 호출
	result := kthSmallest(matrix, k)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
