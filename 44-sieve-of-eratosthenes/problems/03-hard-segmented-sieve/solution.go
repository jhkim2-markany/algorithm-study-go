package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// segmentedSieve는 세그먼트 체를 이용하여 [l, r] 구간의 소수 개수를 구한다.
//
// [매개변수]
//   - l: 구간의 시작값
//   - r: 구간의 끝값
//
// [반환값]
//   - int: [l, r] 구간의 소수 개수
func segmentedSieve(l, r int64) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l, r int64
	fmt.Fscan(reader, &l, &r)

	fmt.Fprintln(writer, segmentedSieve(l, r))

	_ = math.Sqrt // 패키지 사용 보장
}
