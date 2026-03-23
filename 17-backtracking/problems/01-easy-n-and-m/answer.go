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
//
// [알고리즘 힌트]
//
//	백트래킹으로 순열을 생성한다.
//	사용 여부 배열(used)로 중복 선택을 방지하고,
//	1부터 N까지 순서대로 시도하여 사전 순을 보장한다.
//	M개를 모두 선택하면 결과에 추가하고 되돌린다.
func generatePermutations(n, m int) [][]int {
	var result [][]int
	chosen := make([]int, 0, m)
	used := make([]bool, n+1)

	var backtrack func()
	backtrack = func() {
		if len(chosen) == m {
			perm := make([]int, m)
			copy(perm, chosen)
			result = append(result, perm)
			return
		}
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			chosen = append(chosen, i)
			used[i] = true
			backtrack()
			chosen = chosen[:len(chosen)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
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
