package main

import (
	"bufio"
	"fmt"
	"os"
)

// nimGameWinner는 님 게임의 승자를 판별한다.
//
// [매개변수]
//   - piles: 각 돌 더미의 돌 개수 배열
//
// [반환값]
//   - string: 선수 승리이면 "First", 후수 승리이면 "Second"
//
// [알고리즘 힌트]
//
//	님 게임의 핵심 정리: 모든 더미의 돌 개수를 XOR한 값이
//	0이 아니면 선수 승리, 0이면 후수 승리이다.
//	XOR 합이 0이 아닌 상태에서는 항상 XOR 합을 0으로 만드는
//	수가 존재하므로 선수가 유리하다.
func nimGameWinner(piles []int) string {
	xorSum := 0
	for _, a := range piles {
		xorSum ^= a
	}

	if xorSum != 0 {
		return "First"
	}
	return "Second"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	piles := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &piles[i])
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, nimGameWinner(piles))
}
