package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// MaxHeap은 최대 힙을 구현한다 (작은 절반 저장).
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeap은 최소 힙을 구현한다 (큰 절반 저장).
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// runningMedian은 각 원소 삽입 후의 중앙값을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - []float64: 각 단계의 중앙값
//
// [알고리즘 힌트]
//
//	두 개의 힙(최대 힙 + 최소 힙)으로 수열을 반으로 나눈다.
//	최대 힙에는 작은 절반, 최소 힙에는 큰 절반을 저장한다.
//	두 힙의 크기 차이가 1 이하가 되도록 균형을 유지한다.
func runningMedian(arr []int) []float64 {
	// 최대 힙: 작은 절반 (루트가 최댓값)
	maxH := &MaxHeap{}
	// 최소 힙: 큰 절반 (루트가 최솟값)
	minH := &MinHeap{}
	heap.Init(maxH)
	heap.Init(minH)

	results := make([]float64, 0, len(arr))

	for _, num := range arr {
		// 새 원소를 적절한 힙에 삽입
		if maxH.Len() == 0 || num <= (*maxH)[0] {
			heap.Push(maxH, num)
		} else {
			heap.Push(minH, num)
		}

		// 두 힙의 크기 균형 맞추기 (차이가 1 이하)
		if maxH.Len() > minH.Len()+1 {
			heap.Push(minH, heap.Pop(maxH))
		} else if minH.Len() > maxH.Len() {
			heap.Push(maxH, heap.Pop(minH))
		}

		// 중앙값 계산
		if maxH.Len() == minH.Len() {
			// 짝수 개: 두 힙의 루트 평균
			median := float64((*maxH)[0]+(*minH)[0]) / 2.0
			results = append(results, median)
		} else {
			// 홀수 개: 더 큰 힙의 루트
			results = append(results, float64((*maxH)[0]))
		}
	}

	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 정수 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 정수 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	results := runningMedian(arr)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintf(writer, "%.1f\n", r)
	}
}
