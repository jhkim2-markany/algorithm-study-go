package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 결과를 저장할 전역 변수
var result []int
var found bool

// abs 함수는 정수의 절댓값을 반환한다
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// permute 함수는 재귀적으로 순열을 생성하며 조건을 검사한다
func permute(arr []int, perm []int, used []bool, k int) {
	if found {
		return
	}

	// 순열이 완성된 경우
	if len(perm) == len(arr) {
		found = true
		result = make([]int, len(perm))
		copy(result, perm)
		return
	}

	for i := 0; i < len(arr); i++ {
		if used[i] {
			continue
		}

		// 인접 원소 차이 조건 검사 (가지치기)
		if len(perm) > 0 && abs(perm[len(perm)-1]-arr[i]) > k {
			continue
		}

		used[i] = true
		perm = append(perm, arr[i])
		permute(arr, perm, used, k)
		perm = perm[:len(perm)-1]
		used[i] = false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기와 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 사전순으로 가장 앞서는 순열을 찾기 위해 정렬
	sort.Ints(arr)

	// 모든 순열을 탐색하여 조건을 만족하는 첫 번째 순열 찾기
	used := make([]bool, n)
	perm := make([]int, 0, n)
	permute(arr, perm, used, k)

	if found {
		for i, v := range result {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
