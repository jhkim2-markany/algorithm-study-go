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
//
// [알고리즘 힌트]
//
//	스택에 (살충제 양, 사망일) 쌍을 저장한다.
//	현재 식물이 왼쪽 이웃보다 크면 죽는 날짜를 계산한다.
//	스택에서 현재 식물보다 크거나 같은 원소를 팝하며 최대 사망일을 추적한다.
func poisonousPlants(p []int) int {
	type entry struct {
		val int
		day int
	}

	// 스택 초기화
	stack := []entry{}
	maxDays := 0

	for i := 0; i < len(p); i++ {
		// 현재 식물의 사망일 초기화
		curDay := 0

		// 스택에서 현재 식물보다 크거나 같은 원소를 팝
		for len(stack) > 0 && stack[len(stack)-1].val >= p[i] {
			// 팝하면서 최대 사망일 추적
			if stack[len(stack)-1].day > curDay {
				curDay = stack[len(stack)-1].day
			}
			stack = stack[:len(stack)-1]
		}

		// 스택이 비어있으면 왼쪽에 더 작은 식물이 없으므로 죽지 않음
		if len(stack) == 0 {
			curDay = 0
		} else {
			// 왼쪽에 더 작은 식물이 있으므로 curDay + 1일에 죽음
			curDay++
		}

		// 최대 일수 갱신
		if curDay > maxDays {
			maxDays = curDay
		}

		// 현재 식물을 스택에 푸시
		stack = append(stack, entry{p[i], curDay})
	}

	return maxDays
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
