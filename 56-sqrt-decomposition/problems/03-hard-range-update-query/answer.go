package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// sqrtRangeUpdateQuery는 Sqrt Decomposition(Lazy 블록)을 이용한 구간 갱신 + 구간 합 자료구조를 구현한다.
//
// [매개변수]
//   - n: 배열의 크기
//   - a: 정수 배열 (0-indexed)
//
// [반환값]
//   - rangeUpdate func(l, r, v int): [l, r] 구간에 v를 더하는 함수
//   - rangeQuery func(l, r int) int: [l, r] 구간 합을 반환하는 함수
//
// [알고리즘 힌트]
//
//	배열을 √N 크기의 블록으로 분할하고, 각 블록의 합(blockSum)과 lazy 값을 관리한다.
//	구간 갱신: 완전히 포함되는 블록은 lazy += v (O(1)), 부분 블록은 개별 원소에 v를 더한다.
//	구간 쿼리: 완전 블록은 blockSum + lazy × blockSize, 부분 블록은 a[i] + lazy를 합산한다.
//	각 연산의 시간 복잡도는 O(√N)이다.
func sqrtRangeUpdateQuery(n int, a []int) (rangeUpdate func(l, r, v int), rangeQuery func(l, r int) int) {
	b := int(math.Ceil(math.Sqrt(float64(n))))
	numBlocks := (n + b - 1) / b

	blockSum := make([]int, numBlocks)
	lazy := make([]int, numBlocks)

	// 초기 블록 합 계산
	for i := 0; i < n; i++ {
		blockSum[i/b] += a[i]
	}

	rangeUpdate = func(l, r, v int) {
		lb, rb := l/b, r/b

		if lb == rb {
			// 같은 블록 내
			for i := l; i <= r; i++ {
				a[i] += v
				blockSum[lb] += v
			}
			return
		}

		// 왼쪽 부분 블록
		for i := l; i < (lb+1)*b; i++ {
			a[i] += v
			blockSum[lb] += v
		}

		// 중간 완전 블록
		for bi := lb + 1; bi < rb; bi++ {
			lazy[bi] += v
		}

		// 오른쪽 부분 블록
		for i := rb * b; i <= r; i++ {
			a[i] += v
			blockSum[rb] += v
		}
	}

	rangeQuery = func(l, r int) int {
		lb, rb := l/b, r/b
		sum := 0

		if lb == rb {
			for i := l; i <= r; i++ {
				sum += a[i] + lazy[lb]
			}
			return sum
		}

		// 왼쪽 부분 블록
		for i := l; i < (lb+1)*b; i++ {
			sum += a[i] + lazy[lb]
		}

		// 중간 완전 블록
		for bi := lb + 1; bi < rb; bi++ {
			blockSize := b
			if bi == numBlocks-1 && n%b != 0 {
				blockSize = n % b
			}
			sum += blockSum[bi] + lazy[bi]*blockSize
		}

		// 오른쪽 부분 블록
		for i := rb * b; i <= r; i++ {
			sum += a[i] + lazy[rb]
		}

		return sum
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	rangeUpdate, rangeQuery := sqrtRangeUpdateQuery(n, a)

	for ; q > 0; q-- {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			var l, r, v int
			fmt.Fscan(reader, &l, &r, &v)
			rangeUpdate(l-1, r-1, v)
		} else {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, rangeQuery(l-1, r-1))
		}
	}
}
