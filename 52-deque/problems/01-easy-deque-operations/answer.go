package main

import (
	"bufio"
	"fmt"
	"os"
)

// processDequeOps는 덱 명령어 목록을 처리하고 각 출력 명령의 결과를 반환한다.
//
// [매개변수]
//   - commands: 각 원소가 [명령어, 값] 형태인 명령 배열 (값이 없으면 0)
//
// [반환값]
//   - []int: 출력이 필요한 명령(pop_front, pop_back, size, empty, front, back)의 결과 배열
//
// [알고리즘 힌트]
//
//	슬라이스 기반 덱: push_front는 앞에 삽입, push_back은 뒤에 삽입,
//	pop/front/back은 비어있으면 -1을 반환한다.
func processDequeOps(commands []([2]interface{})) []int {
	// 이 문제는 명령어를 직접 main에서 처리하는 구조이다
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	deque := []int{}

	for i := 0; i < n; i++ {
		var cmd string
		fmt.Fscan(reader, &cmd)

		switch cmd {
		case "push_front":
			var x int
			fmt.Fscan(reader, &x)
			deque = append([]int{x}, deque...)
		case "push_back":
			var x int
			fmt.Fscan(reader, &x)
			deque = append(deque, x)
		case "pop_front":
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[0])
				deque = deque[1:]
			}
		case "pop_back":
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[len(deque)-1])
				deque = deque[:len(deque)-1]
			}
		case "size":
			fmt.Fprintln(writer, len(deque))
		case "empty":
			if len(deque) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "front":
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[0])
			}
		case "back":
			if len(deque) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, deque[len(deque)-1])
			}
		}
	}
}
