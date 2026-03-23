package main

import (
	"bufio"
	"fmt"
	"os"
)

var writer *bufio.Writer

// 하노이의 탑 재귀 함수
// n개의 원판을 from 기둥에서 to 기둥으로 옮긴다 (aux를 보조 기둥으로 사용)
func hanoi(n, from, to, aux int) {
	// 기저 조건: 원판이 1개이면 바로 옮긴다
	if n == 1 {
		fmt.Fprintf(writer, "%d %d\n", from, to)
		return
	}
	// 1단계: 위의 n-1개 원판을 보조 기둥으로 옮긴다
	hanoi(n-1, from, aux, to)
	// 2단계: 가장 큰 원판을 목표 기둥으로 옮긴다
	fmt.Fprintf(writer, "%d %d\n", from, to)
	// 3단계: 보조 기둥의 n-1개 원판을 목표 기둥으로 옮긴다
	hanoi(n-1, aux, to, from)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 원판의 수
	var n int
	fmt.Fscan(reader, &n)

	// 총 이동 횟수: 2^N - 1
	moves := 1
	for i := 0; i < n; i++ {
		moves *= 2
	}
	moves--
	fmt.Fprintln(writer, moves)

	// 하노이의 탑 실행 (기둥 1 → 기둥 3, 보조 기둥 2)
	hanoi(n, 1, 3, 2)
}
