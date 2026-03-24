package main

import (
	"bufio"
	"fmt"
	"os"
)

// waiter는 접시 분류 결과를 반환한다.
//
// [매개변수]
//   - number: 접시 번호 배열 (스택의 아래에서 위 순서)
//   - q: 반복 횟수
//
// [반환값]
//   - []int: 결과 접시 번호 목록
func waiter(number []int, q int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 접시 개수와 반복 횟수 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 접시 번호 입력
	number := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &number[i])
	}

	// 핵심 함수 호출
	result := waiter(number, q)

	// 결과 출력
	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}
