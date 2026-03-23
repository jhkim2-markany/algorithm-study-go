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

	// 입력: 수열의 크기
	var n int
	fmt.Fscan(reader, &n)

	// 입력: N개의 정수
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 좌표 압축: 정렬 후 중복 제거
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	// 중복 제거하여 고유 좌표 배열 생성
	unique := []int{sorted[0]}
	for i := 1; i < n; i++ {
		if sorted[i] != sorted[i-1] {
			unique = append(unique, sorted[i])
		}
	}

	// 이진 탐색으로 각 원소의 압축된 좌표(순위) 출력
	for i := 0; i < n; i++ {
		// sort.SearchInts: 정렬된 배열에서 값의 위치를 찾는다
		rank := sort.SearchInts(unique, arr[i])
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, rank)
	}
	fmt.Fprintln(writer)
}
