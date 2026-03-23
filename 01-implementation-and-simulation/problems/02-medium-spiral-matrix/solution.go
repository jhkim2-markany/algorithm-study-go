package main

import (
	"bufio"
	"fmt"
	"os"
)

// buildSpiralMatrix는 N×N 크기의 행렬을 나선형(spiral) 순서로
// 1부터 N²까지 채워 반환한다.
//
// [매개변수]
//   - n: 행렬의 크기 (N×N)
//
// [반환값]
//   - [][]int: 나선형 순서로 채워진 N×N 행렬
func buildSpiralMatrix(n int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 행렬 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	matrix := buildSpiralMatrix(n)

	// 결과 출력
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, matrix[i][j])
		}
		fmt.Fprintln(writer)
	}
}
