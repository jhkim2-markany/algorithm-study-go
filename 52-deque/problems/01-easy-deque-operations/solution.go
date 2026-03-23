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

	// 입력: 명령의 수
	var n int
	fmt.Fscan(reader, &n)

	// 슬라이스 기반 덱 구현
	deque := []int{}

	for i := 0; i < n; i++ {
		var cmd string
		fmt.Fscan(reader, &cmd)

		switch cmd {
		case "push_front":
			// 덱의 앞에 원소를 추가한다
			var x int
			fmt.Fscan(reader, &x)
			deque = append([]int{x}, deque...)

		case "push_back":
			// 덱의 뒤에 원소를 추가한다
			var x int
			fmt.Fscan(reader, &x)
			deque = append(deque, x)

		case "pop_front":
			// 덱의 앞 원소를 제거하고 출력한다
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[0])
				deque = deque[1:]
			}

		case "pop_back":
			// 덱의 뒤 원소를 제거하고 출력한다
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[len(deque)-1])
				deque = deque[:len(deque)-1]
			}

		case "size":
			// 덱의 크기를 출력한다
			fmt.Fprintln(writer, len(deque))

		case "empty":
			// 덱이 비어 있으면 1, 아니면 0을 출력한다
			if len(deque) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}

		case "front":
			// 덱의 앞 원소를 출력한다
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[0])
			}

		case "back":
			// 덱의 뒤 원소를 출력한다
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[len(deque)-1])
			}
		}
	}
}
