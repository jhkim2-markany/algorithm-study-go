package main

import (
	"bufio"
	"fmt"
	"os"
)

// poisonousPlants는 더 이상 식물이 죽지 않을 때까지의 일수를 반환한다.
//
// [매개변수]
//   - p: 각 식물의 살충제 양 배열
//
// [반환값]
//   - int: 더 이상 식물이 죽지 않을 때까지의 일수
func poisonousPlants(p []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 식물 개수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 살충제 양 입력
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	// 핵심 함수 호출
	result := poisonousPlants(p)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
