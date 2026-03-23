package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	arr []int
)

// maxSubarray 함수는 분할 정복으로 최대 부분 배열 합을 구한다
func maxSubarray(lo, hi int) int {
	// 기저 조건: 원소가 하나이면 그 값을 반환
	if lo == hi {
		return arr[lo]
	}

	mid := (lo + hi) / 2

	// 왼쪽 부분의 최대 부분 배열 합
	leftMax := maxSubarray(lo, mid)
	// 오른쪽 부분의 최대 부분 배열 합
	rightMax := maxSubarray(mid+1, hi)
	// 중간을 걸치는 최대 부분 배열 합
	crossMax := maxCrossing(lo, mid, hi)

	// 세 값 중 최댓값 반환
	return max3(leftMax, rightMax, crossMax)
}

// maxCrossing 함수는 중간 지점을 걸치는 최대 부분 배열 합을 구한다
func maxCrossing(lo, mid, hi int) int {
	// 중간에서 왼쪽으로 확장
	leftSum := arr[mid]
	sum := arr[mid]
	for i := mid - 1; i >= lo; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// 중간에서 오른쪽으로 확장
	rightSum := arr[mid+1]
	sum = arr[mid+1]
	for i := mid + 2; i <= hi; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

// max3 함수는 세 정수 중 최댓값을 반환한다
func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= c {
		return b
	}
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	fmt.Fscan(reader, &n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 분할 정복으로 최대 부분 배열 합 계산
	result := maxSubarray(0, n-1)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
