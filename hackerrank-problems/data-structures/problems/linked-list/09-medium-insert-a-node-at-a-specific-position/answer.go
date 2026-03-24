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
//
// [알고리즘 힌트]
//
//	위치가 0이면 새 노드를 헤드로 설정한다.
//	그 외에는 (position-1)번째 노드까지 이동한 뒤,
//	새 노드를 해당 위치에 끼워 넣는다.
func insertNodeAtPosition(head *SinglyLinkedListNode, data int, position int) *SinglyLinkedListNode {
	// 새 노드 생성
	newNode := &SinglyLinkedListNode{Data: data}

	// 헤드 앞에 삽입하는 경우
	if position == 0 {
		newNode.Next = head
		return newNode
	}

	// (position-1)번째 노드까지 순회
	current := head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	// 새 노드를 현재 노드와 다음 노드 사이에 삽입
	newNode.Next = current.Next
	current.Next = newNode

	return head
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
