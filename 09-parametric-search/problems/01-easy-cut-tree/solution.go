package main

import (
	"bufio"
	"fmt"
	"os"
)

// cutTree는 나무들을 높이 h로 잘랐을 때 m 이상의 나무를 얻을 수 있는
// 최대 절단 높이를 반환한다.
//
// [매개변수]
//   - trees: 각 나무의 높이 배열
//   - m: 필요한 나무 길이의 합
//
// [반환값]
//   - int: 조건을 만족하는 최대 절단 높이
func cutTree(trees []int, m int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 나무 수 N과 필요한 나무 길이 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 각 나무의 높이 입력
	trees := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &trees[i])
	}

	// 핵심 함수 호출
	result := cutTree(trees, m)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
