package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// MaxH는 최대 힙을 구현한다 (작은 절반 저장).
type MaxH []int

func (h MaxH) Len() int           { return len(h) }
func (h MaxH) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxH) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinH는 최소 힙을 구현한다 (큰 절반 저장).
type MinH []int

func (h MinH) Len() int           { return len(h) }
func (h MinH) Less(i, j int) bool { return h[i] < h[j] }
func (h MinH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinH) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// medianUpdates는 각 연산 후의 중앙값을 반환한다.
//
// [매개변수]
//   - ops: 연산 목록 (각 연산은 [타입, 값])
//
// [반환값]
//   - []string: 각 연산 후의 중앙값 문자열
//
// [알고리즘 힌트]
//
//	두 개의 힙(최대 힙 + 최소 힙)과 lazy deletion을 조합한다.
//	삭제할 원소를 맵에 기록하고, 힙의 루트가 삭제된 원소이면 제거한다.
//	유효한 크기를 별도로 관리하여 균형을 맞춘다.
func medianUpdates(ops [][]interface{}) []string {
	maxH := &MaxH{}
	minH := &MinH{}
	heap.Init(maxH)
	heap.Init(minH)

	// 원소 개수 추적
	count := make(map[int]int)
	// 유효한 크기
	maxSize, minSize := 0, 0

	// 삭제된 원소를 힙 루트에서 제거
	cleanTop := func() {
		for maxH.Len() > 0 && count[(*maxH)[0]] == 0 {
			heap.Pop(maxH)
		}
		for minH.Len() > 0 && count[(*minH)[0]] == 0 {
			heap.Pop(minH)
		}
	}

	// 균형 맞추기
	balance := func() {
		for maxSize > minSize+1 {
			cleanTop()
			val := heap.Pop(maxH).(int)
			heap.Push(minH, val)
			maxSize--
			minSize++
		}
		for minSize > maxSize {
			cleanTop()
			val := heap.Pop(minH).(int)
			heap.Push(maxH, val)
			minSize--
			maxSize++
		}
		cleanTop()
	}

	var results []string

	for _, op := range ops {
		opType := op[0].(string)
		x := op[1].(int)

		if opType == "a" {
			// 추가 연산
			count[x]++
			if maxH.Len() == 0 || x <= (*maxH)[0] {
				heap.Push(maxH, x)
				maxSize++
			} else {
				heap.Push(minH, x)
				minSize++
			}
			balance()
		} else {
			// 삭제 연산
			if count[x] == 0 {
				results = append(results, "Wrong!")
				continue
			}
			count[x]--
			if count[x] == 0 {
				delete(count, x)
			}
			// 어느 힙에서 삭제되었는지 판단
			if maxH.Len() > 0 && x <= (*maxH)[0] {
				maxSize--
			} else {
				minSize--
			}
			balance()
		}

		// 중앙값 계산
		total := maxSize + minSize
		if total == 0 {
			results = append(results, "Wrong!")
		} else if total%2 == 1 {
			// 홀수: 최대 힙의 루트
			results = append(results, fmt.Sprintf("%d", (*maxH)[0]))
		} else {
			// 짝수: 두 힙의 루트 평균
			sum := (*maxH)[0] + (*minH)[0]
			if sum%2 == 0 {
				results = append(results, fmt.Sprintf("%d", sum/2))
			} else {
				results = append(results, fmt.Sprintf("%.1f", float64(sum)/2.0))
			}
		}
	}

	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 연산 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 연산 입력
	ops := make([][]interface{}, n)
	for i := 0; i < n; i++ {
		var op string
		var x int
		fmt.Fscan(reader, &op, &x)
		ops[i] = []interface{}{op, x}
	}

	// 핵심 함수 호출
	results := medianUpdates(ops)

	// 결과 출력
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
