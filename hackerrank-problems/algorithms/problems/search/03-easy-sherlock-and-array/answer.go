package main

import (
	"bufio"
	"fmt"
	"os"
)

// balancedSums는 왼쪽 합과 오른쪽 합이 같은 원소가 존재하는지 판별한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - string: "YES" 또는 "NO"
//
// [알고리즘 힌트]
//
//	전체 합을 구한 뒤, 왼쪽 합을 누적하면서 오른쪽 합과 비교한다.
//	오른쪽 합 = 전체 합 - 왼쪽 합 - 현재 원소.
func balancedSums(arr []int) string {
	// 전체 합 계산
	total := 0
	for _, v := range arr {
		total += v
	}

	// 왼쪽 합을 누적하면서 탐색
	leftSum := 0
	for _, v := range arr {
		// 오른쪽 합 = 전체 합 - 왼쪽 합 - 현재 원소
		rightSum := total - leftSum - v
		if leftSum == rightSum {
			return "YES"
		}
		// 왼쪽 합에 현재 원소 추가
		leftSum += v
	}

	return "NO"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(reader, &n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		fmt.Fprintln(writer, balancedSums(arr))
	}
}
