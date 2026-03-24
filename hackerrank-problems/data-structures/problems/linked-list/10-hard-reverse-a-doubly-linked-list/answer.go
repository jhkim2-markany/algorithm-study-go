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
//
// [알고리즘 힌트]
//
//	각 노드의 Next와 Prev 포인터를 교환한다.
//	모든 노드를 순회하며 교환하면 리스트가 뒤집힌다.
//	마지막으로 처리한 노드가 새로운 헤드가 된다.
func reverseDoublyLinkedList(head *DoublyLinkedListNode) *DoublyLinkedListNode {
	// 빈 리스트 처리
	if head == nil {
		return nil
	}

	current := head
	var newHead *DoublyLinkedListNode

	// 모든 노드의 Next와 Prev를 교환
	for current != nil {
		// Next와 Prev 교환
		temp := current.Next
		current.Next = current.Prev
		current.Prev = temp

		// 현재 노드를 새 헤드 후보로 저장
		newHead = current

		// 원래의 Next 방향으로 이동 (교환 후에는 Prev)
		current = temp
	}

	return newHead
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
