package main

import (
	"bufio"
	"fmt"
	"os"
)

// printShortestPath는 레드 나이트의 최단 경로를 출력한다.
//
// [매개변수]
//   - n: 체스판 크기
//   - rStart, cStart: 시작 위치
//   - rEnd, cEnd: 목표 위치
//   - writer: 출력 버퍼
//
// [반환값]
//   - 없음 (표준 출력으로 최소 이동 횟수와 경로를 출력)
func printShortestPath(n int, rStart, cStart, rEnd, cEnd int, writer *bufio.Writer) {
	// 여기에 코드를 작성하세요
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var rStart, cStart, rEnd, cEnd int
	fmt.Fscan(reader, &rStart, &cStart)
	fmt.Fscan(reader, &rEnd, &cEnd)

	printShortestPath(n, rStart, cStart, rEnd, cEnd, writer)
}
