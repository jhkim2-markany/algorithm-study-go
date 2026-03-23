package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m   int
	chosen []int
	used   []bool
	writer *bufio.Writer
)

// backtrack 함수는 현재까지 선택한 수열을 확장한다
func backtrack() {
	// 종료 조건: M개를 모두 선택했으면 출력
	if len(chosen) == m {
		for i, v := range chosen {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
		return
	}

	// 1부터 N까지 순서대로 시도 (사전 순 보장)
	for i := 1; i <= n; i++ {
		// 가지치기: 이미 사용한 숫자는 건너뛴다
		if used[i] {
			continue
		}

		// 선택
		chosen = append(chosen, i)
		used[i] = true

		// 재귀 호출
		backtrack()

		// 되돌리기
		chosen = chosen[:len(chosen)-1]
		used[i] = false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	fmt.Fscan(reader, &n, &m)

	chosen = make([]int, 0, m)
	used = make([]bool, n+1)

	// 백트래킹으로 순열 생성
	backtrack()
}
