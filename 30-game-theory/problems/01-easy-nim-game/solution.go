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
func nimGameWinner(piles []int) string {
	// 여기에 코드를 작성하세요
	return ""
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
