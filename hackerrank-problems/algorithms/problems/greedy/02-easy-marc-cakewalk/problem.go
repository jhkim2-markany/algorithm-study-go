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
func marcsCakewalk(calorie []int) int64 {
	// 여기에 코드를 작성하세요
	return 0
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

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Ints
}
