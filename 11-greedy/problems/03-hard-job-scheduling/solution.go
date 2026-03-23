package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Job 구조체는 작업의 마감 기한과 보상을 나타낸다
type Job struct {
	deadline int
	profit   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 작업 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 각 작업의 마감 기한과 보상 입력
	jobs := make([]Job, n)
	maxDeadline := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &jobs[i].deadline, &jobs[i].profit)
		if jobs[i].deadline > maxDeadline {
			maxDeadline = jobs[i].deadline
		}
	}

	// 보상 기준 내림차순 정렬 (보상이 큰 작업 우선)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].profit > jobs[j].profit
	})

	// 날짜별 작업 배정 여부를 추적하는 배열 (1-indexed)
	scheduled := make([]bool, maxDeadline+1)
	totalProfit := 0

	for _, job := range jobs {
		// 마감 기한 당일부터 1일차까지 역순으로 빈 날짜 탐색
		for day := job.deadline; day >= 1; day-- {
			if !scheduled[day] {
				// 빈 날짜에 작업 배정
				scheduled[day] = true
				totalProfit += job.profit
				break
			}
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, totalProfit)
}
