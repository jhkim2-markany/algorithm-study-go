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

	// O(N log N) LIS + 경로 역추적
	tails := []int{}      // tails[k] = 길이 k+1인 증가 부분 수열의 마지막 원소 최솟값
	pos := make([]int, n) // pos[i] = a[i]가 tails에서 들어간 위치

	for i, x := range a {
		// tails에서 x 이상인 첫 번째 위치를 찾는다
		p := sort.SearchInts(tails, x)
		if p == len(tails) {
			tails = append(tails, x)
		} else {
			tails[p] = x
		}
		pos[i] = p // a[i]가 LIS에서 차지하는 위치 기록
	}

	// 역추적: LIS 길이부터 역순으로 실제 LIS 원소를 복원한다
	lisLen := len(tails)
	result := make([]int, lisLen)
	target := lisLen - 1

	for i := n - 1; i >= 0; i-- {
		if pos[i] == target {
			result[target] = a[i]
			target--
			if target < 0 {
				break
			}
		}
	}

	// 출력: LIS 길이와 실제 LIS
	fmt.Fprintln(writer, lisLen)
	for i := 0; i < lisLen; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, result[i])
	}
	fmt.Fprintln(writer)
}
