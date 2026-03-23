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

	// 입력: 배열 크기
	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
	}

	// A+B의 모든 합을 구한다 (N² 개)
	ab := make([]int, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ab = append(ab, a[i]+b[j])
		}
	}

	// C+D의 모든 합을 구한다 (N² 개)
	cd := make([]int, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cd = append(cd, c[i]+d[j])
		}
	}

	// cd를 정렬하여 이분 탐색 준비
	sort.Ints(cd)

	// ab의 각 원소에 대해 cd에서 보완값의 개수를 센다
	count := 0
	for _, x := range ab {
		target := -x
		// lower_bound: target 이상인 첫 위치
		lo := sort.SearchInts(cd, target)
		// upper_bound: target 초과인 첫 위치
		hi := sort.SearchInts(cd, target+1)
		count += hi - lo // target과 같은 원소의 개수
	}

	// 결과 출력
	fmt.Fprintln(writer, count)
}
