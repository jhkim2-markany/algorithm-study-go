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

// reverseLinkedList는 연결 리스트를 뒤집는다.
//
// [매개변수]
//   - head: 연결 리스트의 헤드 포인터
//
// [반환값]
//   - *SinglyLinkedListNode: 뒤집힌 연결 리스트의 헤드 포인터
//
// [알고리즘 힌트]
//
//	세 개의 포인터(prev, current, next)를 사용하여
//	각 노드의 Next 방향을 반대로 바꾼다.
//	순회가 끝나면 prev가 새로운 헤드가 된다.
func reverseLinkedList(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// 이전 노드 포인터 초기화
	var prev *SinglyLinkedListNode
	current := head

	for current != nil {
		// 다음 노드 저장
		next := current.Next
		// 현재 노드의 방향을 반대로 변경
		current.Next = prev
		// 포인터 이동
		prev = current
		current = next
	}

	// prev가 새로운 헤드
	return prev
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

	// 핵심 함수 호출
	head = reverseLinkedList(head)

	// 결과 출력
	current := head
	for current != nil {
		fmt.Fprintln(writer, current.Data)
		current = current.Next
	}
}
