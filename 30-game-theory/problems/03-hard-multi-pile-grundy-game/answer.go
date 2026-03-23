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
//
// [알고리즘 힌트]
//
//	Sprague-Grundy 정리를 사용한다.
//	각 더미에 대해 그런디 수를 계산한다:
//	  grundy[i] = mex({grundy[i-m] | m ∈ moves, i ≥ m})
//	  mex는 집합에 포함되지 않는 가장 작은 음이 아닌 정수이다.
//	모든 더미의 그런디 수를 XOR한 값이 0이 아니면 선수 승리이다.
func multiPileGrundyWinner(piles []int, moves []int) string {
	maxN := 0
	for _, p := range piles {
		if p > maxN {
			maxN = p
		}
	}

	grundy := make([]int, maxN+1)

	for i := 1; i <= maxN; i++ {
		reachable := make(map[int]bool)
		for _, mv := range moves {
			if i >= mv {
				reachable[grundy[i-mv]] = true
			}
		}
		grundy[i] = mex(reachable)
	}

	xorSum := 0
	for _, p := range piles {
		xorSum ^= grundy[p]
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
