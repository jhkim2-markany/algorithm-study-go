package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N: 전체 집합 크기, Q: 연산 수
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 비트마스크로 집합을 표현한다 (0번 비트 = 원소 1)
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
			// 집합 A를 비트마스크로 만든 뒤 합집합을 구한다
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
			// 집합 A를 비트마스크로 만든 뒤 교집합을 구한다
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
