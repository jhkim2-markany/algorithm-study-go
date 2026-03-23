package main

import (
	"bufio"
	"fmt"
	"os"
)

// subarraySumEqualsK는 합이 k인 연속 부분배열의 개수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - k: 목표 합
//
// [반환값]
//   - int: 합이 k인 연속 부분배열의 개수
//
// [알고리즘 힌트]
//
//	누적합 + 해시맵을 조합한다.
//	현재 누적합을 prefixSum이라 하면,
//	prefixSum - k가 이전에 등장한 횟수만큼 합이 k인 부분배열이 존재한다.
//	해시맵에 각 누적합의 등장 횟수를 기록하며,
//	초기값으로 freq[0] = 1을 설정한다 (처음부터 합이 k인 경우 처리).
//
//	시간복잡도: O(N), 공간복잡도: O(N)
func subarraySumEqualsK(arr []int, k int) int {
	count := 0
	prefixSum := 0
	freq := make(map[int]int)
	freq[0] = 1

	for _, num := range arr {
		prefixSum += num
		if c, ok := freq[prefixSum-k]; ok {
			count += c
		}
		freq[prefixSum]++
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 N과 목표 합 K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	count := subarraySumEqualsK(arr, k)

	// 결과 출력
	fmt.Fprintln(writer, count)
}
