package main

import (
	"bufio"
	"fmt"
	"os"
)

// gcd는 두 정수의 최대공약수를 구한다.
//
// [매개변수]
//   - a: 첫 번째 양의 정수
//   - b: 두 번째 양의 정수
//
// [반환값]
//   - int: a와 b의 최대공약수
func gcd(a, b int) int {
	// 여기에 코드를 작성하세요
	return 0
}

// diceProbability는 주사위를 n번 던져서 합이 s가 되는 확률을 기약분수로 반환한다.
//
// [매개변수]
//   - n: 주사위를 던지는 횟수 (1 이상)
//   - s: 목표 합
//
// [반환값]
//   - int: 기약분수의 분자
//   - int: 기약분수의 분모
func diceProbability(n, s int) (int, int) {
	// 여기에 코드를 작성하세요
	return 0, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscan(reader, &n, &s)

	num, den := diceProbability(n, s)
	fmt.Fprintf(writer, "%d/%d\n", num, den)
}
