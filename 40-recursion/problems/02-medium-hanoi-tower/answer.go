package main

import (
	"bufio"
	"fmt"
	"os"
)

// hanoi는 하노이의 탑 문제를 재귀적으로 풀어 이동 과정을 moves 슬라이스에 기록한다.
//
// [매개변수]
//   - n: 원판의 수
//   - from: 출발 기둥 번호
//   - to: 목표 기둥 번호
//   - aux: 보조 기둥 번호
//   - moves: 이동 기록을 저장할 슬라이스 포인터 (각 원소는 [2]int{출발, 도착})
//
// [반환값]
//   - 없음 (moves 슬라이스에 결과를 기록)
//
// [알고리즘 힌트]
//   1. 기저 조건: n = 1이면 from → to 이동을 기록하고 반환한다.
//   2. 위의 n-1개 원판을 from → aux로 옮긴다 (to를 보조로 사용).
//   3. 가장 큰 원판을 from → to로 옮긴다.
//   4. aux에 있는 n-1개 원판을 aux → to로 옮긴다 (from을 보조로 사용).
func hanoi(n, from, to, aux int, moves *[][2]int) {
	if n == 1 {
		*moves = append(*moves, [2]int{from, to})
		return
	}
	hanoi(n-1, from, aux, to, moves)
	*moves = append(*moves, [2]int{from, to})
	hanoi(n-1, aux, to, from, moves)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	// 총 이동 횟수: 2^N - 1
	total := 1
	for i := 0; i < n; i++ {
		total *= 2
	}
	total--
	fmt.Fprintln(writer, total)

	moves := make([][2]int, 0, total)
	hanoi(n, 1, 3, 2, &moves)

	for _, m := range moves {
		fmt.Fprintf(writer, "%d %d\n", m[0], m[1])
	}
}
