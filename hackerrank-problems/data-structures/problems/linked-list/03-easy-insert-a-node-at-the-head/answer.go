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

// insertNodeAtHead는 연결 리스트의 맨 앞에 새 노드를 삽입한다.
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
//	새 노드를 생성하고, 새 노드의 Next를 기존 헤드로 설정한다.
//	새 노드가 새로운 헤드가 된다.
func insertNodeAtHead(head *SinglyLinkedListNode, data int) *SinglyLinkedListNode {
	// 새 노드 생성
	newNode := &SinglyLinkedListNode{Data: data}

	// 새 노드의 Next를 기존 헤드로 설정
	newNode.Next = head

	// 새 노드를 헤드로 반환
	return newNode
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 연결 리스트 생성 (head 삽입)
	var head *SinglyLinkedListNode
	for i := 0; i < n; i++ {
		var data int
		fmt.Fscan(reader, &data)
		head = insertNodeAtHead(head, data)
	}

	// 결과 출력
	current := head
	for current != nil {
		fmt.Fprintln(writer, current.Data)
		current = current.Next
	}
}
