package main

import (
	"bufio"
	"fmt"
	"os"
)

// cutLanCable은 랜선들을 길이 l로 잘라 n개 이상 만들 수 있는
// 최대 랜선 길이를 반환한다.
//
// [매개변수]
//   - cables: 각 랜선의 길이 배열
//   - n: 필요한 랜선 개수
//
// [반환값]
//   - int: 조건을 만족하는 최대 랜선 길이
//
// [알고리즘 힌트]
//
//	파라메트릭 서치: 랜선 길이를 이진 탐색으로 찾는다.
//	결정 함수: 길이 mid로 잘랐을 때 만들 수 있는 랜선 수 >= n인가?
//	각 랜선에서 cable / mid 개를 만들 수 있다.
//	조건 만족 시 lo = mid + 1, 불만족 시 hi = mid - 1.
//
//	시간복잡도: O(K log L), L은 최대 랜선 길이
func cutLanCable(cables []int, n int) int {
	maxLen := 0
	for _, c := range cables {
		if c > maxLen {
			maxLen = c
		}
	}

	lo, hi := 1, maxLen
	result := 0

	for lo <= hi {
		mid := (lo + hi) / 2
		count := 0
		for _, cable := range cables {
			count += cable / mid
		}
		if count >= n {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 랜선 수 K와 필요한 랜선 수 N 입력
	var k, n int
	fmt.Fscan(reader, &k, &n)

	// 각 랜선의 길이 입력
	cables := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &cables[i])
	}

	// 핵심 함수 호출
	result := cutLanCable(cables, n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
