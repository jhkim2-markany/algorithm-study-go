package main

import (
	"bufio"
	"fmt"
	"os"
)

// stoneTakingWinner는 돌 가져가기 게임의 승자를 판별한다.
//
// [매개변수]
//   - n: 돌의 개수
//   - moves: 한 번에 가져갈 수 있는 돌의 개수 집합
//
// [반환값]
//   - string: 선수 승리이면 "First", 후수 승리이면 "Second"
func stoneTakingWinner(n int, moves []int) string {
	// 여기에 코드를 작성하세요
	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	moves := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &moves[i])
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, stoneTakingWinner(n, moves))
}
