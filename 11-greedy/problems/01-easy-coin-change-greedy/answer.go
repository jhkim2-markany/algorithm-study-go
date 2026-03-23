package main

import (
	"bufio"
	"fmt"
	"os"
)

// coinChangeGreedy는 주어진 금액을 거슬러 주는 데 필요한 최소 동전 수를 반환한다.
//
// [매개변수]
//   - n: 거슬러 줘야 할 금액 (10의 배수)
//
// [반환값]
//   - int: 필요한 최소 동전 수
//
// [알고리즘 힌트]
//
//	그리디: 큰 동전부터 최대한 많이 사용한다.
//	동전 종류 [500, 100, 50, 10]을 순서대로 처리하며,
//	각 동전으로 n / coin 개를 사용하고 나머지를 갱신한다.
//
//	시간복잡도: O(1) (동전 종류가 고정)
func coinChangeGreedy(n int) int {
	coins := [4]int{500, 100, 50, 10}
	count := 0
	for _, coin := range coins {
		count += n / coin
		n %= coin
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	result := coinChangeGreedy(n)
	fmt.Fprintln(writer, result)
}
