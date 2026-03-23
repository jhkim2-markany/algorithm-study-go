package main

import "fmt"

// 큐(Queue) - 선입선출(FIFO) 자료구조 기본 구현
// 슬라이스를 활용하여 Enqueue, Dequeue, Front 연산을 구현한다.
// 시간 복잡도: 모든 연산 O(1) amortized
// 공간 복잡도: O(N)

// Queue 구조체 정의
type Queue struct {
	data []int
}

// Enqueue - 큐의 뒤에 원소를 추가한다
func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

// Dequeue - 큐의 앞 원소를 제거하고 반환한다
func (q *Queue) Dequeue() int {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

// Front - 큐의 앞 원소를 제거하지 않고 확인한다
func (q *Queue) Front() int {
	return q.data[0]
}

// IsEmpty - 큐가 비어 있는지 확인한다
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Size - 큐에 저장된 원소의 개수를 반환한다
func (q *Queue) Size() int {
	return len(q.data)
}

func main() {
	// 큐 생성 및 기본 연산 테스트
	q := &Queue{}

	fmt.Println("=== 큐 기본 연산 ===")

	// Enqueue 연산: 원소 삽입
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Printf("Enqueue 10, 20, 30 후 크기: %d\n", q.Size())
	fmt.Printf("Front (맨 앞 원소): %d\n", q.Front())

	// Dequeue 연산: 원소 제거 (FIFO 순서로 나온다)
	fmt.Printf("Dequeue: %d\n", q.Dequeue())
	fmt.Printf("Dequeue: %d\n", q.Dequeue())
	fmt.Printf("Dequeue 후 크기: %d\n", q.Size())
	fmt.Printf("비어 있는가: %v\n", q.IsEmpty())

	// 마지막 원소 Dequeue
	fmt.Printf("Dequeue: %d\n", q.Dequeue())
	fmt.Printf("비어 있는가: %v\n", q.IsEmpty())

	// 큐 활용 예시: 작업 대기열 시뮬레이션
	fmt.Println("\n=== 큐 활용: 작업 대기열 시뮬레이션 ===")
	taskQueue := &Queue{}
	tasks := []string{"작업A", "작업B", "작업C", "작업D"}

	// 작업을 큐에 등록
	for i, task := range tasks {
		taskQueue.Enqueue(i)
		fmt.Printf("등록: %s\n", task)
	}

	// 작업을 순서대로 처리
	fmt.Println("--- 처리 시작 ---")
	for !taskQueue.IsEmpty() {
		idx := taskQueue.Dequeue()
		fmt.Printf("처리 완료: %s\n", tasks[idx])
	}
}
