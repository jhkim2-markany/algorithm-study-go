package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcdLcm은 두 자연수의 최대공약수와 최소공배수를 반환한다.
//
// [매개변수]
//   - a: 첫 번째 자연수
//   - b: 두 번째 자연수
//
// [반환값]
//   - int: 최대공약수(GCD)
//   - int: 최소공배수(LCM)
func gcdLcm(a, b int) (int, int) {
	// 여기에 코드를 작성하세요
	return 0, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 두 자연수 입력
	var a, b int
	fmt.Fscan(reader, &a, &b)

	// 핵심 함수 호출
	g, l := gcdLcm(a, b)

	// 결과 출력
	fmt.Fprintln(writer, g)
	fmt.Fprintln(writer, l)
}
