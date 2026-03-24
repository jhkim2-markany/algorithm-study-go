package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// sherlockAndMinimax는 범위 [p, q]에서 min(|arr[i] - M|)을 최대화하는 M을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - p: 범위 시작
//   - q: 범위 끝
//
// [반환값]
//   - int: 조건을 만족하는 M
func sherlockAndMinimax(arr []int, p int, q int) int {
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

	// 범위 입력
	var p, q int
	fmt.Fscan(reader, &p, &q)

	// 핵심 함수 호출 및 결과 출력
	result := sherlockAndMinimax(arr, p, q)
	fmt.Fprintln(writer, result)

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Ints
}
