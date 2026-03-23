package main

import (
	"container/heap"
	"fmt"
)

// 힙과 우선순위 큐 기본 예시
// 1) 최소 힙: container/heap 인터페이스 구현
// 2) 최대 힙: Less 함수를 반대로 구현
// 3) 우선순위 큐: 구조체에 우선순위를 부여하여 관리

// === 최소 힙 구현 ===

// MinHeap은 정수 최소 힙이다.
// container/heap.Interface를 구현한다.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // 작은 값이 루트
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push는 힙에 원소를 추가한다. heap.Push에서 호출된다.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop은 힙에서 마지막 원소를 제거하고 반환한다. heap.Pop에서 호출된다.
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// === 최대 힙 구현 ===

// MaxHeap은 정수 최대 힙이다.
// Less 함수에서 부등호 방향을 반대로 하면 최대 힙이 된다.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 큰 값이 루트
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// === 우선순위 큐 구현 ===

// Task는 이름과 우선순위를 가진 작업이다.
type Task struct {
	name     string
	priority int // 값이 작을수록 높은 우선순위
}

// TaskQueue는 Task의 우선순위 큐이다.
type TaskQueue []Task

func (q TaskQueue) Len() int           { return len(q) }
func (q TaskQueue) Less(i, j int) bool { return q[i].priority < q[j].priority }
func (q TaskQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *TaskQueue) Push(x interface{}) {
	*q = append(*q, x.(Task))
}

func (q *TaskQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func main() {
	// === 최소 힙 예시 ===
	fmt.Println("=== 최소 힙 ===")
	minH := &MinHeap{}
	heap.Init(minH)

	// 원소 삽입
	values := []int{5, 3, 8, 1, 7, 2}
	for _, v := range values {
		heap.Push(minH, v)
	}
	fmt.Printf("삽입된 값: %v\n", values)

	// 최솟값부터 순서대로 꺼내기
	fmt.Print("최소 힙에서 꺼낸 순서: ")
	for minH.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(minH))
	}
	fmt.Println()

	// === 최대 힙 예시 ===
	fmt.Println("\n=== 최대 힙 ===")
	maxH := &MaxHeap{}
	heap.Init(maxH)

	for _, v := range values {
		heap.Push(maxH, v)
	}
	fmt.Printf("삽입된 값: %v\n", values)

	// 최댓값부터 순서대로 꺼내기
	fmt.Print("최대 힙에서 꺼낸 순서: ")
	for maxH.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(maxH))
	}
	fmt.Println()

	// === 우선순위 큐 예시 ===
	fmt.Println("\n=== 우선순위 큐 ===")
	pq := &TaskQueue{}
	heap.Init(pq)

	// 작업 추가 (우선순위 값이 작을수록 먼저 처리)
	tasks := []Task{
		{"이메일 확인", 3},
		{"버그 수정", 1},
		{"회의 참석", 2},
		{"코드 리뷰", 1},
		{"문서 작성", 4},
	}
	for _, t := range tasks {
		heap.Push(pq, t)
		fmt.Printf("  추가: %s (우선순위: %d)\n", t.name, t.priority)
	}

	// 우선순위 순서대로 작업 처리
	fmt.Println("\n처리 순서:")
	for pq.Len() > 0 {
		t := heap.Pop(pq).(Task)
		fmt.Printf("  처리: %s (우선순위: %d)\n", t.name, t.priority)
	}
}
