package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// lisLength는 주어진 수열에서 최장 증가 부분 수열(LIS)의 길이를 반환한다.
//
// [매개변수]
//   - a: 정수 수열
//
// [반환값]
//   - int: LIS의 길이
//
// [알고리즘 힌트]
//   - O(N log N) 이분 탐색 기반 LIS 알고리즘을 사용한다
//   - tails[k] = 길이 k+1인 증가 부분 수열의 마지막 원소 최솟값을 유지한다
//   - 각 원소에 대해 tails에서 lower_bound를 찾아 교체하거나 뒤에 추가한다
//   - 최종 tails의 길이가 LIS 길이이다
func lisLength(a []int) int {
	tails := []int{}

	for _, x := range a {
		pos := sort.SearchInts(tails, x)
		if pos == len(tails) {
			tails = append(tails, x)
		} else {
			tails[pos] = x
		}
	}

	return len(tails)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	fmt.Fprintln(writer, lisLength(a))
}
