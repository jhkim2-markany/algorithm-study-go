package main

import (
	"bufio"
	"fmt"
	"os"
)

// modInverse는 페르마 소정리를 이용하여 a의 모듈러 역원 a^(m-2) mod m을 반환한다.
//
// [매개변수]
//   - a: 역원을 구할 정수
//   - m: 소수인 모듈러 값
//
// [반환값]
//   - int64: a^(-1) mod m
func modInverse(a, m int64) int64 {
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
		var a, m int64
		fmt.Fscan(reader, &a, &m)
		fmt.Fprintln(writer, modInverse(a, m))
	}
}
