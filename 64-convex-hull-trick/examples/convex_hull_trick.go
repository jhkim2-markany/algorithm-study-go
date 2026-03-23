package main

import "fmt"

// 볼록 껍질 트릭 (Convex Hull Trick) - 최솟값 쿼리
// 일차 함수 f(x) = mx + b 를 추가하고, 특정 x에서의 최솟값을 구한다.
// 기울기가 단조 감소하는 순서로 추가되고, 쿼리 x가 단조 증가할 때 O(N)에 동작한다.
// 시간 복잡도: 삽입 O(N), 쿼리 O(N) (전체 amortized)
// 공간 복잡도: O(N)

// Line은 일차 함수 y = m*x + b를 나타낸다
type Line struct {
	m, b int64 // 기울기, 절편
}

// eval은 직선에 x를 대입한 값을 반환한다
func (l Line) eval(x int64) int64 {
	return l.m*x + l.b
}

// CHT는 볼록 껍질 트릭 자료구조이다
type CHT struct {
	lines []Line // 볼록 껍질을 이루는 직선들
	ptr   int    // 단조 쿼리용 포인터
}

// bad는 직선 b가 직선 a와 c 사이에서 불필요한지 판정한다
// a, b, c 순서로 기울기가 감소할 때, b가 최솟값에 기여하지 않으면 true
func bad(a, b, c Line) bool {
	// 교점 비교를 정수 연산으로 수행 (나눗셈 회피)
	// a와 c의 교점 x좌표에서 b의 값이 a(또는 c)의 값 이상이면 b는 불필요
	return (c.b-a.b)*(a.m-b.m) <= (b.b-a.b)*(a.m-c.m)
}

// AddLine은 직선 y = mx + b를 추가한다
// 기울기 m이 단조 감소하는 순서로 추가해야 한다
func (cht *CHT) AddLine(m, b int64) {
	newLine := Line{m, b}
	// 마지막 두 직선과 새 직선을 비교하여 불필요한 직선 제거
	for len(cht.lines) >= 2 {
		n := len(cht.lines)
		if bad(cht.lines[n-2], cht.lines[n-1], newLine) {
			cht.lines = cht.lines[:n-1]
		} else {
			break
		}
	}
	cht.lines = append(cht.lines, newLine)
}

// QueryMonotone은 단조 증가하는 x에 대해 최솟값을 반환한다
// x가 이전 쿼리보다 크거나 같아야 한다 (포인터 전진만 함)
func (cht *CHT) QueryMonotone(x int64) int64 {
	for cht.ptr+1 < len(cht.lines) &&
		cht.lines[cht.ptr+1].eval(x) <= cht.lines[cht.ptr].eval(x) {
		cht.ptr++
	}
	return cht.lines[cht.ptr].eval(x)
}

// QueryBinary는 임의의 x에 대해 이분 탐색으로 최솟값을 반환한다
func (cht *CHT) QueryBinary(x int64) int64 {
	lo, hi := 0, len(cht.lines)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if cht.lines[mid].eval(x) <= cht.lines[mid+1].eval(x) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return cht.lines[lo].eval(x)
}

func main() {
	// 예시: 직선 3개를 추가하고 여러 x에서 최솟값을 쿼리한다
	//
	// 직선들 (기울기 감소 순서로 추가):
	//   f₁(x) = 3x + 1
	//   f₂(x) = 1x + 3
	//   f₃(x) = -1x + 8
	//
	// 각 x에서의 최솟값:
	//   x=0: min(1, 3, 8) = 1
	//   x=1: min(4, 4, 7) = 4
	//   x=2: min(7, 5, 6) = 5
	//   x=3: min(10, 6, 5) = 5
	//   x=5: min(16, 8, 3) = 3

	cht := &CHT{}

	// 기울기 감소 순서로 직선 추가
	cht.AddLine(3, 1)  // y = 3x + 1
	cht.AddLine(1, 3)  // y = x + 3
	cht.AddLine(-1, 8) // y = -x + 8

	fmt.Println("=== 볼록 껍질 트릭 (Convex Hull Trick) 예시 ===")
	fmt.Println()

	// 볼록 껍질에 남은 직선 출력
	fmt.Printf("볼록 껍질 직선 수: %d\n", len(cht.lines))
	for i, l := range cht.lines {
		fmt.Printf("  직선 %d: y = %dx + %d\n", i, l.m, l.b)
	}
	fmt.Println()

	// 단조 증가 쿼리 (포인터 방식)
	fmt.Println("[단조 쿼리]")
	queries := []int64{0, 1, 2, 3, 5}
	for _, x := range queries {
		val := cht.QueryMonotone(x)
		fmt.Printf("  x=%d → 최솟값 = %d\n", x, val)
	}
	fmt.Println()

	// 이분 탐색 쿼리 (임의 순서 가능)
	cht2 := &CHT{}
	cht2.AddLine(3, 1)
	cht2.AddLine(1, 3)
	cht2.AddLine(-1, 8)

	fmt.Println("[이분 탐색 쿼리]")
	arbitraryQueries := []int64{5, 0, 3, 2, 1}
	for _, x := range arbitraryQueries {
		val := cht2.QueryBinary(x)
		fmt.Printf("  x=%d → 최솟값 = %d\n", x, val)
	}

	// 출력:
	// === 볼록 껍질 트릭 (Convex Hull Trick) 예시 ===
	//
	// 볼록 껍질 직선 수: 3
	//   직선 0: y = 3x + 1
	//   직선 1: y = 1x + 3
	//   직선 2: y = -1x + 8
	//
	// [단조 쿼리]
	//   x=0 → 최솟값 = 1
	//   x=1 → 최솟값 = 4
	//   x=2 → 최솟값 = 5
	//   x=3 → 최솟값 = 5
	//   x=5 → 최솟값 = 3
	//
	// [이분 탐색 쿼리]
	//   x=5 → 최솟값 = 3
	//   x=0 → 최솟값 = 1
	//   x=3 → 최솟값 = 5
	//   x=2 → 최솟값 = 5
	//   x=1 → 최솟값 = 4
}
