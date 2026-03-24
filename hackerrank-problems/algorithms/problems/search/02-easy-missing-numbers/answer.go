package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// missingNumbers는 원본 배열에는 있지만 첫 번째 배열에는 빠진 숫자들을 반환한다.
//
// [매개변수]
//   - arr: 일부 숫자가 빠진 배열
//   - brr: 원본 배열
//
// [반환값]
//   - []int: 빠진 숫자들 (오름차순)
//
// [알고리즘 힌트]
//
//	해시맵으로 원본 배열의 빈도를 세고, 첫 번째 배열의 빈도를 차감한다.
//	빈도가 양수인 숫자를 수집하여 정렬한다.
func missingNumbers(arr []int, brr []int) []int {
	// 원본 배열의 빈도 계산
	freq := make(map[int]int)
	for _, v := range brr {
		freq[v]++
	}

	// 첫 번째 배열의 빈도 차감
	for _, v := range arr {
		freq[v]--
	}

	// 빈도가 양수인 숫자 수집
	var result []int
	for k, v := range freq {
		if v > 0 {
			result = append(result, k)
		}
	}

	// 오름차순 정렬
	sort.Ints(result)
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	brr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &brr[i])
	}

	result := missingNumbers(arr, brr)
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
