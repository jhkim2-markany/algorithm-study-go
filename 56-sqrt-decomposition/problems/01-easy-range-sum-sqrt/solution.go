package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기와 쿼리 수
	var n, q int
	fmt.Fscan(reader, &n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 블록 크기 결정
	b := int(math.Ceil(math.Sqrt(float64(n))))
	if b == 0 {
		b = 1
	}
	numBlocks := (n + b - 1) / b
	block := make([]int, numBlocks)

	// 각 블록의 합을 미리 계산한다
	for i := 0; i < n; i++ {
		block[i/b] += a[i]
	}

	for ; q > 0; q-- {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			// 점 갱신: a[i-1]을 v로 변경
			var idx, val int
			fmt.Fscan(reader, &idx, &val)
			idx-- // 0-indexed로 변환
			block[idx/b] += val - a[idx]
			a[idx] = val
		} else {
			// 구간 합 쿼리 [l-1, r-1]
			var l, r int
			fmt.Fscan(reader, &l, &r)
			l--
			r--

			sum := 0
			for i := l; i <= r; {
				// 블록 시작이고 블록 끝까지 포함되면 블록 합 사용
				if i%b == 0 && i+b-1 <= r {
					sum += block[i/b]
					i += b
				} else {
					sum += a[i]
					i++
				}
			}
			fmt.Fprintln(writer, sum)
		}
	}
}
