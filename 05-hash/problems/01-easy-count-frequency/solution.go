package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력 및 해시맵으로 빈도수 계산
	freq := make(map[int]int)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(reader, &num)
		freq[num]++
	}

	// 해시맵의 키-값 쌍을 슬라이스로 변환
	type pair struct {
		value int
		count int
	}
	pairs := make([]pair, 0, len(freq))
	for v, c := range freq {
		pairs = append(pairs, pair{v, c})
	}

	// 등장 횟수 내림차순, 같으면 값 오름차순으로 정렬
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].value < pairs[j].value
	})

	// 결과 출력
	for _, p := range pairs {
		fmt.Fprintf(writer, "%d %d\n", p.value, p.count)
	}
}
