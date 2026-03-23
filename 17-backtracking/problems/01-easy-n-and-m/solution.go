package main

import (
	"bufio"
	"fmt"
	"os"
)

// generatePermutations는 1부터 N까지의 수에서 M개를 선택한 순열을 사전 순으로 반환한다.
//
// [매개변수]
//   - n: 수의 범위 (1~N)
//   - m: 선택할 수의 개수
//
// [반환값]
//   - [][]int: 사전 순으로 정렬된 순열 목록
func generatePermutations(n, m int) [][]int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 핵심 함수 호출
	perms := generatePermutations(n, m)

	// 결과 출력
	for _, perm := range perms {
		for i, v := range perm {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}
