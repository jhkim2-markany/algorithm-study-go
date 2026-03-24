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

// printLinkedList는 연결 리스트의 모든 원소를 순서대로 출력한다.
//
// [매개변수]
//   - head: 연결 리스트의 헤드 포인터
//
// [반환값]
//   - 없음 (표준 출력으로 각 노드의 데이터를 한 줄에 하나씩 출력)
func printLinkedList(head *SinglyLinkedListNode) {
	// 여기에 코드를 작성하세요
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
	printLinkedList(head)
}
