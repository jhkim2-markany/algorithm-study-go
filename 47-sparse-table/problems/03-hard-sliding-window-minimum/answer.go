package main

import (
	"bufio"
	"fmt"
	"os"
)

// slidingWindowMinMax는 Sparse Table을 이용하여 크기 K인 슬라이딩 윈도우의
// 최솟값 배열과 최댓값 배열을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - k: 윈도우 크기
//
// [반환값]
//   - []int: 각 윈도우 위치의 최솟값 배열
//   - []int: 각 윈도우 위치의 최댓값 배열
//
// [알고리즘 힌트]
//
//	Sparse Table을 최솟값/최댓값 각각 구축한 뒤,
//	각 윈도우 위치에서 O(1) 쿼리로 결과를 구한다.
//	전처리 O(N log N), 전체 쿼리 O(N).
func slidingWindowMinMax(arr []int, k int) ([]int, []int) {
	n := len(arr)

	logArr := make([]int, n+2)
	for i := 2; i <= n; i++ {
		logArr[i] = logArr[i/2] + 1
	}

	maxJ := logArr[n] + 1
	sparseMin := make([][]int, maxJ)
	sparseMax := make([][]int, maxJ)
	for j := 0; j < maxJ; j++ {
		sparseMin[j] = make([]int, n)
		sparseMax[j] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		sparseMin[0][i] = arr[i]
		sparseMax[0][i] = arr[i]
	}
	for j := 1; j < maxJ; j++ {
		for i := 0; i+(1<<j)-1 < n; i++ {
			left := sparseMin[j-1][i]
			right := sparseMin[j-1][i+(1<<(j-1))]
			if left <= right {
				sparseMin[j][i] = left
			} else {
				sparseMin[j][i] = right
			}
			leftMax := sparseMax[j-1][i]
			rightMax := sparseMax[j-1][i+(1<<(j-1))]
			if leftMax >= rightMax {
				sparseMax[j][i] = leftMax
			} else {
				sparseMax[j][i] = rightMax
			}
		}
	}

	kLog := logArr[k]

	mins := make([]int, n-k+1)
	maxs := make([]int, n-k+1)
	for i := 0; i <= n-k; i++ {
		l := i
		r := i + k - 1
		leftVal := sparseMin[kLog][l]
		rightVal := sparseMin[kLog][r-(1<<kLog)+1]
		if leftVal <= rightVal {
			mins[i] = leftVal
		} else {
			mins[i] = rightVal
		}
		leftMax := sparseMax[kLog][l]
		rightMax := sparseMax[kLog][r-(1<<kLog)+1]
		if leftMax >= rightMax {
			maxs[i] = leftMax
		} else {
			maxs[i] = rightMax
		}
	}
	return mins, maxs
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	mins, maxs := slidingWindowMinMax(arr, k)

	for i, v := range mins {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	for i, v := range maxs {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
