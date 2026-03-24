package main

import (
	"bufio"
	"fmt"
	"os"
)

// unboundedKnapsack은 원소를 중복 사용하여 K 이하의 최대 합을 반환한다.
//
// [매개변수]
//   - k: 목표값 (상한)
//   - arr: 정수 배열
//
// [반환값]
//   - int: K를 넘지 않는 최대 합
func unboundedKnapsack(k int, arr []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		// 배열 크기와 목표값 입력
		var n, k int
		fmt.Fscan(reader, &n, &k)

		// 배열 입력
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := unboundedKnapsack(k, arr)
		fmt.Fprintln(writer, result)
	}
}
