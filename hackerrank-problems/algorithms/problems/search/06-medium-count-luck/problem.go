package main

import (
	"bufio"
	"fmt"
	"os"
)

// countLuck은 시작점에서 포트키까지 경로의 갈림길 수가 k와 같은지 판별한다.
//
// [매개변수]
//   - matrix: N×M 격자 ('.' 이동 가능, 'X' 나무, 'M' 시작, '*' 포트키)
//   - k: 예상 갈림길 수
//
// [반환값]
//   - string: "Impressed" 또는 "Oops!"
func countLuck(matrix []string, k int) string {
	// 여기에 코드를 작성하세요
	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)

		matrix := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &matrix[i])
		}

		var k int
		fmt.Fscan(reader, &k)

		fmt.Fprintln(writer, countLuck(matrix, k))
	}
}
