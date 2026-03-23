package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// lisWithPath는 최장 증가 부분 수열(LIS)의 길이와 실제 LIS를 반환한다.
//
// [매개변수]
//   - a: 정수 수열
//
// [반환값]
//   - int: LIS의 길이
//   - []int: 실제 LIS 원소 배열
//
// [알고리즘 힌트]
//   - O(N log N) 이분 탐색 기반 LIS + 경로 역추적을 사용한다
//   - tails 배열로 LIS 길이를 구하면서 각 원소의 위치(pos)를 기록한다
//   - 역추적: 배열 끝에서부터 pos[i] == target인 원소를 찾아 LIS를 복원한다
func lisWithPath(a []int) (int, []int) {
	n := len(a)
	tails := []int{}
	pos := make([]int, n)

	for i, x := range a {
		p := sort.SearchInts(tails, x)
		if p == len(tails) {
			tails = append(tails, x)
		} else {
			tails[p] = x
		}
		pos[i] = p
	}

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

	return lisLen, result
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

	lisLen, result := lisWithPath(a)

	fmt.Fprintln(writer, lisLen)
	for i := 0; i < lisLen; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, result[i])
	}
	fmt.Fprintln(writer)
}
