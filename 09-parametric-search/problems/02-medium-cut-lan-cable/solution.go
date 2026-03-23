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

	// 랜선 수 K와 필요한 랜선 수 N 입력
	var k, n int
	fmt.Fscan(reader, &k, &n)

	// 각 랜선의 길이 입력
	cables := make([]int, k)
	maxLen := 0
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &cables[i])
		if cables[i] > maxLen {
			maxLen = cables[i]
		}
	}

	// 파라메트릭 서치: 랜선 길이의 최댓값을 이진 탐색으로 찾는다
	// 탐색 범위: 1부터 가장 긴 랜선 길이까지
	lo, hi := 1, maxLen
	result := 0

	for lo <= hi {
		mid := (lo + hi) / 2

		// 결정 함수: 길이 mid로 잘랐을 때 N개 이상을 만들 수 있는가?
		count := 0
		for _, cable := range cables {
			count += cable / mid
		}

		if count >= n {
			// 조건 만족: 더 긴 길이도 가능한지 확인
			result = mid
			lo = mid + 1
		} else {
			// 조건 불만족: 길이를 줄인다
			hi = mid - 1
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, result)
}
