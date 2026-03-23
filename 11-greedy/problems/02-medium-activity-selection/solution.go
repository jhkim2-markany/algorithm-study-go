package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Meeting 구조체는 회의의 시작 시간과 종료 시간을 나타낸다
type Meeting struct {
	start int
	end   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 회의 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 각 회의의 시작/종료 시간 입력
	meetings := make([]Meeting, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &meetings[i].start, &meetings[i].end)
	}

	// 종료 시간 기준 오름차순 정렬 (종료 시간이 같으면 시작 시간 기준)
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].end == meetings[j].end {
			return meetings[i].start < meetings[j].start
		}
		return meetings[i].end < meetings[j].end
	})

	// 그리디: 종료 시간이 빠른 회의부터 선택
	count := 1                 // 첫 번째 회의 선택
	lastEnd := meetings[0].end // 마지막으로 선택한 회의의 종료 시간

	for i := 1; i < n; i++ {
		// 현재 회의의 시작 시간이 마지막 선택 회의의 종료 시간 이후이면 선택
		if meetings[i].start >= lastEnd {
			count++
			lastEnd = meetings[i].end
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, count)
}
