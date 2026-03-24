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

// deleteNode는 연결 리스트에서 지정된 위치의 노드를 삭제한다.
//
// [매개변수]
//   - head: 연결 리스트의 헤드 포인터
//   - position: 삭제할 노드의 위치 (0-based)
//
// [반환값]
//   - *SinglyLinkedListNode: 갱신된 연결 리스트의 헤드 포인터
//
// [알고리즘 힌트]
//
//	위치가 0이면 헤드를 다음 노드로 교체한다.
//	그 외에는 (position-1)번째 노드까지 이동한 뒤,
//	해당 노드의 Next를 건너뛰어 연결한다.
func deleteNode(head *SinglyLinkedListNode, position int) *SinglyLinkedListNode {
	// 헤드 노드를 삭제하는 경우
	if position == 0 {
		return head.Next
	}

	// 삭제할 노드의 이전 노드까지 순회
	current := head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	// 이전 노드의 Next를 삭제할 노드의 Next로 연결
	current.Next = current.Next.Next

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

	// 삭제할 위치 입력
	var position int
	fmt.Fscan(reader, &position)

	// 핵심 함수 호출
	head = deleteNode(head, position)

	// 결과 출력
	current := head
	for current != nil {
		fmt.Fprintln(writer, current.Data)
		current = current.Next
	}
}
