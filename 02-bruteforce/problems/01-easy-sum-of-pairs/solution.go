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

	// 배열 크기와 목표 합 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 2중 반복문으로 모든 쌍을 탐색
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 두 수의 합이 목표값과 같으면 출력 후 종료
			if arr[i]+arr[j] == m {
				fmt.Fprintf(writer, "%d %d\n", arr[i], arr[j])
				return
			}
		}
	}

	// 조건을 만족하는 쌍이 없는 경우
	fmt.Fprintln(writer, "NO")
}
