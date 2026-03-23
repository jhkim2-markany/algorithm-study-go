package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 날짜 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 주식 가격 입력
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &prices[i])
	}

	// 결과 배열: 각 날짜의 가격이 떨어지지 않은 기간
	answer := make([]int, n)

	// 스택에는 아직 가격이 떨어지는 날을 찾지 못한 날짜의 인덱스를 저장한다
	stack := []int{}

	for i := 0; i < n; i++ {
		// 현재 가격이 스택 맨 위 날짜의 가격보다 낮으면
		// 스택 맨 위 날짜의 답을 확정한다
		for len(stack) > 0 && prices[stack[len(stack)-1]] > prices[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 가격이 떨어지지 않은 기간 = 현재 날짜 - 해당 날짜
			answer[top] = i - top
		}
		stack = append(stack, i)
	}

	// 스택에 남은 날짜들은 끝까지 가격이 떨어지지 않은 경우
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		answer[top] = (n - 1) - top
	}

	// 결과 출력
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, answer[i])
	}
	fmt.Fprintln(writer)
}
