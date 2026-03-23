package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// abs는 정수의 절댓값을 반환한다
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// checkPermutation은 정렬된 배열의 순열 중 인접 원소 차이가
// 모두 k 이하인 순열을 찾아 반환한다.
//
// [매개변수]
//   - arr: 오름차순 정렬된 정수 배열
//   - k: 인접 원소 간 허용되는 최대 차이
//
// [반환값]
//   - []int: 조건을 만족하는 사전순 가장 앞선 순열
//   - bool: 조건을 만족하는 순열을 찾았으면 true, 없으면 false
//
// [알고리즘 힌트]
//
//	백트래킹(재귀)으로 순열을 생성하며 조건을 검사한다.
//	used[] 배열로 이미 사용한 원소를 추적하고,
//	현재 순열의 마지막 원소와 다음 후보 원소의 차이가 k를 초과하면
//	가지치기(pruning)하여 탐색을 줄인다.
//	입력이 정렬되어 있으므로 첫 번째로 찾은 순열이 사전순 가장 앞선다.
//
//	시간복잡도: 최악 O(N!), 가지치기로 실제로는 훨씬 빠름
func checkPermutation(arr []int, k int) ([]int, bool) {
	n := len(arr)
	result := make([]int, 0, n)
	used := make([]bool, n)
	found := false

	var permute func()
	permute = func() {
		if found {
			return
		}
		if len(result) == n {
			found = true
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			// 인접 원소 차이 조건 검사 (가지치기)
			if len(result) > 0 && abs(result[len(result)-1]-arr[i]) > k {
				continue
			}
			used[i] = true
			result = append(result, arr[i])
			permute()
			if found {
				return
			}
			result = result[:len(result)-1]
			used[i] = false
		}
	}

	permute()

	if found {
		return result, true
	}
	return nil, false
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

	// 핵심 함수 호출
	result, found := checkPermutation(arr, k)

	// 결과 출력
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
