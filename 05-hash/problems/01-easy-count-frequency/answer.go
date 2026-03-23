package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
//
// [알고리즘 힌트]
//
//	해시맵(map)으로 각 원소의 등장 횟수를 O(N)에 계산한다.
//	해시맵의 키-값 쌍을 슬라이스로 변환한 뒤,
//	sort.Slice로 등장 횟수 내림차순, 값 오름차순으로 정렬한다.
//
//	시간복잡도: O(N + K log K), K는 서로 다른 원소의 수
func countFrequency(arr []int) []FreqPair {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	pairs := make([]FreqPair, 0, len(freq))
	for v, c := range freq {
		pairs = append(pairs, FreqPair{v, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count != pairs[j].Count {
			return pairs[i].Count > pairs[j].Count
		}
		return pairs[i].Value < pairs[j].Value
	})

	return pairs
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
