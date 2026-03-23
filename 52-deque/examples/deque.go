package main

import "fmt"

// 덱(Deque) - 양쪽 끝에서 삽입/삭제가 가능한 자료구조
// 시간 복잡도: 모든 연산 O(1) (슬라이스 기반 PushFront 제외)
// 공간 복잡도: O(N)

// Deque는 슬라이스 기반 덱 구현이다
type Deque struct {
	data []int
}

// NewDeque는 빈 덱을 생성한다
func NewDeque() *Deque {
	return &Deque{}
}

// PushFront는 덱의 앞에 원소를 추가한다
func (d *Deque) PushFront(val int) {
	d.data = append([]int{val}, d.data...)
}

// PushBack은 덱의 뒤에 원소를 추가한다
func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

// PopFront는 덱의 앞 원소를 제거하고 반환한다
func (d *Deque) PopFront() int {
	val := d.data[0]
	d.data = d.data[1:]
	return val
}

// PopBack은 덱의 뒤 원소를 제거하고 반환한다
func (d *Deque) PopBack() int {
	val := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return val
}

// Front는 덱의 앞 원소를 확인한다
func (d *Deque) Front() int {
	return d.data[0]
}

// Back은 덱의 뒤 원소를 확인한다
func (d *Deque) Back() int {
	return d.data[len(d.data)-1]
}

// IsEmpty는 덱이 비어 있는지 확인한다
func (d *Deque) IsEmpty() bool {
	return len(d.data) == 0
}

// Size는 덱의 크기를 반환한다
func (d *Deque) Size() int {
	return len(d.data)
}

// slidingWindowMax는 모노톤 덱을 이용한 슬라이딩 윈도우 최댓값을 구한다
// 크기 K인 윈도우를 이동하며 각 위치에서의 최댓값을 반환한다
func slidingWindowMax(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k == 0 {
		return nil
	}

	var result []int
	// 덱에는 인덱스를 저장한다 (앞쪽이 항상 최댓값의 인덱스)
	deque := []int{}

	for i := 0; i < n; i++ {
		// 윈도우 범위를 벗어난 인덱스를 앞에서 제거한다
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 현재 원소보다 작은 원소의 인덱스를 뒤에서 제거한다
		for len(deque) > 0 && arr[deque[len(deque)-1]] <= arr[i] {
			deque = deque[:len(deque)-1]
		}

		// 현재 인덱스를 덱의 뒤에 추가한다
		deque = append(deque, i)

		// 윈도우가 완성된 시점부터 최댓값을 기록한다
		if i >= k-1 {
			result = append(result, arr[deque[0]])
		}
	}

	return result
}

func main() {
	// 기본 덱 연산 예시
	dq := NewDeque()

	fmt.Println("=== 기본 덱 연산 ===")
	dq.PushBack(1)
	dq.PushBack(2)
	dq.PushBack(3)
	fmt.Printf("PushBack 1, 2, 3 후 Front: %d, Back: %d\n", dq.Front(), dq.Back())

	dq.PushFront(0)
	fmt.Printf("PushFront 0 후 Front: %d, Back: %d\n", dq.Front(), dq.Back())

	val := dq.PopFront()
	fmt.Printf("PopFront: %d, 현재 Front: %d\n", val, dq.Front())

	val = dq.PopBack()
	fmt.Printf("PopBack: %d, 현재 Back: %d\n", val, dq.Back())

	fmt.Printf("덱 크기: %d, 비어있는가: %v\n", dq.Size(), dq.IsEmpty())

	// 모노톤 덱을 이용한 슬라이딩 윈도우 최댓값 예시
	fmt.Println("\n=== 슬라이딩 윈도우 최댓값 (모노톤 덱) ===")
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Printf("배열: %v, 윈도우 크기: %d\n", arr, k)

	maxValues := slidingWindowMax(arr, k)
	fmt.Printf("각 윈도우의 최댓값: %v\n", maxValues)
	// 윈도우 [1,3,-1]=3, [3,-1,-3]=3, [-1,-3,5]=5, [-3,5,3]=5, [5,3,6]=6, [3,6,7]=7
}
