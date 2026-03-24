package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// marcsCakewalk는 컵케이크를 먹은 후 걸어야 하는 최소 마일 수를 반환한다.
//
// [매개변수]
//   - calorie: 각 컵케이크의 칼로리 배열
//
// [반환값]
//   - int64: 최소 마일 수
//
// [알고리즘 힌트]
//
//	칼로리를 내림차순 정렬하여 높은 칼로리에 작은 2^i를 곱한다.
//	재배열 부등식에 의해 이 방법이 최적이다.
func marcsCakewalk(calorie []int) int64 {
	// 칼로리를 내림차순으로 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(calorie)))

	// 총 마일 수 계산
	var total int64
	for i, c := range calorie {
		// 2^i를 비트 시프트로 계산
		total += int64(c) << uint(i)
	}

	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 컵케이크 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 칼로리 배열 입력
	calorie := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &calorie[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := marcsCakewalk(calorie)
	fmt.Fprintln(writer, result)
}
