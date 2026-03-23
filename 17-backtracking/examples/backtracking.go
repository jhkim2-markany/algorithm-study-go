package main

import "fmt"

// 백트래킹 기본 구현 - 순열과 조합 생성
// 시간 복잡도: 순열 O(N!), 조합 O(C(N, R))
// 공간 복잡도: O(N) (재귀 깊이)

// permutation 함수는 1~N에서 R개를 뽑는 순열을 생성한다
func permutation(n, r int) [][]int {
	result := [][]int{}
	chosen := make([]int, 0, r)
	used := make([]bool, n+1)

	var backtrack func()
	backtrack = func() {
		// 종료 조건: R개를 모두 선택했으면 결과에 추가
		if len(chosen) == r {
			temp := make([]int, r)
			copy(temp, chosen)
			result = append(result, temp)
			return
		}

		for i := 1; i <= n; i++ {
			// 가지치기: 이미 사용한 숫자는 건너뛴다
			if used[i] {
				continue
			}

			// 선택: 숫자를 추가하고 사용 표시
			chosen = append(chosen, i)
			used[i] = true

			// 재귀 호출: 다음 숫자 선택
			backtrack()

			// 되돌리기: 선택을 취소하고 사용 표시 해제
			chosen = chosen[:len(chosen)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

// combination 함수는 1~N에서 R개를 뽑는 조합을 생성한다
func combination(n, r int) [][]int {
	result := [][]int{}
	chosen := make([]int, 0, r)

	var backtrack func(start int)
	backtrack = func(start int) {
		// 종료 조건: R개를 모두 선택했으면 결과에 추가
		if len(chosen) == r {
			temp := make([]int, r)
			copy(temp, chosen)
			result = append(result, temp)
			return
		}

		for i := start; i <= n; i++ {
			// 선택: 숫자를 추가
			chosen = append(chosen, i)

			// 재귀 호출: 다음 숫자는 i+1부터 선택 (중복 방지)
			backtrack(i + 1)

			// 되돌리기: 선택을 취소
			chosen = chosen[:len(chosen)-1]
		}
	}

	backtrack(1)
	return result
}

func main() {
	// 순열 예제: 1~3에서 2개를 뽑는 순열
	fmt.Println("=== 순열 (N=3, R=2) ===")
	perms := permutation(3, 2)
	for _, p := range perms {
		fmt.Println(p)
	}

	// 조합 예제: 1~4에서 2개를 뽑는 조합
	fmt.Println("\n=== 조합 (N=4, R=2) ===")
	combs := combination(4, 2)
	for _, c := range combs {
		fmt.Println(c)
	}
}
