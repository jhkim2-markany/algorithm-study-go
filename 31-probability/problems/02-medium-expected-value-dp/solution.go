package main

import (
	"bufio"
	"fmt"
	"os"
)

// expectedMoves는 1번 칸에서 n번 칸까지 도착하기 위한 기대 주사위 횟수를 반환한다.
//
// [매개변수]
//   - n: 보드의 마지막 칸 번호 (2 이상)
//
// [반환값]
//   - float64: 1번 칸에서 출발하여 n번 칸에 도착하기까지의 기대 횟수
func expectedMoves(n int) float64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	fmt.Fprintf(writer, "%.6f\n", expectedMoves(n))
}
