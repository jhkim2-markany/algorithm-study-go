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
func hanoi(n, from, to, aux int, moves *[][2]int) {
	// 여기에 코드를 작성하세요
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
