package main

import (
	"bufio"
	"fmt"
	"os"
)

// twoSumSorted는 정렬된 배열에서 합이 target인 두 원소의 인덱스를 반환한다.
//
// [매개변수]
//   - arr: 오름차순 정렬된 정수 배열
//   - target: 두 수의 합이 되어야 하는 목표값
//
// [반환값]
//   - int, int: 합이 target인 두 원소의 인덱스 (1-indexed)
//   - bool: 조건을 만족하는 쌍을 찾았으면 true, 없으면 false
func twoSumSorted(arr []int, target int) (int, int, bool) {
	// 여기에 코드를 작성하세요
	return 0, 0, false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 목표값 입력
	var n, t int
	fmt.Fscan(reader, &n, &t)

	// 정렬된 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	l, r, found := twoSumSorted(arr, t)
	if found {
		fmt.Fprintf(writer, "%d %d\n", l, r)
	}
}
