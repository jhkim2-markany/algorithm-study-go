package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// 쿼리 구조체
type Query struct {
	left  int
	right int
	index int
}

// BIT (Binary Indexed Tree): 구간 내 역전 쌍 계산에 사용한다
var bit []int
var bitSize int

// BIT 업데이트: idx 위치에 val을 더한다
func update(idx, val int) {
	for ; idx <= bitSize; idx += idx & (-idx) {
		bit[idx] += val
	}
}

// BIT 쿼리: 1부터 idx까지의 합을 구한다
func query(idx int) int {
	sum := 0
	for ; idx > 0; idx -= idx & (-idx) {
		sum += bit[idx]
	}
	return sum
}

var (
	blockSize  int
	inversions int // 현재 구간의 역전 쌍 개수
)

// 구간 오른쪽 끝에 원소를 추가한다
// 새 원소보다 큰 값의 개수가 역전 쌍 증가분이다
func addRight(val int) {
	// val보다 큰 값의 개수 = 전체 개수 - val 이하의 개수
	total := query(bitSize)
	lessOrEqual := query(val)
	inversions += total - lessOrEqual
	update(val, 1)
}

// 구간 오른쪽 끝에서 원소를 제거한다
func removeRight(val int) {
	update(val, -1)
	total := query(bitSize)
	lessOrEqual := query(val)
	inversions -= total - lessOrEqual
}

// 구간 왼쪽 끝에 원소를 추가한다
// 새 원소보다 작은 값의 개수가 역전 쌍 증가분이다
func addLeft(val int) {
	// val보다 작은 값의 개수
	less := query(val - 1)
	inversions += less
	update(val, 1)
}

// 구간 왼쪽 끝에서 원소를 제거한다
func removeLeft(val int) {
	update(val, -1)
	less := query(val - 1)
	inversions -= less
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 쿼리 수 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 배열 입력 (1-indexed)
	arr := make([]int, n+1)
	maxVal := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}

	// 쿼리 입력
	queries := make([]Query, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i].left, &queries[i].right)
		queries[i].index = i
	}

	// 블록 크기 설정
	blockSize = int(math.Sqrt(float64(n)))
	if blockSize == 0 {
		blockSize = 1
	}

	// Mo's 알고리즘 정렬
	sort.Slice(queries, func(i, j int) bool {
		bi := queries[i].left / blockSize
		bj := queries[j].left / blockSize
		if bi != bj {
			return bi < bj
		}
		// 홀짝 최적화: 홀수 블록은 R 내림차순
		if bi%2 == 1 {
			return queries[i].right > queries[j].right
		}
		return queries[i].right < queries[j].right
	})

	// BIT 초기화
	bitSize = maxVal
	bit = make([]int, bitSize+1)
	inversions = 0

	// 결과 배열
	answers := make([]int, q)

	// 현재 구간 초기화
	curL, curR := 1, 0

	// Mo's 알고리즘으로 쿼리 처리
	for _, qr := range queries {
		l, r := qr.left, qr.right

		// 오른쪽으로 확장
		for curR < r {
			curR++
			addRight(arr[curR])
		}
		// 왼쪽으로 확장
		for curL > l {
			curL--
			addLeft(arr[curL])
		}
		// 오른쪽에서 축소
		for curR > r {
			removeRight(arr[curR])
			curR--
		}
		// 왼쪽에서 축소
		for curL < l {
			removeLeft(arr[curL])
			curL++
		}

		// 현재 구간의 역전 쌍 개수를 기록한다
		answers[qr.index] = inversions
	}

	// 원래 순서대로 결과 출력
	for _, ans := range answers {
		fmt.Fprintln(writer, ans)
	}
}
