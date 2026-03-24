package main

import (
	"bufio"
	"fmt"
	"os"
)

// knightlOnAChessboard는 모든 (a, b) 쌍에 대해 KnightL의 최소 이동 횟수를 반환한다.
//
// [매개변수]
//   - n: 체스판 크기
//
// [반환값]
//   - [][]int: (N-1)×(N-1) 행렬, 각 원소는 최소 이동 횟수 (-1은 도달 불가)
func knightlOnAChessboard(n int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	result := knightlOnAChessboard(n)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, result[i][j])
		}
		fmt.Fprintln(writer)
	}
}
