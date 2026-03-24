package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// hackerlandRadioTransmitters는 모든 집에 신호를 보내기 위한 최소 송신기 수를 반환한다.
//
// [매개변수]
//   - x: 집 위치 배열
//   - k: 송신기 도달 범위
//
// [반환값]
//   - int: 최소 송신기 수
//
// [알고리즘 힌트]
//
//	정렬 후 그리디하게 송신기를 배치한다.
//	커버되지 않은 첫 집에서 K 이내 가장 오른쪽 집에 설치하고,
//	설치 위치 + K 밖의 첫 집으로 이동한다.
func hackerlandRadioTransmitters(x []int, k int) int {
	// 집 위치 정렬
	sort.Ints(x)
	n := len(x)

	count := 0
	i := 0

	for i < n {
		// 커버되지 않은 가장 왼쪽 집의 위치
		loc := x[i]

		// 송신기 설치 위치: loc + k 이내인 가장 오른쪽 집
		for i < n && x[i] <= loc+k {
			i++
		}
		// i-1이 송신기 설치 위치
		transmitter := x[i-1]

		// 송신기가 커버하는 범위: transmitter + k
		for i < n && x[i] <= transmitter+k {
			i++
		}

		// 송신기 하나 설치 완료
		count++
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	fmt.Fprintln(writer, hackerlandRadioTransmitters(x, k))
}
