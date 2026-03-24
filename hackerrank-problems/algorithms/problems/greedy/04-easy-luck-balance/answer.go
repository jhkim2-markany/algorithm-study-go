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
//
// [알고리즘 힌트]
//
//	중요하지 않은 대회는 모두 진다. 중요한 대회는 행운 값이 큰 순서대로 K개를 지고,
//	나머지는 이긴다. 이것이 행운 균형을 최대화하는 그리디 전략이다.
func luckBalance(k int, contests [][]int) int {
	// 중요한 대회의 행운 값을 별도로 수집
	var important []int
	luck := 0

	for _, c := range contests {
		if c[1] == 0 {
			// 중요하지 않은 대회는 모두 진다 (행운 추가)
			luck += c[0]
		} else {
			// 중요한 대회의 행운 값 수집
			important = append(important, c[0])
		}
	}

	// 중요한 대회를 행운 값 내림차순으로 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(important)))

	// 상위 K개는 지고 (행운 추가), 나머지는 이긴다 (행운 차감)
	for i, v := range important {
		if i < k {
			luck += v
		} else {
			luck -= v
		}
	}

	return luck
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
}
