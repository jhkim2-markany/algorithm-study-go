package main

import (
	"bufio"
	"fmt"
	"os"
)

// getMaximumElement는 스택 쿼리를 처리하여 쿼리 타입 3에 대한 최대값 목록을 반환한다.
//
// [매개변수]
//   - queries: 쿼리 목록 (각 쿼리는 [타입] 또는 [타입, 값] 형태)
//
// [반환값]
//   - []int: 쿼리 타입 3에 대한 최대값 목록
func getMaximumElement(queries [][]int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 쿼리 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 쿼리 입력
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		var qType int
		fmt.Fscan(reader, &qType)
		if qType == 1 {
			var x int
			fmt.Fscan(reader, &x)
			queries[i] = []int{qType, x}
		} else {
			queries[i] = []int{qType}
		}
	}

	// 핵심 함수 호출
	result := getMaximumElement(queries)

	// 결과 출력
	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}
