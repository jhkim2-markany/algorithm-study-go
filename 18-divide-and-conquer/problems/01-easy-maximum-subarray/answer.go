package main

import (
	"bufio"
	"fmt"
	"os"
)

// maxSubarraySum은 분할 정복으로 배열의 최대 부분 배열 합을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - int: 최대 부분 배열 합
//
// [알고리즘 힌트]
//
//	배열을 반으로 나누어 세 가지 경우를 비교한다:
//	1) 왼쪽 부분의 최대 부분 배열 합
//	2) 오른쪽 부분의 최대 부분 배열 합
//	3) 중간을 걸치는 최대 부분 배열 합
//	중간을 걸치는 경우는 mid에서 왼쪽/오른쪽으로 확장하며 최대 합을 구한다.
//	기저 조건: 원소가 하나이면 그 값을 반환한다.
func maxSubarraySum(arr []int) int {
	var solve func(lo, hi int) int
	solve = func(lo, hi int) int {
		if lo == hi {
			return arr[lo]
		}

		mid := (lo + hi) / 2
		leftMax := solve(lo, mid)
		rightMax := solve(mid+1, hi)

		// 중간을 걸치는 최대 합
		leftSum := arr[mid]
		sum := arr[mid]
		for i := mid - 1; i >= lo; i-- {
			sum += arr[i]
			if sum > leftSum {
				leftSum = sum
			}
		}

		rightSum := arr[mid+1]
		sum = arr[mid+1]
		for i := mid + 2; i <= hi; i++ {
			sum += arr[i]
			if sum > rightSum {
				rightSum = sum
			}
		}

		crossMax := leftSum + rightSum

		// 세 값 중 최댓값
		best := leftMax
		if rightMax > best {
			best = rightMax
		}
		if crossMax > best {
			best = crossMax
		}
		return best
	}

	return solve(0, len(arr)-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := maxSubarraySum(arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
