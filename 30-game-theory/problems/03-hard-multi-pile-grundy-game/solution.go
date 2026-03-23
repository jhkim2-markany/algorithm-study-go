package main

import (
	"bufio"
	"fmt"
	"os"
)

// mex는 집합에 포함되지 않는 가장 작은 음이 아닌 정수를 구한다
func mex(set map[int]bool) int {
	val := 0
	for set[val] {
		val++
	}
	return val
}

// multiPileGrundyWinner는 Sprague-Grundy 정리로 다중 더미 게임의 승자를 판별한다.
//
// [매개변수]
//   - piles: 각 더미의 돌 개수 배열
//   - moves: 한 번에 가져갈 수 있는 돌의 개수 집합
//
// [반환값]
//   - string: 선수 승리이면 "First", 후수 승리이면 "Second"
func multiPileGrundyWinner(piles []int, moves []int) string {
	// 여기에 코드를 작성하세요
	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m, k int
	fmt.Fscan(reader, &m, &k)

	moves := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &moves[i])
	}

	piles := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &piles[i])
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, multiPileGrundyWinner(piles, moves))
}
