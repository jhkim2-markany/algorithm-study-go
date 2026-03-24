package main

import (
	"bufio"
	"fmt"
	"os"
)

// pairs는 차이가 정확히 k인 쌍의 개수를 반환한다.
//
// [매개변수]
//   - k: 목표 차이값
//   - arr: 정수 배열 (모든 원소가 서로 다름)
//
// [반환값]
//   - int: 차이가 k인 쌍의 개수
//
// [알고리즘 힌트]
//
//	해시셋에 모든 원소를 저장한 뒤, 각 원소 x에 대해
//	x + k가 셋에 있으면 쌍으로 센다.
func pairs(k int, arr []int) int {
	// 해시셋에 모든 원소 저장
	set := make(map[int]bool)
	for _, v := range arr {
		set[v] = true
	}

	// 각 원소에 대해 x + k가 존재하는지 확인
	count := 0
	for _, v := range arr {
		if set[v+k] {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	fmt.Fprintln(writer, pairs(k, arr))
}
