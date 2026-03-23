package main

import (
	"bufio"
	"fmt"
	"os"
)

// countEqualPairs는 이중 다항식 해싱을 이용하여 동일한 문자열 쌍의 개수를 반환한다.
//
// [매개변수]
//   - strs: 문자열 배열
//
// [반환값]
//   - int64: 동일한 문자열 쌍의 개수 (k개의 같은 문자열이면 k*(k-1)/2)
func countEqualPairs(strs []string) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &strs[i])
	}

	fmt.Fprintln(writer, countEqualPairs(strs))
}
