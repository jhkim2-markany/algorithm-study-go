package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// luckBalance는 최대 K번의 중요한 대회에서 질 수 있을 때, 최대 행운 균형을 반환한다.
//
// [매개변수]
//   - k: 중요한 대회에서 질 수 있는 최대 횟수
//   - contests: 각 대회의 [행운값, 중요도] 배열
//
// [반환값]
//   - int: 최대 행운 균형
func luckBalance(k int, contests [][]int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N, K 입력
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 대회 정보 입력
	contests := make([][]int, n)
	for i := 0; i < n; i++ {
		contests[i] = make([]int, 2)
		fmt.Fscan(reader, &contests[i][0], &contests[i][1])
	}

	// 핵심 함수 호출 및 결과 출력
	result := luckBalance(k, contests)
	fmt.Fprintln(writer, result)

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Slice
}
