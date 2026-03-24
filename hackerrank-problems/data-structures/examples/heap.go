package main

import "fmt"

// 최소 힙 (Min-Heap) - 기본 연산 예시
// 배열 기반 최소 힙의 삽입(push), 추출(pop), 최솟값 확인(peek) 연산을 구현한다.
// container/heap 패키지를 사용하지 않고 직접 구현하여 내부 동작 원리를 보여준다.
//
// 시간 복잡도:
//   - 삽입 (Push):    O(log N) (N: 힙의 원소 수, 삽입 후 위로 올리기)
//   - 추출 (Pop):     O(log N) (루트 제거 후 아래로 내리기)
//   - 최솟값 (Peek):  O(1)     (루트 노드 반환)
//   - 힙 생성:        O(N log N) (N개 원소를 하나씩 삽입하는 경우)
//
// 공간 복잡도: O(N) (N개의 원소를 배열에 저장)
//
// 힙 속성 (Heap Property):
//   부모 노드의 값은 항상 자식 노드의 값보다 작거나 같다.
//   배열 인덱스 관계:
//     - 부모: (i - 1) / 2
//     - 왼쪽 자식: 2*i + 1
//     - 오른쪽 자식: 2*i + 2

// MinHeap은 배열 기반 최소 힙을 나타낸다.
// data 슬라이스에 힙의 원소를 저장한다.
type MinHeap struct {
	data []int
}

// NewMinHeap 함수는 빈 최소 힙을 생성하여 반환한다.
func NewMinHeap() *MinHeap {
	return &MinHeap{data: []int{}}
}

// Size 함수는 힙에 저장된 원소의 개수를 반환한다.
func (h *MinHeap) Size() int {
	return len(h.data)
}

// IsEmpty 함수는 힙이 비어있는지 확인한다.
func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// Peek 함수는 힙의 최솟값(루트)을 제거하지 않고 반환한다.
// 힙이 비어있으면 -1과 false를 반환한다.
func (h *MinHeap) Peek() (int, bool) {
	if h.IsEmpty() {
		return -1, false
	}
	// 루트 노드(인덱스 0)가 항상 최솟값
	return h.data[0], true
}

// Push 함수는 힙에 새 원소를 삽입한다.
// 1. 배열의 맨 끝에 원소를 추가한다.
// 2. 힙 속성을 만족할 때까지 부모와 비교하며 위로 올린다 (sift up).
func (h *MinHeap) Push(val int) {
	// 배열 맨 끝에 새 원소 추가
	h.data = append(h.data, val)
	// 삽입된 원소를 올바른 위치로 올리기
	h.siftUp(len(h.data) - 1)
}

// siftUp 함수는 주어진 인덱스의 원소를 부모와 비교하며 위로 올린다.
// 부모보다 작으면 교환하고, 루트에 도달하거나 부모보다 크면 멈춘다.
func (h *MinHeap) siftUp(index int) {
	for index > 0 {
		// 부모 인덱스 계산
		parent := (index - 1) / 2

		// 부모보다 작으면 교환
		if h.data[index] < h.data[parent] {
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
			index = parent
		} else {
			// 힙 속성을 만족하면 종료
			break
		}
	}
}

// Pop 함수는 힙에서 최솟값(루트)을 제거하고 반환한다.
// 1. 루트와 마지막 원소를 교환한다.
// 2. 마지막 원소를 제거한다.
// 3. 새 루트를 올바른 위치로 내린다 (sift down).
// 힙이 비어있으면 -1과 false를 반환한다.
func (h *MinHeap) Pop() (int, bool) {
	if h.IsEmpty() {
		return -1, false
	}

	// 루트(최솟값) 저장
	min := h.data[0]
	lastIdx := len(h.data) - 1

	// 루트와 마지막 원소 교환
	h.data[0] = h.data[lastIdx]
	// 마지막 원소 제거
	h.data = h.data[:lastIdx]

	// 힙이 비어있지 않으면 새 루트를 아래로 내리기
	if !h.IsEmpty() {
		h.siftDown(0)
	}

	return min, true
}

// siftDown 함수는 주어진 인덱스의 원소를 자식과 비교하며 아래로 내린다.
// 더 작은 자식과 교환하고, 리프에 도달하거나 자식보다 작으면 멈춘다.
func (h *MinHeap) siftDown(index int) {
	size := len(h.data)

	for {
		smallest := index
		left := 2*index + 1  // 왼쪽 자식 인덱스
		right := 2*index + 2 // 오른쪽 자식 인덱스

		// 왼쪽 자식이 현재보다 작으면 교환 대상으로 설정
		if left < size && h.data[left] < h.data[smallest] {
			smallest = left
		}

		// 오른쪽 자식이 현재 최솟값보다 작으면 교환 대상으로 갱신
		if right < size && h.data[right] < h.data[smallest] {
			smallest = right
		}

		// 현재 노드가 가장 작으면 힙 속성 만족 → 종료
		if smallest == index {
			break
		}

		// 더 작은 자식과 교환 후 해당 위치에서 계속 내리기
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

// PrintHeap 함수는 힙의 내부 배열 상태를 출력한다.
func (h *MinHeap) PrintHeap() {
	fmt.Printf("힙 배열: %v (크기: %d)\n", h.data, h.Size())
}

func main() {
	heap := NewMinHeap()

	// --- 삽입 (Push) 예시 ---
	fmt.Println("=== 최소 힙 기본 연산 ===")
	fmt.Println()

	fmt.Println("--- 삽입 (Push) ---")
	values := []int{35, 20, 50, 10, 25, 45, 15}
	for _, v := range values {
		heap.Push(v)
		fmt.Printf("Push(%d) → ", v)
		heap.PrintHeap()
	}

	// --- 최솟값 확인 (Peek) 예시 ---
	fmt.Println("\n--- 최솟값 확인 (Peek) ---")
	if val, ok := heap.Peek(); ok {
		fmt.Printf("최솟값: %d (제거하지 않음)\n", val)
	}
	heap.PrintHeap()

	// --- 추출 (Pop) 예시 ---
	fmt.Println("\n--- 추출 (Pop) ---")
	fmt.Println("최솟값부터 순서대로 추출:")
	for !heap.IsEmpty() {
		val, _ := heap.Pop()
		fmt.Printf("Pop() → %d  |  ", val)
		heap.PrintHeap()
	}

	// --- 빈 힙에서 연산 ---
	fmt.Println("\n--- 빈 힙 테스트 ---")
	if _, ok := heap.Peek(); !ok {
		fmt.Println("Peek: 힙이 비어있음")
	}
	if _, ok := heap.Pop(); !ok {
		fmt.Println("Pop: 힙이 비어있음")
	}

	// --- 힙 정렬 시연 ---
	fmt.Println("\n--- 힙 정렬 (Heap Sort) 시연 ---")
	unsorted := []int{64, 25, 12, 22, 11, 90, 33}
	fmt.Printf("정렬 전: %v\n", unsorted)

	sortHeap := NewMinHeap()
	for _, v := range unsorted {
		sortHeap.Push(v)
	}

	sorted := []int{}
	for !sortHeap.IsEmpty() {
		val, _ := sortHeap.Pop()
		sorted = append(sorted, val)
	}
	fmt.Printf("정렬 후: %v\n", sorted)
}
