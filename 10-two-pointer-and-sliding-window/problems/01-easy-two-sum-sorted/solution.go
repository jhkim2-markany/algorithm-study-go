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

	// 배열 크기와 목표값 입력
	var n, t int
	fmt.Fscan(reader, &n, &t)

	// 정렬된 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 투 포인터: 양쪽 끝에서 시작
	left, right := 0, n-1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == t {
			// 합이 목표값과 같으면 1-based 인덱스 출력
			fmt.Fprintf(writer, "%d %d\n", left+1, right+1)
			return
		} else if sum < t {
			// 합이 작으면 왼쪽 포인터를 오른쪽으로 이동
			left++
		} else {
			// 합이 크면 오른쪽 포인터를 왼쪽으로 이동
			right--
		}
	}
}
