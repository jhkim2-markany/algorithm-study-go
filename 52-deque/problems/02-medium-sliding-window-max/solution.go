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

	// 입력: 배열 크기 N, 윈도우 크기 K
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 모노톤 덱을 이용한 슬라이딩 윈도우 최댓값
	// 덱에는 인덱스를 저장하며, 앞쪽이 항상 최댓값의 인덱스이다
	deque := make([]int, 0, n)

	for i := 0; i < n; i++ {
		// 윈도우 범위를 벗어난 인덱스를 앞에서 제거한다
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 현재 원소보다 작거나 같은 원소의 인덱스를 뒤에서 제거한다
		// 이들은 현재 원소가 있는 한 최댓값이 될 수 없다
		for len(deque) > 0 && arr[deque[len(deque)-1]] <= arr[i] {
			deque = deque[:len(deque)-1]
		}

		// 현재 인덱스를 덱의 뒤에 추가한다
		deque = append(deque, i)

		// 윈도우가 완성된 시점(i >= k-1)부터 최댓값을 출력한다
		if i >= k-1 {
			if i > k-1 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, arr[deque[0]])
		}
	}
	fmt.Fprintln(writer)
}
