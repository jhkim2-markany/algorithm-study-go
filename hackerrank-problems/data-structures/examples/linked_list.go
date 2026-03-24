package main

import "fmt"

// 단일 연결 리스트 (Singly Linked List) - 기본 연산 예시
// 연결 리스트의 삽입, 삭제, 탐색, 출력 연산을 구현한다.
//
// 시간 복잡도:
//   - 맨 앞 삽입 (InsertHead):  O(1)
//   - 맨 뒤 삽입 (InsertTail):  O(N) (N: 노드 수)
//   - 특정 위치 삽입 (InsertAt): O(N)
//   - 삭제 (Delete):            O(N)
//   - 탐색 (Search):            O(N)
//   - 역순 (ReverseList):       O(N)
//   - 출력 (PrintList):         O(N)
//
// 공간 복잡도: O(N) (N개의 노드를 저장)

// Node는 연결 리스트의 노드를 나타낸다.
// Data 필드에 값을 저장하고, Next 필드로 다음 노드를 가리킨다.
type Node struct {
	Data int
	Next *Node
}

// LinkedList는 단일 연결 리스트를 나타낸다.
// Head는 리스트의 첫 번째 노드를 가리킨다.
type LinkedList struct {
	Head *Node
}

// InsertHead 함수는 리스트의 맨 앞에 새 노드를 삽입한다.
// 새 노드의 Next를 기존 Head로 설정하고, Head를 새 노드로 갱신한다.
func (ll *LinkedList) InsertHead(data int) {
	// 새 노드를 생성하고 기존 Head를 다음 노드로 연결
	newNode := &Node{Data: data, Next: ll.Head}
	// Head를 새 노드로 갱신
	ll.Head = newNode
}

// InsertTail 함수는 리스트의 맨 뒤에 새 노드를 삽입한다.
// 리스트가 비어있으면 Head로 설정하고, 아니면 마지막 노드를 찾아 연결한다.
func (ll *LinkedList) InsertTail(data int) {
	newNode := &Node{Data: data}

	// 리스트가 비어있는 경우: 새 노드를 Head로 설정
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	// 마지막 노드까지 순회
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	// 마지막 노드의 Next에 새 노드를 연결
	current.Next = newNode
}

// InsertAt 함수는 리스트의 특정 위치(0-indexed)에 새 노드를 삽입한다.
// position이 0이면 맨 앞에 삽입하고, 범위를 벗어나면 맨 뒤에 삽입한다.
func (ll *LinkedList) InsertAt(position int, data int) {
	// 위치가 0이면 맨 앞에 삽입
	if position <= 0 {
		ll.InsertHead(data)
		return
	}

	newNode := &Node{Data: data}

	// 삽입할 위치의 바로 앞 노드까지 이동
	current := ll.Head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.Next
	}

	// 범위를 벗어나면 맨 뒤에 삽입
	if current == nil {
		ll.InsertTail(data)
		return
	}

	// 새 노드를 현재 노드와 다음 노드 사이에 삽입
	newNode.Next = current.Next
	current.Next = newNode
}

// Delete 함수는 리스트에서 주어진 값을 가진 첫 번째 노드를 삭제한다.
// 삭제 성공 시 true, 값을 찾지 못하면 false를 반환한다.
func (ll *LinkedList) Delete(data int) bool {
	// 리스트가 비어있는 경우
	if ll.Head == nil {
		return false
	}

	// 삭제할 노드가 Head인 경우: Head를 다음 노드로 갱신
	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		return true
	}

	// 삭제할 노드를 찾기 위해 순회 (이전 노드를 추적)
	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			// 이전 노드의 Next를 삭제할 노드의 다음 노드로 연결
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}

	// 값을 찾지 못한 경우
	return false
}

// Search 함수는 리스트에서 주어진 값을 가진 노드의 위치(0-indexed)를 반환한다.
// 값을 찾지 못하면 -1을 반환한다.
func (ll *LinkedList) Search(data int) int {
	current := ll.Head
	index := 0

	// 리스트를 순회하며 값을 비교
	for current != nil {
		if current.Data == data {
			return index
		}
		current = current.Next
		index++
	}

	// 값을 찾지 못한 경우
	return -1
}

// ReverseList 함수는 단일 연결 리스트를 역순으로 뒤집는다.
// 세 개의 포인터(prev, current, next)를 사용하여 각 노드의 방향을 반전시킨다.
// 시간 복잡도: O(N), 공간 복잡도: O(1)
func ReverseList(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		// 다음 노드를 임시 저장
		next := current.Next
		// 현재 노드의 방향을 반전 (다음 → 이전)
		current.Next = prev
		// 포인터를 한 칸씩 전진
		prev = current
		current = next
	}

	// prev가 새로운 Head가 된다
	return prev
}

// PrintList 함수는 리스트의 모든 노드를 순서대로 출력한다.
// 각 노드의 값을 화살표(->)로 연결하여 표시한다.
func (ll *LinkedList) PrintList() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d", current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	ll := &LinkedList{}

	// --- 삽입 연산 예시 ---
	fmt.Println("=== 연결 리스트 기본 연산 ===")
	fmt.Println()

	// 맨 뒤에 삽입
	fmt.Println("--- 맨 뒤 삽입 (InsertTail) ---")
	ll.InsertTail(10)
	ll.InsertTail(20)
	ll.InsertTail(30)
	fmt.Print("리스트: ")
	ll.PrintList() // 10 -> 20 -> 30

	// 맨 앞에 삽입
	fmt.Println("\n--- 맨 앞 삽입 (InsertHead) ---")
	ll.InsertHead(5)
	fmt.Print("리스트: ")
	ll.PrintList() // 5 -> 10 -> 20 -> 30

	// 특정 위치에 삽입
	fmt.Println("\n--- 특정 위치 삽입 (InsertAt) ---")
	ll.InsertAt(2, 15) // 인덱스 2에 15 삽입
	fmt.Print("리스트: ")
	ll.PrintList() // 5 -> 10 -> 15 -> 20 -> 30

	// --- 탐색 연산 예시 ---
	fmt.Println("\n--- 탐색 (Search) ---")
	idx := ll.Search(15)
	fmt.Printf("값 15의 위치: %d\n", idx) // 2

	idx = ll.Search(99)
	fmt.Printf("값 99의 위치: %d (찾지 못함)\n", idx) // -1

	// --- 삭제 연산 예시 ---
	fmt.Println("\n--- 삭제 (Delete) ---")
	fmt.Printf("값 15 삭제: %v\n", ll.Delete(15))
	fmt.Print("리스트: ")
	ll.PrintList() // 5 -> 10 -> 20 -> 30

	fmt.Printf("값 5 삭제 (Head): %v\n", ll.Delete(5))
	fmt.Print("리스트: ")
	ll.PrintList() // 10 -> 20 -> 30

	fmt.Printf("값 99 삭제 (없는 값): %v\n", ll.Delete(99))
	fmt.Print("리스트: ")
	ll.PrintList() // 10 -> 20 -> 30

	// --- 역순 연산 예시 ---
	fmt.Println("\n--- 역순 (ReverseList) ---")
	fmt.Print("역순 전: ")
	ll.PrintList() // 10 -> 20 -> 30

	ll.Head = ReverseList(ll.Head)
	fmt.Print("역순 후: ")
	ll.PrintList() // 30 -> 20 -> 10

	// 빈 리스트 역순 테스트
	emptyLL := &LinkedList{}
	emptyLL.Head = ReverseList(emptyLL.Head)
	fmt.Print("빈 리스트 역순: ")
	emptyLL.PrintList()

	// 단일 노드 리스트 역순 테스트
	singleLL := &LinkedList{}
	singleLL.InsertHead(42)
	singleLL.Head = ReverseList(singleLL.Head)
	fmt.Print("단일 노드 역순: ")
	singleLL.PrintList() // 42
}
