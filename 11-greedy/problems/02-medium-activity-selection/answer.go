package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Meeting은 회의의 시작 시간과 종료 시간을 나타낸다.
type Meeting struct {
	start, end int
}

// activitySelection은 회의 목록에서 겹치지 않게 최대한 많은 회의를 선택한 수를 반환한다.
//
// [매개변수]
//   - starts: 각 회의의 시작 시간 배열
//   - ends: 각 회의의 종료 시간 배열
//
// [반환값]
//   - int: 겹치지 않게 선택할 수 있는 최대 회의 수
//
// [알고리즘 힌트]
//
//	그리디: 종료 시간이 빠른 회의부터 선택한다.
//	종료 시간 기준 오름차순 정렬 후,
//	현재 회의의 시작 시간이 마지막 선택 회의의 종료 시간 이후이면 선택.
//
//	시간복잡도: O(N log N)
func activitySelection(starts, ends []int) int {
	n := len(starts)
	meetings := make([]Meeting, n)
	for i := 0; i < n; i++ {
		meetings[i] = Meeting{starts[i], ends[i]}
	}

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].end == meetings[j].end {
			return meetings[i].start < meetings[j].start
		}
		return meetings[i].end < meetings[j].end
	})

	count := 1
	lastEnd := meetings[0].end
	for i := 1; i < n; i++ {
		if meetings[i].start >= lastEnd {
			count++
			lastEnd = meetings[i].end
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	starts := make([]int, n)
	ends := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &starts[i], &ends[i])
	}

	result := activitySelection(starts, ends)
	fmt.Fprintln(writer, result)
}
