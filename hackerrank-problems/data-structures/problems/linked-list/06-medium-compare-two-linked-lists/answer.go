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

// compareLists는 두 연결 리스트가 동일한지 비교한다.
//
// [매개변수]
//   - head1: 첫 번째 연결 리스트의 헤드 포인터
//   - head2: 두 번째 연결 리스트의 헤드 포인터
//
// [반환값]
//   - int: 동일하면 1, 다르면 0
//
// [알고리즘 힌트]
//
//	두 리스트를 동시에 순회하면서 각 노드의 데이터를 비교한다.
//	데이터가 다르거나 한쪽만 먼저 끝나면 0을 반환한다.
//	둘 다 동시에 nil이 되면 1을 반환한다.
func compareLists(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) int {
	// 두 리스트를 동시에 순회
	curr1 := head1
	curr2 := head2

	for curr1 != nil && curr2 != nil {
		// 데이터가 다르면 0 반환
		if curr1.Data != curr2.Data {
			return 0
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	// 둘 다 nil이면 동일, 아니면 길이가 다름
	if curr1 == nil && curr2 == nil {
		return 1
	}
	return 0
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
		fmt.Fprintln(writer, compareLists(head1, head2))
	}
}
