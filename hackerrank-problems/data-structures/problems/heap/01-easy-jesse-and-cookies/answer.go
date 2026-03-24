package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// IntMinHeap은 정수 최소 힙을 구현한다.
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// cookies는 모든 쿠키의 달콤함이 K 이상이 되기 위한 최소 연산 횟수를 반환한다.
//
// [매개변수]
//   - k: 목표 달콤함
//   - arr: 쿠키의 달콤함 배열
//
// [반환값]
//   - int: 최소 연산 횟수 (불가능하면 -1)
//
// [알고리즘 힌트]
//
//	최소 힙을 사용하여 항상 가장 작은 두 쿠키를 꺼내 합친다.
//	새 쿠키 = 가장 작은 쿠키 + 2 × 두 번째로 작은 쿠키
//	힙의 최솟값이 K 이상이 되면 종료한다.
func cookies(k int, arr []int) int {
	// 최소 힙 초기화
	h := &IntMinHeap{}
	for _, v := range arr {
		*h = append(*h, v)
	}
	heap.Init(h)

	// 연산 횟수 카운터
	count := 0

	// 최솟값이 K 미만인 동안 반복
	for h.Len() >= 2 && (*h)[0] < k {
		// 가장 작은 쿠키 두 개를 꺼냄
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)

		// 새 쿠키 생성: 가장 작은 + 2 × 두 번째로 작은
		newCookie := first + 2*second
		heap.Push(h, newCookie)

		// 연산 횟수 증가
		count++
	}

	// 모든 쿠키가 K 이상인지 확인
	if h.Len() > 0 && (*h)[0] < k {
		return -1
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 쿠키 개수 N, 목표 달콤함 K
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 쿠키 달콤함 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := cookies(k, arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
