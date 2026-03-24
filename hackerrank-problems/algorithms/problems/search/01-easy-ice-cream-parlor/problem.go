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
func icecreamParlor(m int, arr []int) [2]int {
	// 여기에 코드를 작성하세요
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
