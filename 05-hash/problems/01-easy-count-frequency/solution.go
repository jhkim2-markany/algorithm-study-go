package main

import (
	"bufio"
	"fmt"
	"os"
)

// FreqPair는 값과 등장 횟수를 저장하는 구조체이다.
type FreqPair struct {
	Value int
	Count int
}

// countFrequency는 배열의 각 원소 등장 횟수를 구하여
// 등장 횟수 내림차순, 같으면 값 오름차순으로 정렬된 결과를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - []FreqPair: (값, 등장횟수) 쌍의 배열 (등장횟수 내림차순, 값 오름차순)
func countFrequency(arr []int) []FreqPair {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := countFrequency(arr)

	// 결과 출력
	for _, p := range result {
		fmt.Fprintf(writer, "%d %d\n", p.Value, p.Count)
	}
}
