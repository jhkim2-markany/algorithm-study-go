package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 수열 길이
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// O(N log N) 이분 탐색 기반 LIS
	tails := []int{} // tails[k] = 길이 k+1인 증가 부분 수열의 마지막 원소 최솟값

	for _, x := range a {
		// tails에서 x 이상인 첫 번째 위치를 찾는다 (lower_bound)
		pos := sort.SearchInts(tails, x)
		if pos == len(tails) {
			// x가 모든 원소보다 크면 뒤에 추가
			tails = append(tails, x)
		} else {
			// 해당 위치를 x로 교체
			tails[pos] = x
		}
	}

	// 출력: LIS 길이
	fmt.Fprintln(writer, len(tails))
}
