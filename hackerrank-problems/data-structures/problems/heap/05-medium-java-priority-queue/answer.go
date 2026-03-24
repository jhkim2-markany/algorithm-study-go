package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Student는 학생 정보를 나타낸다.
type Student struct {
	name string
	cgpa float64
	id   int
}

// StudentHeap은 학생 우선순위 큐를 구현한다.
type StudentHeap []Student

func (h StudentHeap) Len() int { return len(h) }
func (h StudentHeap) Less(i, j int) bool {
	if h[i].cgpa != h[j].cgpa {
		return h[i].cgpa > h[j].cgpa // CGPA 내림차순
	}
	if h[i].name != h[j].name {
		return h[i].name < h[j].name // 이름 오름차순
	}
	return h[i].id < h[j].id // ID 오름차순
}
func (h StudentHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StudentHeap) Push(x interface{}) {
	*h = append(*h, x.(Student))
}

func (h *StudentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// processEvents는 모든 이벤트를 처리한 후 큐에 남은 학생 이름 목록을 반환한다.
//
// [매개변수]
//   - events: 이벤트 문자열 목록
//
// [반환값]
//   - []string: 남은 학생 이름 목록 (우선순위 순서), 비어있으면 ["EMPTY"]
//
// [알고리즘 힌트]
//
//	커스텀 비교 함수를 가진 힙을 사용한다.
//	ENTER 이벤트에서 학생을 힙에 삽입하고,
//	SERVED 이벤트에서 최우선순위 학생을 제거한다.
//	최종적으로 힙에서 하나씩 꺼내 결과를 구성한다.
func processEvents(events []string) []string {
	h := &StudentHeap{}
	heap.Init(h)

	for _, event := range events {
		parts := strings.Fields(event)
		if parts[0] == "ENTER" {
			name := parts[1]
			cgpa, _ := strconv.ParseFloat(parts[2], 64)
			id, _ := strconv.Atoi(parts[3])
			heap.Push(h, Student{name: name, cgpa: cgpa, id: id})
		} else if parts[0] == "SERVED" {
			if h.Len() > 0 {
				heap.Pop(h)
			}
		}
	}

	if h.Len() == 0 {
		return []string{"EMPTY"}
	}

	var results []string
	for h.Len() > 0 {
		s := heap.Pop(h).(Student)
		results = append(results, s.name)
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 이벤트 개수 입력
	var n int
	fmt.Fscan(reader, &n)
	reader.ReadString('\n')

	// 이벤트 입력
	events := make([]string, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		events[i] = strings.TrimSpace(line)
	}

	// 핵심 함수 호출
	results := processEvents(events)

	// 결과 출력
	for _, name := range results {
		fmt.Fprintln(writer, name)
	}
}
