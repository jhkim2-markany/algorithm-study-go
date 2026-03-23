package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	// 누적합 + 해시맵을 이용한 풀이
	// 현재 누적합에서 K를 뺀 값이 이전에 등장한 횟수를 센다
	count := 0
	prefixSum := 0
	// 해시맵: 누적합 값 → 등장 횟수
	freq := make(map[int]int)
	// 누적합이 0인 경우를 위해 초기값 설정
	freq[0] = 1

	for i := 0; i < n; i++ {
		// 현재까지의 누적합 계산
		prefixSum += arr[i]

		// prefixSum - K가 이전에 등장한 적 있으면
		// 그 지점부터 현재까지의 부분배열 합이 K이다
		if c, ok := freq[prefixSum-k]; ok {
			count += c
		}

		// 현재 누적합을 해시맵에 기록
		freq[prefixSum]++
	}

	fmt.Fprintln(writer, count)
}
