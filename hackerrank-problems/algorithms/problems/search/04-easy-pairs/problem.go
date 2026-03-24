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
func pairs(k int, arr []int) int {
	// 여기에 코드를 작성하세요
	return 0
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
