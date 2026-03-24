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

// insertNodeAtTail은 연결 리스트의 끝에 새 노드를 삽입한다.
//
// [매개변수]
//   - head: 연결 리스트의 헤드 포인터
//   - data: 삽입할 데이터 값
//
// [반환값]
//   - *SinglyLinkedListNode: 갱신된 연결 리스트의 헤드 포인터
//
// [알고리즘 힌트]
//
//	리스트가 비어 있으면 새 노드가 곧 헤드이다.
//	비어 있지 않으면 마지막 노드(Next == nil)까지 순회한 뒤
//	마지막 노드의 Next를 새 노드로 연결한다.
func insertNodeAtTail(head *SinglyLinkedListNode, data int) *SinglyLinkedListNode {
	// 새 노드 생성
	newNode := &SinglyLinkedListNode{Data: data}

	// 리스트가 비어 있으면 새 노드가 헤드
	if head == nil {
		return newNode
	}

	// 마지막 노드까지 순회
	current := head
	for current.Next != nil {
		current = current.Next
	}

	// 마지막 노드의 Next를 새 노드로 설정
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

	// 연결 리스트 생성 (tail 삽입)
	var head *SinglyLinkedListNode
	for i := 0; i < n; i++ {
		var data int
		fmt.Fscan(reader, &data)
		head = insertNodeAtTail(head, data)
	}

	// 결과 출력
	current := head
	for current != nil {
		fmt.Fprintln(writer, current.Data)
		current = current.Next
	}
}
