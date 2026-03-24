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

// mergeLists는 정렬된 두 연결 리스트를 하나로 병합한다.
//
// [매개변수]
//   - head1: 첫 번째 정렬된 연결 리스트의 헤드 포인터
//   - head2: 두 번째 정렬된 연결 리스트의 헤드 포인터
//
// [반환값]
//   - *SinglyLinkedListNode: 병합된 정렬 연결 리스트의 헤드 포인터
func mergeLists(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	// 여기에 코드를 작성하세요
	return nil
}

func buildList(reader *bufio.Reader) *SinglyLinkedListNode {
	var n int
	fmt.Fscan(reader, &n)
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
	return head
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 테스트 케이스 수 입력
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		head1 := buildList(reader)
		head2 := buildList(reader)
		merged := mergeLists(head1, head2)

		current := merged
		for current != nil {
			fmt.Fprintln(writer, current.Data)
			current = current.Next
		}
	}
}
