package main

import (
	"bufio"
	"fmt"
	"os"
)

// SinglyLinkedListNode는 단일 연결 리스트의 노드를 나타낸다.
type SinglyLinkedListNode struct {
	Data int
	Next *SinglyLinkedListNode
}

// insertNodeAtPosition은 연결 리스트의 특정 위치에 새 노드를 삽입한다.
//
// [매개변수]
//   - head: 연결 리스트의 헤드 포인터
//   - data: 삽입할 데이터 값
//   - position: 삽입할 위치 (0-based)
//
// [반환값]
//   - *SinglyLinkedListNode: 갱신된 연결 리스트의 헤드 포인터
func insertNodeAtPosition(head *SinglyLinkedListNode, data int, position int) *SinglyLinkedListNode {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 연결 리스트 생성
	var head, tail *SinglyLinkedListNode
	for i := 0; i < n; i++ {
		var data int
		fmt.Fscan(reader, &data)
		node := &SinglyLinkedListNode{Data: data}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	// 삽입할 데이터와 위치 입력
	var data, position int
	fmt.Fscan(reader, &data)
	fmt.Fscan(reader, &position)

	// 핵심 함수 호출
	head = insertNodeAtPosition(head, data, position)

	// 결과 출력
	current := head
	for current != nil {
		fmt.Fprintln(writer, current.Data)
		current = current.Next
	}
}
