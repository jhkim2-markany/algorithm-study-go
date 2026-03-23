package main

import (
	"bufio"
	"fmt"
	"os"
)

// countPrimesInRange는 [l, r] 범위의 소수 개수를 반환한다.
//
// [매개변수]
//   - l: 범위의 시작 (1 이상)
//   - r: 범위의 끝 (l 이상)
//
// [반환값]
//   - int: [l, r] 범위에 포함된 소수의 개수
func countPrimesInRange(l, r int64) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, r int64
	fmt.Fscan(reader, &l, &r)

	fmt.Fprintln(writer, countPrimesInRange(l, r))
}
