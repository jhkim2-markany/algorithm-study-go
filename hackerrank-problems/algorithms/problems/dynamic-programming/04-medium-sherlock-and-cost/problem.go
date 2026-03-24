package main

import (
	"bufio"
	"fmt"
	"os"
)

// sherlockAndCost는 인접 원소 간 절댓값 차이의 합의 최댓값을 반환한다.
//
// [매개변수]
//   - b: 각 위치의 상한값 배열
//
// [반환값]
//   - int: S의 최댓값
func sherlockAndCost(b []int) int {
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
		// 배열 크기 입력
		var n int
		fmt.Fscan(reader, &n)

		// 배열 B 입력
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := sherlockAndCost(b)
		fmt.Fprintln(writer, result)
	}
}
