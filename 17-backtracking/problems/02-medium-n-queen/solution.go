package main

import (
	"bufio"
	"fmt"
	"os"
)

// solveNQueen은 N×N 체스판에 N개의 퀸을 서로 공격하지 않게 배치하는 경우의 수를 반환한다.
//
// [매개변수]
//   - n: 체스판의 크기 및 퀸의 개수
//
// [반환값]
//   - int: 가능한 배치의 수
func solveNQueen(n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)

	// 핵심 함수 호출
	result := solveNQueen(n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
