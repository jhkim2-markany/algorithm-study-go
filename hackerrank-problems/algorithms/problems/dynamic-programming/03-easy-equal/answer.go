package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// equal은 모든 동료가 같은 수의 초콜릿을 갖도록 하는 최소 연산 횟수를 반환한다.
//
// [매개변수]
//   - arr: 각 동료의 초콜릿 수 배열
//
// [반환값]
//   - int: 최소 연산 횟수
//
// [알고리즘 힌트]
//
//	"나머지에게 주기"를 "한 명에게서 빼기"로 변환한다.
//	최솟값 기준으로 0~4를 추가로 뺀 5가지 목표값을 시도하여
//	각 차이를 5, 2, 1 단위로 줄이는 최소 연산을 계산한다.
func equal(arr []int) int {
	// 배열의 최솟값 찾기
	minVal := arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
	}

	best := math.MaxInt64

	// 목표값을 minVal, minVal-1, ..., minVal-4로 시도
	for delta := 0; delta < 5; delta++ {
		target := minVal - delta
		total := 0

		for _, v := range arr {
			diff := v - target
			// diff를 5, 2, 1 단위로 줄이는 최소 연산 횟수
			total += diff / 5
			diff %= 5
			total += diff / 2
			diff %= 2
			total += diff
		}

		if total < best {
			best = total
		}
	}

	return best
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		// 동료 수 입력
		var n int
		fmt.Fscan(reader, &n)

		// 초콜릿 수 배열 입력
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		// 핵심 함수 호출 및 결과 출력
		result := equal(arr)
		fmt.Fprintln(writer, result)
	}
}
