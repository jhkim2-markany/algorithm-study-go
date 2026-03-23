package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 거슬러 줘야 할 금액 입력
	var n int
	fmt.Fscan(reader, &n)

	// 동전 종류 (큰 단위부터)
	coins := [4]int{500, 100, 50, 10}
	count := 0

	// 큰 동전부터 최대한 많이 사용하는 그리디 전략
	for _, coin := range coins {
		count += n / coin // 현재 동전으로 거슬러 줄 수 있는 개수
		n %= coin         // 나머지 금액 갱신
	}

	// 결과 출력
	fmt.Fprintln(writer, count)
}
