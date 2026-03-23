package main

import (
	"bufio"
	"fmt"
	"os"
)

// mex: 집합에 포함되지 않는 가장 작은 음이 아닌 정수를 구한다
func mex(set map[int]bool) int {
	val := 0
	for set[val] {
		val++
	}
	return val
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 게임 수 M과 이동 종류 수 K 입력
	var m, k int
	fmt.Fscan(reader, &m, &k)

	// 가져갈 수 있는 돌의 개수 집합 입력
	moves := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &moves[i])
	}

	// 각 게임의 돌 개수 입력
	piles := make([]int, m)
	maxN := 0
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &piles[i])
		if piles[i] > maxN {
			maxN = piles[i]
		}
	}

	// 0부터 maxN까지 그런디 수를 계산한다
	grundy := make([]int, maxN+1)
	// grundy[0] = 0 (돌 0개: 이동 불가, 그런디 수 0)

	for i := 1; i <= maxN; i++ {
		// 이동 가능한 상태들의 그런디 수를 수집한다
		reachable := make(map[int]bool)
		for _, mv := range moves {
			if i >= mv {
				reachable[grundy[i-mv]] = true
			}
		}
		// mex를 적용하여 그런디 수를 결정한다
		grundy[i] = mex(reachable)
	}

	// Sprague-Grundy 정리: 각 게임의 그런디 수를 XOR한다
	xorSum := 0
	for _, p := range piles {
		xorSum ^= grundy[p]
	}

	// XOR 합이 0이 아니면 선수 승리, 0이면 후수 승리
	if xorSum != 0 {
		fmt.Fprintln(writer, "First")
	} else {
		fmt.Fprintln(writer, "Second")
	}
}
