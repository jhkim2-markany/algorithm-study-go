package main

import (
	"bufio"
	"fmt"
	"os"
)

// processSetOperations는 비트마스크 집합 연산을 처리하여 최종 집합을 반환한다.
//
// [매개변수]
//   - n: 전체 집합 크기
//   - operations: 연산 목록 (각 연산은 문자열 슬라이스)
//
// [반환값]
//   - int: 최종 비트마스크 상태
func processSetOperations(n int, operations [][]string) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	mask := 0

	for i := 0; i < q; i++ {
		var op string
		fmt.Fscan(reader, &op)

		switch op {
		case "add":
			var x int
			fmt.Fscan(reader, &x)
			mask |= 1 << (x - 1)

		case "remove":
			var x int
			fmt.Fscan(reader, &x)
			mask &^= 1 << (x - 1)

		case "toggle":
			var x int
			fmt.Fscan(reader, &x)
			mask ^= 1 << (x - 1)

		case "union":
			var k int
			fmt.Fscan(reader, &k)
			other := 0
			for j := 0; j < k; j++ {
				var x int
				fmt.Fscan(reader, &x)
				other |= 1 << (x - 1)
			}
			mask |= other

		case "intersect":
			var k int
			fmt.Fscan(reader, &k)
			other := 0
			for j := 0; j < k; j++ {
				var x int
				fmt.Fscan(reader, &x)
				other |= 1 << (x - 1)
			}
			mask &= other
		}
	}

	// 최종 집합 출력
	if mask == 0 {
		fmt.Fprintln(writer, "empty")
		return
	}

	first := true
	for i := 0; i < n; i++ {
		if mask&(1<<i) != 0 {
			if !first {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, i+1)
			first = false
		}
	}
	fmt.Fprintln(writer)
}
