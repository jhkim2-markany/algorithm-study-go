package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// closestSubsetSum은 Meet in the Middle 기법으로 목표값에 가장 가까운 부분집합 합을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - s: 목표값
//
// [반환값]
//   - int: 목표값에 가장 가까운 부분집합 합
func closestSubsetSum(arr []int, s int) int {
	// 여기에 코드를 작성하세요
	_ = sort.SearchInts(nil, 0)
	return 0
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

	fmt.Fprintln(writer, closestSubsetSum(arr, s))
}
