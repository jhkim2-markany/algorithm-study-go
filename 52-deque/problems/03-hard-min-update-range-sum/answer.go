package main

import (
	"bufio"
	"fmt"
	"os"
)

// minUpdateRangeSum은 배열에서 각 원소가 포함되는 크기 K 윈도우의 최솟값 중
// 최댓값을 B[i]로 정의할 때, B[i]의 총합을 반환한다.
//
// [매개변수]
//   - a: 정수 배열
//   - k: 윈도우 크기
//
// [반환값]
//   - int64: B[i]의 총합
//
// [알고리즘 힌트]
//
//	1단계: 모노톤 덱으로 각 윈도우의 최솟값(windowMin)을 구한다.
//	2단계: windowMin 배열에 대해 Sparse Table을 구축하여 구간 최댓값 쿼리를 O(1)에 처리한다.
//	3단계: 각 인덱스 i가 포함되는 윈도우 범위에서 windowMin의 최댓값이 B[i]이다.
func minUpdateRangeSum(a []int, k int) int64 {
	n := len(a)
	numWindows := n - k + 1
	windowMin := make([]int, numWindows)

	// 1단계: 모노톤 덱으로 윈도우 최솟값
	deque := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && a[deque[len(deque)-1]] >= a[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			windowMin[i-k+1] = a[deque[0]]
		}
	}

	// 2단계: Sparse Table 구축 (구간 최댓값)
	logN := 0
	for (1 << logN) <= numWindows {
		logN++
	}
	sparse := make([][]int, logN)
	sparse[0] = make([]int, numWindows)
	copy(sparse[0], windowMin)
	for p := 1; p < logN; p++ {
		length := 1 << p
		sparse[p] = make([]int, numWindows)
		for j := 0; j+length-1 < numWindows; j++ {
			if sparse[p-1][j] > sparse[p-1][j+length/2] {
				sparse[p][j] = sparse[p-1][j]
			} else {
				sparse[p][j] = sparse[p-1][j+length/2]
			}
		}
	}

	queryMax := func(l, r int) int {
		if l > r {
			return 0
		}
		length := r - l + 1
		p := 0
		for (1 << (p + 1)) <= length {
			p++
		}
		a := sparse[p][l]
		b := sparse[p][r-(1<<p)+1]
		if a > b {
			return a
		}
		return b
	}

	// 3단계: B[i]의 합 계산
	sum := int64(0)
	for i := 0; i < n; i++ {
		lo := i - k + 1
		if lo < 0 {
			lo = 0
		}
		hi := i
		if hi > numWindows-1 {
			hi = numWindows - 1
		}
		sum += int64(queryMax(lo, hi))
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	fmt.Fprintln(writer, minUpdateRangeSum(a, k))
}
