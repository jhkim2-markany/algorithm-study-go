package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// heapSort는 힙 정렬을 수행하며 각 단계의 배열 상태를 반환한다.
//
// [매개변수]
//   - arr: 정렬할 정수 배열
//
// [반환값]
//   - [][]int: 각 교환 후 배열 상태
func heapSort(arr []int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
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

	// 핵심 함수 호출
	steps := heapSort(arr)

	// 결과 출력
	for _, step := range steps {
		strs := make([]string, len(step))
		for i, v := range step {
			strs[i] = fmt.Sprintf("%d", v)
		}
		fmt.Fprintln(writer, strings.Join(strs, " "))
	}
}
