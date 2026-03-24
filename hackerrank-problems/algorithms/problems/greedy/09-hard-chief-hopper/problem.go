package main

import (
	"bufio"
	"fmt"
	"os"
)

// chiefHopper는 모든 건물을 통과할 수 있는 최소 초기 에너지를 반환한다.
//
// [매개변수]
//   - arr: 건물 높이 배열
//
// [반환값]
//   - int: 최소 초기 에너지
func chiefHopper(arr []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 건물 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 건물 높이 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := chiefHopper(arr)
	fmt.Fprintln(writer, result)
}
