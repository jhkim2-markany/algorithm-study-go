package main

import (
	"bufio"
	"fmt"
	"os"
)

// journeyToMoon은 서로 다른 나라에서 2명을 선발하는 경우의 수를 반환한다.
//
// [매개변수]
//   - n: 우주비행사 수
//   - pairs: 같은 나라 출신 쌍 목록
//
// [반환값]
//   - int64: 가능한 조합 수
func journeyToMoon(n int, pairs [][2]int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, p int
	fmt.Fscan(reader, &n, &p)

	pairs := make([][2]int, p)
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &pairs[i][0], &pairs[i][1])
	}

	result := journeyToMoon(n, pairs)
	fmt.Fprintln(writer, result)
}
