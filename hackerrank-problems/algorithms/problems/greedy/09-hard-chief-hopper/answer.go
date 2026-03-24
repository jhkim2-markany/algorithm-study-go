package main

import (
	"bufio"
	"fmt"
	"os"
)

// chiefHopper는 모든 건물을 통과할 수 있는 최소 초기 에너지를 반환한다.
//
// [매개변수]
//   - arr: 건물 높이 배열
//
// [반환값]
//   - int: 최소 초기 에너지
//
// [알고리즘 힌트]
//
//	역방향으로 필요한 최소 에너지를 계산한다.
//	E_new = 2E - h 에서 E = ceil((E_new + h) / 2)를 역추적한다.
func chiefHopper(arr []int) int {
	// 마지막 건물 이후 에너지는 0 이상이면 됨
	energy := 0

	// 역순으로 각 건물에 필요한 최소 에너지 계산
	for i := len(arr) - 1; i >= 0; i-- {
		// E = ceil((energy + arr[i]) / 2)
		// 정수 올림: (a + b - 1) / b 대신 (energy + arr[i] + 1) / 2
		energy = (energy + arr[i] + 1) / 2
	}

	return energy
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 건물 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 건물 높이 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := chiefHopper(arr)
	fmt.Fprintln(writer, result)
}
