package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Job은 작업의 마감 기한과 보상을 나타낸다.
type Job struct {
	deadline, profit int
}

// jobScheduling은 마감 기한과 보상이 주어진 작업들에서 최대 보상을 반환한다.
//
// [매개변수]
//   - deadlines: 각 작업의 마감 기한 배열
//   - profits: 각 작업의 보상 배열
//
// [반환값]
//   - int: 얻을 수 있는 최대 보상
//
// [알고리즘 힌트]
//
//	그리디: 보상이 큰 작업부터 처리한다.
//	보상 내림차순 정렬 후, 각 작업의 마감 기한 당일부터
//	1일차까지 역순으로 빈 날짜를 탐색하여 배정한다.
//	scheduled 배열로 날짜별 배정 여부를 추적한다.
//
//	시간복잡도: O(N × D), D는 최대 마감 기한
func jobScheduling(deadlines, profits []int) int {
	n := len(deadlines)
	jobs := make([]Job, n)
	maxDeadline := 0
	for i := 0; i < n; i++ {
		jobs[i] = Job{deadlines[i], profits[i]}
		if deadlines[i] > maxDeadline {
			maxDeadline = deadlines[i]
		}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].profit > jobs[j].profit
	})

	scheduled := make([]bool, maxDeadline+1)
	totalProfit := 0

	for _, job := range jobs {
		for day := job.deadline; day >= 1; day-- {
			if !scheduled[day] {
				scheduled[day] = true
				totalProfit += job.profit
				break
			}
		}
	}
	return totalProfit
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	deadlines := make([]int, n)
	profits := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &deadlines[i], &profits[i])
	}

	result := jobScheduling(deadlines, profits)
	fmt.Fprintln(writer, result)
}
