package main

import (
	"bufio"
	"fmt"
	"os"
)

// DoublyLinkedListNode는 이중 연결 리스트의 노드를 나타낸다.
type DoublyLinkedListNode struct {
	Data int
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

// reverseDoublyLinkedList는 이중 연결 리스트를 뒤집는다.
//
// [매개변수]
//   - head: 이중 연결 리스트의 헤드 포인터
//
// [반환값]
//   - *DoublyLinkedListNode: 뒤집힌 이중 연결 리스트의 헤드 포인터
func reverseDoublyLinkedList(head *DoublyLinkedListNode) *DoublyLinkedListNode {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		// 이중 연결 리스트 생성
		var head, tail *DoublyLinkedListNode
		for j := 0; j < n; j++ {
			var data int
			fmt.Fscan(reader, &data)
			node := &DoublyLinkedListNode{Data: data}
			if head == nil {
				head = node
				tail = node
			} else {
				tail.Next = node
				node.Prev = tail
				tail = node
			}
		}

		// 핵심 함수 호출
		head = reverseDoublyLinkedList(head)

		// 결과 출력
		current := head
		for current != nil {
			if current != head {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, current.Data)
			current = current.Next
		}
		fmt.Fprintln(writer)
	}
}
