package main

import (
	"bufio"
	"fmt"
	"os"
)

// cutLanCable은 랜선들을 길이 l로 잘라 n개 이상 만들 수 있는
// 최대 랜선 길이를 반환한다.
//
// [매개변수]
//   - cables: 각 랜선의 길이 배열
//   - n: 필요한 랜선 개수
//
// [반환값]
//   - int: 조건을 만족하는 최대 랜선 길이
func cutLanCable(cables []int, n int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 랜선 수 K와 필요한 랜선 수 N 입력
	var k, n int
	fmt.Fscan(reader, &k, &n)

	// 각 랜선의 길이 입력
	cables := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &cables[i])
	}

	// 핵심 함수 호출
	result := cutLanCable(cables, n)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
