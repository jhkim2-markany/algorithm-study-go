package main

import (
	"bufio"
	"fmt"
	"os"
)

// MyQueue는 두 개의 스택으로 구현한 큐이다.
type MyQueue struct {
	inStack  []int
	outStack []int
}

// Enqueue는 큐의 뒤에 원소를 삽입한다.
//
// [매개변수]
//   - x: 삽입할 원소
func (q *MyQueue) Enqueue(x int) {
	// 여기에 코드를 작성하세요
}

// Dequeue는 큐의 앞에서 원소를 제거한다.
func (q *MyQueue) Dequeue() {
	// 여기에 코드를 작성하세요
}

// Peek는 큐의 앞에 있는 원소를 반환한다.
//
// [반환값]
//   - int: 큐의 앞에 있는 원소
func (q *MyQueue) Peek() int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 쿼리 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	queue := &MyQueue{}

	for i := 0; i < n; i++ {
		var qType int
		fmt.Fscan(reader, &qType)
		switch qType {
		case 1:
			var x int
			fmt.Fscan(reader, &x)
			queue.Enqueue(x)
		case 2:
			queue.Dequeue()
		case 3:
			fmt.Fprintln(writer, queue.Peek())
		}
	}
}
