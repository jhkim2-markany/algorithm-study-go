package main

import (
	"bufio"
	"fmt"
	"os"
)

// lineContainer는 Li Chao Tree를 이용하여 직선 추가와 최솟값 쿼리를 처리한다.
//
// [매개변수]
//   - ops: 연산 목록 (op=1: 직선 추가 [1, m, b], op=2: 최솟값 쿼리 [2, x])
//
// [반환값]
//   - []int64: 최솟값 쿼리(op=2)의 결과 배열
func lineContainer(ops [][]int64) []int64 {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	ops := make([][]int64, q)
	for i := 0; i < q; i++ {
		var op int64
		fmt.Fscan(reader, &op)
		if op == 1 {
			var m, b int64
			fmt.Fscan(reader, &m, &b)
			ops[i] = []int64{op, m, b}
		} else {
			var x int64
			fmt.Fscan(reader, &x)
			ops[i] = []int64{op, x}
		}
	}

	results := lineContainer(ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
