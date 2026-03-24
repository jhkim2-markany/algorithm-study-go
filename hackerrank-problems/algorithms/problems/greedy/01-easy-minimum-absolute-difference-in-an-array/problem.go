package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// minimumAbsoluteDifference는 배열에서 서로 다른 두 원소의 차이의 절댓값 중 최솟값을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - int: 최소 절대 차이
func minimumAbsoluteDifference(arr []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := minimumAbsoluteDifference(arr)
	fmt.Fprintln(writer, result)

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Ints
}
