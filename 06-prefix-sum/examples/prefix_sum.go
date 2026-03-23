package main

import "fmt"

// 누적합(Prefix Sum) 활용 예시 - 1차원 및 2차원 누적합
// 시간 복잡도: 전처리 O(N), 쿼리 O(1)
// 공간 복잡도: O(N)

// buildPrefixSum1D 함수는 1차원 누적합 배열을 생성한다.
// P[i] = A[0] + A[1] + ... + A[i-1]
func buildPrefixSum1D(arr []int) []int {
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}
	return prefix
}

// rangeSum1D 함수는 구간 [l, r]의 합을 O(1)에 반환한다. (0-indexed)
func rangeSum1D(prefix []int, l, r int) int {
	return prefix[r+1] - prefix[l]
}

// buildPrefixSum2D 함수는 2차원 누적합 배열을 생성한다.
// 포함-배제 원리를 이용하여 (0,0)부터 (i-1,j-1)까지의 합을 저장한다.
func buildPrefixSum2D(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	n := len(matrix)
	m := len(matrix[0])

	// (n+1) x (m+1) 크기의 누적합 배열 생성
	prefix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prefix[i] = make([]int, m+1)
	}

	// 포함-배제 원리로 누적합 계산
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return prefix
}

// rangeSum2D 함수는 (r1,c1)부터 (r2,c2)까지의 부분 행렬 합을 O(1)에 반환한다. (0-indexed)
func rangeSum2D(prefix [][]int, r1, c1, r2, c2 int) int {
	return prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
}

func main() {
	// 예시 1: 1차원 누적합
	fmt.Println("=== 1차원 누적합 ===")
	arr := []int{3, 1, 4, 1, 5, 9}
	prefix := buildPrefixSum1D(arr)
	fmt.Printf("원본 배열: %v\n", arr)
	fmt.Printf("누적합 배열: %v\n", prefix)

	// 구간 합 쿼리
	fmt.Printf("구간 [1, 3]의 합: %d\n", rangeSum1D(prefix, 1, 3)) // 1+4+1 = 6
	fmt.Printf("구간 [0, 5]의 합: %d\n", rangeSum1D(prefix, 0, 5)) // 3+1+4+1+5+9 = 23
	fmt.Printf("구간 [2, 4]의 합: %d\n", rangeSum1D(prefix, 2, 4)) // 4+1+5 = 10

	// 예시 2: 2차원 누적합
	fmt.Println("\n=== 2차원 누적합 ===")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("원본 행렬:")
	for _, row := range matrix {
		fmt.Printf("  %v\n", row)
	}

	prefix2D := buildPrefixSum2D(matrix)

	// 부분 행렬 합 쿼리
	fmt.Printf("(0,0)~(1,1) 합: %d\n", rangeSum2D(prefix2D, 0, 0, 1, 1)) // 1+2+4+5 = 12
	fmt.Printf("(1,1)~(2,2) 합: %d\n", rangeSum2D(prefix2D, 1, 1, 2, 2)) // 5+6+8+9 = 28
	fmt.Printf("(0,0)~(2,2) 합: %d\n", rangeSum2D(prefix2D, 0, 0, 2, 2)) // 전체 합 = 45
}
