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
//
// [알고리즘 힌트]
//
//	헤드부터 시작하여 Next 포인터를 따라가며 순회한다.
//	각 노드에서 Data를 출력하고, nil을 만나면 종료한다.
func printLinkedList(head *SinglyLinkedListNode) {
	// 현재 노드를 헤드로 초기화
	current := head

	// nil이 아닌 동안 순회
	for current != nil {
		// 현재 노드의 데이터 출력
		fmt.Println(current.Data)
		// 다음 노드로 이동
		current = current.Next
	}
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
