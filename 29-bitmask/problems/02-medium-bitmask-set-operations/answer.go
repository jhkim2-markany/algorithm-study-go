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
//
// [알고리즘 힌트]
//
//	비트마스크로 집합을 표현한다. 원소 x는 (x-1)번 비트에 대응한다.
//	add: OR 연산으로 비트 설정, remove: AND NOT으로 비트 해제,
//	toggle: XOR로 비트 반전, union: OR로 합집합,
//	intersect: AND로 교집합을 구한다.
func processSetOperations(n int, operations [][]string) int {
	// 이 문제는 I/O 처리가 복잡하여 main에서 직접 처리한다
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 비트마스크로 집합을 표현한다
	mask := 0

	for i := 0; i < q; i++ {
		var op string
		fmt.Fscan(reader, &op)

		switch op {
		case "add":
			var x int
			fmt.Fscan(reader, &x)
			// x번째 원소를 추가한다 (x-1번 비트를 설정)
			mask |= 1 << (x - 1)

		case "remove":
			var x int
			fmt.Fscan(reader, &x)
			// x번째 원소를 제거한다 (x-1번 비트를 해제)
			mask &^= 1 << (x - 1)

		case "toggle":
			var x int
			fmt.Fscan(reader, &x)
			// x번째 원소를 토글한다 (x-1번 비트를 반전)
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
			// OR 연산으로 합집합
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
			// AND 연산으로 교집합
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
