package main

import (
	"fmt"
	"math"
)

// 제곱근 분할법 (Sqrt Decomposition) - 구간 합 쿼리 + 점 갱신
// 배열을 √N 크기의 블록으로 분할하여 구간 합을 O(√N)에 처리한다.
// 시간 복잡도: 점 갱신 O(1), 구간 쿼리 O(√N)
// 공간 복잡도: O(N)

// SqrtDecomp는 제곱근 분할법 자료구조이다
type SqrtDecomp struct {
	a     []int // 원본 배열
	block []int // 블록별 합
	b     int   // 블록 크기
	n     int   // 배열 크기
}

// NewSqrtDecomp는 배열로부터 제곱근 분할법 자료구조를 생성한다
func NewSqrtDecomp(a []int) *SqrtDecomp {
	n := len(a)
	b := int(math.Ceil(math.Sqrt(float64(n))))
	if b == 0 {
		b = 1
	}
	// 블록 개수 계산
	numBlocks := (n + b - 1) / b
	block := make([]int, numBlocks)

	// 각 블록의 합을 미리 계산한다
	for i := 0; i < n; i++ {
		block[i/b] += a[i]
	}

	// 원본 배열 복사
	arr := make([]int, n)
	copy(arr, a)

	return &SqrtDecomp{a: arr, block: block, b: b, n: n}
}

// Update는 인덱스 idx의 값을 val로 변경한다
func (sd *SqrtDecomp) Update(idx, val int) {
	// 블록 합에서 기존 값을 빼고 새 값을 더한다
	sd.block[idx/sd.b] += val - sd.a[idx]
	sd.a[idx] = val
}

// Query는 구간 [l, r]의 합을 반환한다
func (sd *SqrtDecomp) Query(l, r int) int {
	sum := 0
	// l부터 r까지 순회하되, 완전한 블록은 블록 합을 사용한다
	for i := l; i <= r; {
		// 현재 위치가 블록의 시작이고, 블록 끝까지 r에 포함되면 블록 합 사용
		if i%sd.b == 0 && i+sd.b-1 <= r {
			sum += sd.block[i/sd.b]
			i += sd.b
		} else {
			// 부분 블록은 개별 원소를 더한다
			sum += sd.a[i]
			i++
		}
	}
	return sum
}

func main() {
	// 예시 배열
	a := []int{1, 3, 5, 2, 7, 6, 3, 8}
	sd := NewSqrtDecomp(a)

	fmt.Println("배열:", a)
	fmt.Printf("블록 크기: %d\n", sd.b)
	fmt.Printf("블록 합: %v\n\n", sd.block)

	// 구간 합 쿼리
	fmt.Printf("구간 합 [1, 6] = %d\n", sd.Query(1, 6))
	// 3 + 5 + 2 + 7 + 6 + 3 = 26

	fmt.Printf("구간 합 [0, 7] = %d\n", sd.Query(0, 7))
	// 전체 합 = 35

	fmt.Printf("구간 합 [3, 5] = %d\n\n", sd.Query(3, 5))
	// 2 + 7 + 6 = 15

	// 점 갱신: a[4] = 10 (기존 7)
	sd.Update(4, 10)
	fmt.Println("a[4] = 10으로 갱신 후:")
	fmt.Printf("구간 합 [1, 6] = %d\n", sd.Query(1, 6))
	// 3 + 5 + 2 + 10 + 6 + 3 = 29

	fmt.Printf("구간 합 [0, 7] = %d\n", sd.Query(0, 7))
	// 전체 합 = 38
}
