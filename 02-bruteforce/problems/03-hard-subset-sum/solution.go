package main

import (
	"bufio"
	"fmt"
	"os"
)

// countSubsetSum은 배열의 부분집합 중 합이 target인
// 부분집합의 개수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열 (원소는 양수, 음수, 0 모두 가능)
//   - target: 부분집합의 합이 되어야 하는 목표값
//
// [반환값]
//   - int: 합이 target인 부분집합의 개수 (빈 집합 제외)
func countSubsetSum(arr []int, target int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 목표 합 입력
	var n, s int
	fmt.Fscan(reader, &n, &s)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	count := countSubsetSum(arr, s)

	// 결과 출력
	fmt.Fprintln(writer, count)
}
