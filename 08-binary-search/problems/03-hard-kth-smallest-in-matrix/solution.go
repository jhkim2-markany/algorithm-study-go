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
func kthSmallest(matrix [][]int, k int) int {
	// 여기에 코드를 작성하세요
	return 0
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
