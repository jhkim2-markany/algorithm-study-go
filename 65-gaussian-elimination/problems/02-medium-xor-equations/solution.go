package main

import (
	"bufio"
	"fmt"
	"os"
)

// solveXorEquations는 GF(2) 위의 XOR 연립방정식을 가우스 소거법으로 풀어
// 해를 구한다. 해가 없으면 nil을 반환한다.
//
// [매개변수]
//   - n: 미지수 수
//   - m: 방정식 수
//   - equations: 각 방정식의 변수 인덱스 목록과 결과값
//     (equations[i] = {indices: 변수 인덱스들(1-indexed), b: 결과값})
//
// [반환값]
//   - []int: 해 벡터 (자유 변수는 0), 해가 없으면 nil
func solveXorEquations(n, m int, indices [][]int, b []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	indices := make([][]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		var k int
		fmt.Fscan(reader, &k)
		indices[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &indices[i][j])
		}
		fmt.Fscan(reader, &b[i])
	}

	x := solveXorEquations(n, m, indices, b)
	if x == nil {
		fmt.Fprintln(writer, "IMPOSSIBLE")
	} else {
		for i := 0; i < n; i++ {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, x[i])
		}
		fmt.Fprintln(writer)
	}
}
