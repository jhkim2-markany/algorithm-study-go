package main

import (
	"fmt"
	"math"
)

// Sparse Table (희소 배열)
// 정적 배열에서 구간 최솟값/최댓값 등 멱등 연산 쿼리를 O(1)에 응답한다.
// 전처리: O(N log N), 쿼리: O(1)

// SparseTable 구조체 - 구간 최솟값(RMQ)용
type SparseTable struct {
	table [][]int // table[k][i] = 인덱스 i에서 길이 2^k 구간의 최솟값
	log   []int   // log[i] = floor(log2(i))
}

// NewSparseTable은 배열 arr로부터 Sparse Table을 구성한다
// 시간 복잡도: O(N log N)
// 공간 복잡도: O(N log N)
func NewSparseTable(arr []int) *SparseTable {
	n := len(arr)
	if n == 0 {
		return &SparseTable{}
	}

	// 로그 값 전처리
	logTable := make([]int, n+1)
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	maxK := logTable[n] + 1
	// 테이블 초기화
	table := make([][]int, maxK)
	for k := 0; k < maxK; k++ {
		table[k] = make([]int, n)
	}

	// 길이 1인 구간: 원래 배열 값을 그대로 저장한다
	for i := 0; i < n; i++ {
		table[0][i] = arr[i]
	}

	// 길이 2^k인 구간을 점화식으로 채운다
	for k := 1; k < maxK; k++ {
		for i := 0; i+(1<<k)-1 < n; i++ {
			// 두 구간의 최솟값을 합친다
			left := table[k-1][i]
			right := table[k-1][i+(1<<(k-1))]
			if left <= right {
				table[k][i] = left
			} else {
				table[k][i] = right
			}
		}
	}

	return &SparseTable{table: table, log: logTable}
}

// Query는 구간 [l, r]의 최솟값을 O(1)에 반환한다
func (st *SparseTable) Query(l, r int) int {
	// 구간 길이에 해당하는 로그 값을 구한다
	length := r - l + 1
	k := st.log[length]
	// 겹치는 두 구간의 최솟값을 합친다 (멱등 연산이므로 겹쳐도 정확하다)
	left := st.table[k][l]
	right := st.table[k][r-(1<<k)+1]
	if left <= right {
		return left
	}
	return right
}

// SparseTableMax 구조체 - 구간 최댓값용
type SparseTableMax struct {
	table [][]int
	log   []int
}

// NewSparseTableMax는 배열 arr로부터 구간 최댓값 Sparse Table을 구성한다
func NewSparseTableMax(arr []int) *SparseTableMax {
	n := len(arr)
	if n == 0 {
		return &SparseTableMax{}
	}

	logTable := make([]int, n+1)
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	maxK := logTable[n] + 1
	table := make([][]int, maxK)
	for k := 0; k < maxK; k++ {
		table[k] = make([]int, n)
	}

	// 길이 1인 구간 초기화
	for i := 0; i < n; i++ {
		table[0][i] = arr[i]
	}

	// 길이 2^k인 구간: 두 구간의 최댓값을 합친다
	for k := 1; k < maxK; k++ {
		for i := 0; i+(1<<k)-1 < n; i++ {
			left := table[k-1][i]
			right := table[k-1][i+(1<<(k-1))]
			if left >= right {
				table[k][i] = left
			} else {
				table[k][i] = right
			}
		}
	}

	return &SparseTableMax{table: table, log: logTable}
}

// Query는 구간 [l, r]의 최댓값을 O(1)에 반환한다
func (st *SparseTableMax) Query(l, r int) int {
	length := r - l + 1
	k := st.log[length]
	left := st.table[k][l]
	right := st.table[k][r-(1<<k)+1]
	if left >= right {
		return left
	}
	return right
}

// GCD를 구하는 함수
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// SparseTableGCD 구조체 - 구간 GCD용
type SparseTableGCD struct {
	table [][]int
	log   []int
}

// NewSparseTableGCD는 배열 arr로부터 구간 GCD Sparse Table을 구성한다
func NewSparseTableGCD(arr []int) *SparseTableGCD {
	n := len(arr)
	if n == 0 {
		return &SparseTableGCD{}
	}

	logTable := make([]int, n+1)
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	maxK := logTable[n] + 1
	table := make([][]int, maxK)
	for k := 0; k < maxK; k++ {
		table[k] = make([]int, n)
	}

	// 길이 1인 구간: 원래 값 저장
	for i := 0; i < n; i++ {
		table[0][i] = arr[i]
	}

	// 길이 2^k인 구간: 두 구간의 GCD를 합친다
	for k := 1; k < maxK; k++ {
		for i := 0; i+(1<<k)-1 < n; i++ {
			table[k][i] = gcd(table[k-1][i], table[k-1][i+(1<<(k-1))])
		}
	}

	return &SparseTableGCD{table: table, log: logTable}
}

// Query는 구간 [l, r]의 GCD를 O(1)에 반환한다
func (st *SparseTableGCD) Query(l, r int) int {
	length := r - l + 1
	k := st.log[length]
	return gcd(st.table[k][l], st.table[k][r-(1<<k)+1])
}

func main() {
	// 예제 배열
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	fmt.Println("배열:", arr)

	// 1. 구간 최솟값 (RMQ)
	fmt.Println("\n=== 구간 최솟값 (Sparse Table) ===")
	stMin := NewSparseTable(arr)
	queries := [][2]int{{0, 3}, {2, 7}, {0, 9}, {4, 6}}
	for _, q := range queries {
		l, r := q[0], q[1]
		fmt.Printf("min(%d, %d) = %d\n", l, r, stMin.Query(l, r))
	}

	// 2. 구간 최댓값
	fmt.Println("\n=== 구간 최댓값 (Sparse Table) ===")
	stMax := NewSparseTableMax(arr)
	for _, q := range queries {
		l, r := q[0], q[1]
		fmt.Printf("max(%d, %d) = %d\n", l, r, stMax.Query(l, r))
	}

	// 3. 구간 GCD
	fmt.Println("\n=== 구간 GCD (Sparse Table) ===")
	arrGCD := []int{12, 18, 24, 36, 48, 60}
	fmt.Println("배열:", arrGCD)
	stGCD := NewSparseTableGCD(arrGCD)
	gcdQueries := [][2]int{{0, 2}, {1, 4}, {0, 5}, {3, 5}}
	for _, q := range gcdQueries {
		l, r := q[0], q[1]
		fmt.Printf("gcd(%d, %d) = %d\n", l, r, stGCD.Query(l, r))
	}

	// 4. 세그먼트 트리와의 비교
	fmt.Println("\n=== Sparse Table vs Segment Tree 비교 ===")
	fmt.Println("Sparse Table: 전처리 O(N log N), 쿼리 O(1), 갱신 불가")
	fmt.Println("Segment Tree: 전처리 O(N), 쿼리 O(log N), 갱신 O(log N)")
	fmt.Printf("배열 크기 N = %d일 때 log₂(N) = %.1f\n", len(arr), math.Log2(float64(len(arr))))
}
