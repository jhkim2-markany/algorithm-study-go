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
func insertNodeAtHead(head *SinglyLinkedListNode, data int) *SinglyLinkedListNode {
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
