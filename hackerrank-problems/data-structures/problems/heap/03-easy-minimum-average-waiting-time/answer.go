package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

// Customer는 고객 정보를 나타낸다.
type Customer struct {
	Arrive int // 도착 시간
	Cook   int // 조리 시간
}

// CookHeap은 조리 시간 기준 최소 힙이다.
type CookHeap []Customer

func (h CookHeap) Len() int           { return len(h) }
func (h CookHeap) Less(i, j int) bool { return h[i].Cook < h[j].Cook }
func (h CookHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CookHeap) Push(x interface{}) {
	*h = append(*h, x.(Customer))
}

func (h *CookHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// minimumAverage는 최소 평균 대기 시간을 반환한다.
//
// [매개변수]
//   - customers: 고객 목록 (도착 시간, 조리 시간)
//
// [반환값]
//   - int64: 최소 평균 대기 시간 (소수점 이하 버림)
//
// [알고리즘 힌트]
//
//	SJF(Shortest Job First) 스케줄링을 적용한다.
//	도착 시간 순으로 정렬 후, 현재 시점에서 도착한 고객 중
//	조리 시간이 가장 짧은 고객을 먼저 처리한다.
func minimumAverage(customers []Customer) int64 {
	// 도착 시간 기준으로 정렬
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].Arrive < customers[j].Arrive
	})

	h := &CookHeap{}
	heap.Init(h)

	var totalWait int64
	currentTime := 0
	idx := 0
	n := len(customers)

	for idx < n || h.Len() > 0 {
		// 현재 시간까지 도착한 고객을 힙에 추가
		for idx < n && customers[idx].Arrive <= currentTime {
			heap.Push(h, customers[idx])
			idx++
		}

		if h.Len() == 0 {
			// 힙이 비어있으면 다음 고객의 도착 시간으로 점프
			currentTime = customers[idx].Arrive
			continue
		}

		// 조리 시간이 가장 짧은 고객을 처리
		c := heap.Pop(h).(Customer)
		currentTime += c.Cook
		// 대기 시간 = 완료 시간 - 도착 시간
		totalWait += int64(currentTime - c.Arrive)
	}

	// 평균 대기 시간 (소수점 이하 버림)
	return totalWait / int64(n)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 고객 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 고객 정보 입력
	customers := make([]Customer, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &customers[i].Arrive, &customers[i].Cook)
	}

	// 핵심 함수 호출
	result := minimumAverage(customers)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
