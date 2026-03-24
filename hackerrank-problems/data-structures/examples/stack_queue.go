package main

import "fmt"

// 스택, 큐, 덱 (Stack, Queue, Deque) - 기본 연산 예시
// 배열 기반 스택/큐와 이중 연결 리스트 기반 덱의 기본 연산을 구현한다.
//
// 시간 복잡도:
//   [스택 - 배열 기반]
//   - Push:    O(1) (amortized, 슬라이스 append)
//   - Pop:     O(1)
//   - Peek:    O(1)
//   - IsEmpty: O(1)
//
//   [큐 - 배열 기반]
//   - Enqueue: O(1) (amortized, 슬라이스 append)
//   - Dequeue: O(N) (맨 앞 제거 시 슬라이스 이동, 링 버퍼로 O(1) 가능)
//   - Front:   O(1)
//   - IsEmpty: O(1)
//
//   [덱 - 이중 연결 리스트 기반]
//   - PushFront:  O(1)
//   - PushBack:   O(1)
//   - PopFront:   O(1)
//   - PopBack:    O(1)
//
// 공간 복잡도: 모두 O(N) (N개의 원소를 저장)

// ============================================================
// 스택 (Stack) - 배열(슬라이스) 기반
// LIFO(Last In First Out) 구조
// ============================================================

// Stack은 배열 기반 스택을 나타낸다.
type Stack struct {
	data []int
}

// NewStack 함수는 빈 스택을 생성하여 반환한다.
func NewStack() *Stack {
	return &Stack{data: []int{}}
}

// Push 함수는 스택의 맨 위에 원소를 추가한다.
func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

// Pop 함수는 스택의 맨 위 원소를 제거하고 반환한다.
// 스택이 비어있으면 -1과 false를 반환한다.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

// Peek 함수는 스택의 맨 위 원소를 제거하지 않고 반환한다.
// 스택이 비어있으면 -1과 false를 반환한다.
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	return s.data[len(s.data)-1], true
}

// IsEmpty 함수는 스택이 비어있는지 확인한다.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// ============================================================
// 큐 (Queue) - 배열(슬라이스) 기반
// FIFO(First In First Out) 구조
// ============================================================

// Queue는 배열 기반 큐를 나타낸다.
type Queue struct {
	data []int
}

// NewQueue 함수는 빈 큐를 생성하여 반환한다.
func NewQueue() *Queue {
	return &Queue{data: []int{}}
}

// Enqueue 함수는 큐의 맨 뒤에 원소를 추가한다.
func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

// Dequeue 함수는 큐의 맨 앞 원소를 제거하고 반환한다.
// 큐가 비어있으면 -1과 false를 반환한다.
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	front := q.data[0]
	q.data = q.data[1:]
	return front, true
}

// Front 함수는 큐의 맨 앞 원소를 제거하지 않고 반환한다.
// 큐가 비어있으면 -1과 false를 반환한다.
func (q *Queue) Front() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	return q.data[0], true
}

// IsEmpty 함수는 큐가 비어있는지 확인한다.
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// ============================================================
// 덱 (Deque) - 이중 연결 리스트 기반
// 양쪽 끝에서 삽입/삭제가 가능한 구조
// ============================================================

// DequeNode는 이중 연결 리스트의 노드를 나타낸다.
type DequeNode struct {
	Data int
	Prev *DequeNode
	Next *DequeNode
}

// Deque는 이중 연결 리스트 기반 덱을 나타낸다.
// Front와 Back은 각각 덱의 앞과 뒤를 가리킨다.
type Deque struct {
	front *DequeNode
	back  *DequeNode
	size  int
}

// NewDeque 함수는 빈 덱을 생성하여 반환한다.
func NewDeque() *Deque {
	return &Deque{}
}

// PushFront 함수는 덱의 앞에 원소를 추가한다.
func (d *Deque) PushFront(val int) {
	node := &DequeNode{Data: val}
	if d.front == nil {
		// 덱이 비어있으면 front와 back 모두 새 노드를 가리킴
		d.front = node
		d.back = node
	} else {
		// 기존 front 앞에 새 노드를 연결
		node.Next = d.front
		d.front.Prev = node
		d.front = node
	}
	d.size++
}

// PushBack 함수는 덱의 뒤에 원소를 추가한다.
func (d *Deque) PushBack(val int) {
	node := &DequeNode{Data: val}
	if d.back == nil {
		// 덱이 비어있으면 front와 back 모두 새 노드를 가리킴
		d.front = node
		d.back = node
	} else {
		// 기존 back 뒤에 새 노드를 연결
		node.Prev = d.back
		d.back.Next = node
		d.back = node
	}
	d.size++
}

// PopFront 함수는 덱의 앞 원소를 제거하고 반환한다.
// 덱이 비어있으면 -1과 false를 반환한다.
func (d *Deque) PopFront() (int, bool) {
	if d.front == nil {
		return -1, false
	}
	val := d.front.Data
	d.front = d.front.Next
	if d.front != nil {
		d.front.Prev = nil
	} else {
		// 마지막 원소를 제거한 경우 back도 nil로 설정
		d.back = nil
	}
	d.size--
	return val, true
}

// PopBack 함수는 덱의 뒤 원소를 제거하고 반환한다.
// 덱이 비어있으면 -1과 false를 반환한다.
func (d *Deque) PopBack() (int, bool) {
	if d.back == nil {
		return -1, false
	}
	val := d.back.Data
	d.back = d.back.Prev
	if d.back != nil {
		d.back.Next = nil
	} else {
		// 마지막 원소를 제거한 경우 front도 nil로 설정
		d.front = nil
	}
	d.size--
	return val, true
}

// IsEmpty 함수는 덱이 비어있는지 확인한다.
func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

// PrintDeque 함수는 덱의 모든 원소를 앞에서 뒤 순서로 출력한다.
func (d *Deque) PrintDeque() {
	current := d.front
	fmt.Print("[")
	for current != nil {
		fmt.Printf("%d", current.Data)
		if current.Next != nil {
			fmt.Print(", ")
		}
		current = current.Next
	}
	fmt.Println("]")
}

func main() {
	// ============================================================
	// 스택 사용 예시
	// ============================================================
	fmt.Println("=== 스택 (Stack) - LIFO ===")
	fmt.Println()

	st := NewStack()

	fmt.Println("--- Push ---")
	st.Push(10)
	st.Push(20)
	st.Push(30)
	fmt.Printf("Push(10), Push(20), Push(30) → 스택: %v\n", st.data)

	fmt.Println("\n--- Peek ---")
	if val, ok := st.Peek(); ok {
		fmt.Printf("Peek() → %d (제거하지 않음)\n", val)
	}

	fmt.Println("\n--- Pop ---")
	for !st.IsEmpty() {
		val, _ := st.Pop()
		fmt.Printf("Pop() → %d  |  스택: %v\n", val, st.data)
	}

	fmt.Println("\n--- 빈 스택 테스트 ---")
	if _, ok := st.Pop(); !ok {
		fmt.Println("Pop: 스택이 비어있음")
	}
	if _, ok := st.Peek(); !ok {
		fmt.Println("Peek: 스택이 비어있음")
	}

	// ============================================================
	// 큐 사용 예시
	// ============================================================
	fmt.Println("\n=== 큐 (Queue) - FIFO ===")
	fmt.Println()

	q := NewQueue()

	fmt.Println("--- Enqueue ---")
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Printf("Enqueue(10), Enqueue(20), Enqueue(30) → 큐: %v\n", q.data)

	fmt.Println("\n--- Front ---")
	if val, ok := q.Front(); ok {
		fmt.Printf("Front() → %d (제거하지 않음)\n", val)
	}

	fmt.Println("\n--- Dequeue ---")
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		fmt.Printf("Dequeue() → %d  |  큐: %v\n", val, q.data)
	}

	fmt.Println("\n--- 빈 큐 테스트 ---")
	if _, ok := q.Dequeue(); !ok {
		fmt.Println("Dequeue: 큐가 비어있음")
	}
	if _, ok := q.Front(); !ok {
		fmt.Println("Front: 큐가 비어있음")
	}

	// ============================================================
	// 덱 사용 예시
	// ============================================================
	fmt.Println("\n=== 덱 (Deque) - 이중 연결 리스트 기반 ===")
	fmt.Println()

	dq := NewDeque()

	fmt.Println("--- PushBack / PushFront ---")
	dq.PushBack(10)
	dq.PushBack(20)
	dq.PushFront(5)
	dq.PushBack(30)
	dq.PushFront(1)
	fmt.Print("PushBack(10), PushBack(20), PushFront(5), PushBack(30), PushFront(1) → 덱: ")
	dq.PrintDeque() // [1, 5, 10, 20, 30]

	fmt.Println("\n--- PopFront ---")
	if val, ok := dq.PopFront(); ok {
		fmt.Printf("PopFront() → %d  |  덱: ", val)
		dq.PrintDeque() // [5, 10, 20, 30]
	}

	fmt.Println("\n--- PopBack ---")
	if val, ok := dq.PopBack(); ok {
		fmt.Printf("PopBack() → %d  |  덱: ", val)
		dq.PrintDeque() // [5, 10, 20]
	}

	fmt.Println("\n--- 모든 원소 제거 ---")
	for !dq.IsEmpty() {
		val, _ := dq.PopFront()
		fmt.Printf("PopFront() → %d  |  덱: ", val)
		dq.PrintDeque()
	}

	fmt.Println("\n--- 빈 덱 테스트 ---")
	if _, ok := dq.PopFront(); !ok {
		fmt.Println("PopFront: 덱이 비어있음")
	}
	if _, ok := dq.PopBack(); !ok {
		fmt.Println("PopBack: 덱이 비어있음")
	}
}
