package main

import (
	"bufio"
	"fmt"
	"os"
)

// icecreamParlor는 가진 돈으로 정확히 두 가지 아이스크림을 살 수 있는 인덱스 쌍을 반환한다.
//
// [매개변수]
//   - m: 가진 돈
//   - arr: 아이스크림 가격 배열
//
// [반환값]
//   - [2]int: 두 아이스크림의 1-indexed 인덱스 (오름차순)
//
// [알고리즘 힌트]
//
//	해시맵에 가격→인덱스를 저장하면서 순회한다.
//	현재 가격의 보수(m - price)가 해시맵에 있으면 답이다.
func icecreamParlor(m int, arr []int) [2]int {
	// 가격 → 인덱스 해시맵
	seen := make(map[int]int)

	for i, price := range arr {
		// 보수 계산
		complement := m - price

		// 보수가 이미 해시맵에 있으면 답 반환
		if idx, ok := seen[complement]; ok {
			return [2]int{idx + 1, i + 1}
		}

		// 현재 가격과 인덱스를 해시맵에 저장
		seen[price] = i
	}

	return [2]int{}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var m, n int
		fmt.Fscan(reader, &m)
		fmt.Fscan(reader, &n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		result := icecreamParlor(m, arr)
		fmt.Fprintf(writer, "%d %d\n", result[0], result[1])
	}
}
