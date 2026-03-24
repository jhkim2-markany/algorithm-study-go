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
//
// [알고리즘 힌트]
//
//	보조 스택(max stack)을 사용하여 현재까지의 최대값을 추적한다.
//	푸시 시 현재 값과 보조 스택 최상위 중 큰 값을 보조 스택에 넣는다.
//	팝 시 두 스택 모두에서 제거한다.
func getMaximumElement(queries [][]int) []int {
	// 메인 스택과 최대값 추적용 보조 스택 초기화
	stack := []int{}
	maxStack := []int{}
	result := []int{}

	for _, q := range queries {
		switch q[0] {
		case 1:
			// 원소를 메인 스택에 푸시
			x := q[1]
			stack = append(stack, x)

			// 보조 스택이 비어있거나 새 원소가 현재 최대값 이상이면 갱신
			if len(maxStack) == 0 || x >= maxStack[len(maxStack)-1] {
				maxStack = append(maxStack, x)
			} else {
				maxStack = append(maxStack, maxStack[len(maxStack)-1])
			}
		case 2:
			// 두 스택 모두에서 팝
			stack = stack[:len(stack)-1]
			maxStack = maxStack[:len(maxStack)-1]
		case 3:
			// 보조 스택의 최상위가 현재 최대값
			result = append(result, maxStack[len(maxStack)-1])
		}
	}

	return result
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
