package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Operation은 비트마스크 집합 연산을 나타낸다
type Operation struct {
	Op       string
	Operands []int
}

// processSetOperations는 비트마스크 집합 연산을 처리하여 최종 비트마스크를 반환한다.
//
// [매개변수]
//   - n: 전체 집합 크기
//   - operations: 연산 목록 (각 연산은 연산자와 피연산자를 포함)
//
// [반환값]
//   - int: 최종 비트마스크 상태
func processSetOperations(n int, operations []Operation) int {
	// 여기에 코드를 작성하세요
	return 0
}

// formatSet은 비트마스크를 집합 문자열로 변환한다
func formatSet(mask, n int) string {
	if mask == 0 {
		return "empty"
	}
	result := ""
	for i := 0; i < n; i++ {
		if mask&(1<<i) != 0 {
			if result != "" {
				result += " "
			}
			result += strconv.Itoa(i + 1)
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 연산 목록 입력
	operations := make([]Operation, q)
	for i := 0; i < q; i++ {
		var op string
		fmt.Fscan(reader, &op)

		switch op {
		case "add", "remove", "toggle":
			var x int
			fmt.Fscan(reader, &x)
			operations[i] = Operation{Op: op, Operands: []int{x}}

		case "union", "intersect":
			var k int
			fmt.Fscan(reader, &k)
			operands := make([]int, k)
			for j := 0; j < k; j++ {
				fmt.Fscan(reader, &operands[j])
			}
			operations[i] = Operation{Op: op, Operands: operands}
		}
	}

	// 핵심 함수 호출
	mask := processSetOperations(n, operations)

	// 결과 출력
	fmt.Fprintln(writer, formatSet(mask, n))
}
