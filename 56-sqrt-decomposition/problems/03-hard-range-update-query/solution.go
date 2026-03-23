package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 구간 갱신 + 구간 합 쿼리를 Sqrt Decomposition (Lazy 블록)으로 처리한다
// 각 블록에 lazy 값을 두어 구간 갱신을 O(√N)에 처리한다
// 시간 복잡도: 구간 갱신 O(√N), 구간 쿼리 O(√N)

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

	// 블록별 합과 lazy 값
	blockSum := make([]int, numBlocks)
	lazy := make([]int, numBlocks)

	// 각 블록의 실제 크기를 계산한다
	blockSize := make([]int, numBlocks)
	for i := 0; i < numBlocks; i++ {
		start := i * b
		end := start + b
		if end > n {
			end = n
		}
		blockSize[i] = end - start
	}

	// 초기 블록 합 계산
	for i := 0; i < n; i++ {
		blockSum[i/b] += a[i]
	}

	// 구간 갱신: [l, r]에 v를 더한다
	rangeUpdate := func(l, r, v int) {
		lb, rb := l/b, r/b

		if lb == rb {
			// 같은 블록 안에 있으면 개별 원소를 갱신한다
			for i := l; i <= r; i++ {
				a[i] += v
				blockSum[lb] += v
			}
			return
		}

		// 왼쪽 부분 블록: 개별 원소 갱신
		for i := l; i < (lb+1)*b; i++ {
			a[i] += v
			blockSum[lb] += v
		}

		// 완전히 포함되는 중간 블록: lazy 값만 갱신
		for bi := lb + 1; bi < rb; bi++ {
			lazy[bi] += v
		}

		// 오른쪽 부분 블록: 개별 원소 갱신
		for i := rb * b; i <= r; i++ {
			a[i] += v
			blockSum[rb] += v
		}
	}

	// 구간 쿼리: [l, r]의 합을 반환한다
	rangeQuery := func(l, r int) int {
		sum := 0
		lb, rb := l/b, r/b

		if lb == rb {
			// 같은 블록 안에 있으면 개별 원소 + lazy를 합산한다
			for i := l; i <= r; i++ {
				sum += a[i] + lazy[lb]
			}
			return sum
		}

		// 왼쪽 부분 블록
		for i := l; i < (lb+1)*b; i++ {
			sum += a[i] + lazy[lb]
		}

		// 완전히 포함되는 중간 블록: 블록합 + lazy × 블록크기
		for bi := lb + 1; bi < rb; bi++ {
			sum += blockSum[bi] + lazy[bi]*blockSize[bi]
		}

		// 오른쪽 부분 블록
		for i := rb * b; i <= r; i++ {
			sum += a[i] + lazy[rb]
		}

		return sum
	}

	for ; q > 0; q-- {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			// 구간 갱신
			var l, r, v int
			fmt.Fscan(reader, &l, &r, &v)
			rangeUpdate(l-1, r-1, v)
		} else {
			// 구간 합 쿼리
			var l, r int
			fmt.Fscan(reader, &l, &r)
			fmt.Fprintln(writer, rangeQuery(l-1, r-1))
		}
	}
}
