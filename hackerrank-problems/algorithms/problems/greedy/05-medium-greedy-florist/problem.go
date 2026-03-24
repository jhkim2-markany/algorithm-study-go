package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// getMinimumCost는 K명의 친구가 모든 꽃을 구매할 때 최소 비용을 반환한다.
//
// [매개변수]
//   - k: 친구의 수
//   - c: 각 꽃의 원래 가격 배열
//
// [반환값]
//   - int: 최소 총 비용
func getMinimumCost(k int, c []int) int {
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

	// 꽃 가격 배열 입력
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := getMinimumCost(k, c)
	fmt.Fprintln(writer, result)

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Ints
}
