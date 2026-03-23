package main

import (
	"bufio"
	"fmt"
	"os"
)

// 슬라이딩 윈도우 최솟값/최댓값 - Sparse Table 풀이
// 정적 배열이므로 Sparse Table로 전처리 후 각 윈도우 위치의 쿼리를 O(1)에 응답한다
// 전처리: O(N log N), 전체 쿼리: O(N)

const MAXLOG = 21 // log2(1000000) ≈ 20

var sparseMin [MAXLOG][1000001]int
var sparseMax [MAXLOG][1000001]int
var logArr [1000002]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기와 윈도우 크기
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 로그 값 전처리
	for i := 2; i <= n; i++ {
		logArr[i] = logArr[i/2] + 1
	}

	// Sparse Table 구성: 길이 1인 구간 초기화
	for i := 0; i < n; i++ {
		sparseMin[0][i] = arr[i]
		sparseMax[0][i] = arr[i]
	}

	// 길이 2^j인 구간을 점화식으로 채운다
	maxJ := logArr[n] + 1
	for j := 1; j < maxJ; j++ {
		for i := 0; i+(1<<j)-1 < n; i++ {
			// 최솟값 테이블
			left := sparseMin[j-1][i]
			right := sparseMin[j-1][i+(1<<(j-1))]
			if left <= right {
				sparseMin[j][i] = left
			} else {
				sparseMin[j][i] = right
			}
			// 최댓값 테이블
			leftMax := sparseMax[j-1][i]
			rightMax := sparseMax[j-1][i+(1<<(j-1))]
			if leftMax >= rightMax {
				sparseMax[j][i] = leftMax
			} else {
				sparseMax[j][i] = rightMax
			}
		}
	}

	// 윈도우 크기 K에 대한 로그 값 (모든 쿼리에서 동일)
	kLog := logArr[k]

	// 최솟값 출력: 각 윈도우 위치에서 O(1) 쿼리
	for i := 0; i <= n-k; i++ {
		l := i
		r := i + k - 1
		leftVal := sparseMin[kLog][l]
		rightVal := sparseMin[kLog][r-(1<<kLog)+1]
		if leftVal <= rightVal {
			fmt.Fprint(writer, leftVal)
		} else {
			fmt.Fprint(writer, rightVal)
		}
		if i < n-k {
			fmt.Fprint(writer, " ")
		}
	}
	fmt.Fprintln(writer)

	// 최댓값 출력: 각 윈도우 위치에서 O(1) 쿼리
	for i := 0; i <= n-k; i++ {
		l := i
		r := i + k - 1
		leftVal := sparseMax[kLog][l]
		rightVal := sparseMax[kLog][r-(1<<kLog)+1]
		if leftVal >= rightVal {
			fmt.Fprint(writer, leftVal)
		} else {
			fmt.Fprint(writer, rightVal)
		}
		if i < n-k {
			fmt.Fprint(writer, " ")
		}
	}
	fmt.Fprintln(writer)
}
