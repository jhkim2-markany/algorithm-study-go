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

// getNode는 연결 리스트의 끝에서 positionFromTail번째 노드의 데이터를 반환한다.
//
// [매개변수]
//   - head: 연결 리스트의 헤드 포인터
//   - positionFromTail: 끝에서부터의 위치 (0-based)
//
// [반환값]
//   - int: 해당 위치 노드의 데이터 값
//
// [알고리즘 힌트]
//
//	투 포인터 기법을 사용한다.
//	fast 포인터를 positionFromTail만큼 먼저 이동시킨 뒤,
//	두 포인터를 동시에 이동시키면 fast가 끝에 도달했을 때
//	slow가 원하는 위치에 있게 된다.
func getNode(head *SinglyLinkedListNode, positionFromTail int) int {
	// 두 포인터를 헤드에 위치
	fast := head
	slow := head

	// fast를 positionFromTail만큼 먼저 이동
	for i := 0; i < positionFromTail; i++ {
		fast = fast.Next
	}

	// 두 포인터를 동시에 이동
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// slow가 끝에서 positionFromTail번째 노드
	return slow.Data
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

		var head, tail *SinglyLinkedListNode
		for j := 0; j < n; j++ {
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

		var position int
		fmt.Fscan(reader, &position)
		fmt.Fprintln(writer, getNode(head, position))
	}
}
