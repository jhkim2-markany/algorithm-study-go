package main

import (
	"bufio"
	"fmt"
	"os"
)

// modInverse는 확장 유클리드 알고리즘을 이용하여 a의 모듈러 역원을 구한다.
// a * x ≡ 1 (mod m)을 만족하는 x를 반환한다.
//
// [매개변수]
//   - a: 역원을 구할 정수
//   - m: 모듈러 값
//
// [반환값]
//   - int: 모듈러 역원 (0 이상 m 미만), 역원이 없으면 -1
func modInverse(a, m int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var a, m int
		fmt.Fscan(reader, &a, &m)
		fmt.Fprintln(writer, modInverse(a, m))
	}
}
