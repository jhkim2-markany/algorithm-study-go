package main

import (
	"bufio"
	"fmt"
	"os"
)

// equalStacks는 세 스택의 높이가 같아질 때의 최대 높이를 반환한다.
//
// [매개변수]
//   - h1: 첫 번째 스택의 블록 높이 배열 (위에서 아래 순서)
//   - h2: 두 번째 스택의 블록 높이 배열
//   - h3: 세 번째 스택의 블록 높이 배열
//
// [반환값]
//   - int: 세 스택의 최대 동일 높이
func equalStacks(h1, h2, h3 []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 스택 크기 입력
	var n1, n2, n3 int
	fmt.Fscan(reader, &n1, &n2, &n3)

	// 각 스택의 블록 높이 입력
	h1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Fscan(reader, &h1[i])
	}
	h2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Fscan(reader, &h2[i])
	}
	h3 := make([]int, n3)
	for i := 0; i < n3; i++ {
		fmt.Fscan(reader, &h3[i])
	}

	// 핵심 함수 호출
	result := equalStacks(h1, h2, h3)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
