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
//
// [알고리즘 힌트]
//
//	입력 스택에 원소를 푸시한다.
func (q *MyQueue) Enqueue(x int) {
	// 입력 스택에 푸시
	q.inStack = append(q.inStack, x)
}

// transfer는 입력 스택의 모든 원소를 출력 스택으로 옮긴다.
func (q *MyQueue) transfer() {
	// 출력 스택이 비어있을 때만 옮김
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			// 입력 스택에서 팝
			top := q.inStack[len(q.inStack)-1]
			q.inStack = q.inStack[:len(q.inStack)-1]
			// 출력 스택에 푸시 (순서가 뒤집힘)
			q.outStack = append(q.outStack, top)
		}
	}
}

// Dequeue는 큐의 앞에서 원소를 제거한다.
//
// [알고리즘 힌트]
//
//	출력 스택이 비어있으면 입력 스택에서 옮긴 뒤 팝한다.
func (q *MyQueue) Dequeue() {
	// 필요시 원소 이동
	q.transfer()
	// 출력 스택에서 팝
	q.outStack = q.outStack[:len(q.outStack)-1]
}

// Peek는 큐의 앞에 있는 원소를 반환한다.
//
// [반환값]
//   - int: 큐의 앞에 있는 원소
//
// [알고리즘 힌트]
//
//	출력 스택이 비어있으면 입력 스택에서 옮긴 뒤 최상위를 반환한다.
func (q *MyQueue) Peek() int {
	// 필요시 원소 이동
	q.transfer()
	// 출력 스택의 최상위 반환
	return q.outStack[len(q.outStack)-1]
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
