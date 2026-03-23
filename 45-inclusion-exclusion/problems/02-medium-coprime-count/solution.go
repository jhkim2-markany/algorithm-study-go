package main

import (
	"bufio"
	"fmt"
	"os"
)

// eulerPhi는 오일러 피 함수를 계산한다.
// n 이하의 양의 정수 중 n과 서로소인 수의 개수를 반환한다.
//
// [매개변수]
//   - n: 양의 정수
//
// [반환값]
//   - int64: φ(n) 값 (n과 서로소인 수의 개수)
func eulerPhi(n int64) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscan(reader, &n)
		fmt.Fprintln(writer, eulerPhi(n))
	}
}
