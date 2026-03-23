package main

import (
	"bufio"
	"fmt"
	"os"
)

// findSumPair는 배열에서 합이 target인 두 수를 찾아 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - target: 두 수의 합이 되어야 하는 목표값
//
// [반환값]
//   - int, int: 합이 target인 두 수 (첫 번째로 발견된 쌍)
//   - bool: 조건을 만족하는 쌍을 찾았으면 true, 없으면 false
func findSumPair(arr []int, target int) (int, int, bool) {
	// 여기에 코드를 작성하세요
	return 0, 0, false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 목표 합 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	a, b, found := findSumPair(arr, m)

	// 결과 출력
	if found {
		fmt.Fprintf(writer, "%d %d\n", a, b)
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
