package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// subsetSumExists는 Meet in the Middle 기법으로 부분집합 합이 목표값과 같은지 판별한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - s: 목표 합
//
// [반환값]
//   - bool: 부분집합 합이 s와 같은 경우 true
func subsetSumExists(arr []int, s int) bool {
	// 여기에 코드를 작성하세요
	_ = sort.SearchInts(nil, 0)
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscan(reader, &n, &s)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	if subsetSumExists(arr, s) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
