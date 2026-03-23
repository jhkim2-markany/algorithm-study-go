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

	// 배열 크기와 목표 합 입력
	var n, s int
	fmt.Fscan(reader, &n, &s)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	count := 0

	// 비트마스크로 모든 부분집합을 탐색 (빈 집합 제외)
	// mask가 1부터 시작하여 빈 집합을 제외한다
	for mask := 1; mask < (1 << n); mask++ {
		sum := 0
		// 각 비트를 확인하여 선택된 원소의 합을 계산
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += arr[i]
			}
		}
		// 합이 목표값과 같으면 카운트 증가
		if sum == s {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}
