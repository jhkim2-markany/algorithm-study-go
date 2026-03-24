package main

import (
	"bufio"
	"fmt"
	"os"
)

// knightlOnAChessboard는 모든 (a, b) 쌍에 대해 KnightL의 최소 이동 횟수를 반환한다.
//
// [매개변수]
//   - n: 체스판 크기
//
// [반환값]
//   - [][]int: (N-1)×(N-1) 행렬, 각 원소는 최소 이동 횟수 (-1은 도달 불가)
//
// [알고리즘 힌트]
//
//	각 (a, b) 쌍에 대해 BFS로 (0,0)→(N-1,N-1) 최단 경로를 구한다.
//	KnightL(a,b) = KnightL(b,a) 대칭성을 활용한다.
func knightlOnAChessboard(n int) [][]int {
	// 결과 행렬 초기화
	result := make([][]int, n-1)
	for i := range result {
		result[i] = make([]int, n-1)
	}

	// BFS로 KnightL(a, b)의 최소 이동 횟수를 구하는 함수
	bfs := func(a, b int) int {
		// 시작점이 곧 도착점인 경우
		if n == 1 {
			return 0
		}

		// 방문 배열
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		// 8가지 이동 방향
		moves := [][2]int{
			{a, b}, {a, -b}, {-a, b}, {-a, -b},
			{b, a}, {b, -a}, {-b, a}, {-b, -a},
		}

		// BFS 시작
		type pos struct{ r, c int }
		queue := []pos{{0, 0}}
		visited[0][0] = true
		steps := 0

		for len(queue) > 0 {
			steps++
			size := len(queue)
			for s := 0; s < size; s++ {
				cur := queue[s]
				for _, m := range moves {
					nr, nc := cur.r+m[0], cur.c+m[1]
					// 범위 확인
					if nr < 0 || nr >= n || nc < 0 || nc >= n {
						continue
					}
					// 도착점 확인
					if nr == n-1 && nc == n-1 {
						return steps
					}
					// 미방문 셀 추가
					if !visited[nr][nc] {
						visited[nr][nc] = true
						queue = append(queue, pos{nr, nc})
					}
				}
			}
			// 현재 레벨 제거
			queue = queue[size:]
		}

		return -1
	}

	// 모든 (a, b) 쌍에 대해 BFS 수행
	for a := 1; a < n; a++ {
		for b := 1; b < n; b++ {
			if b < a {
				// 대칭성 활용: KnightL(a,b) = KnightL(b,a)
				result[a-1][b-1] = result[b-1][a-1]
			} else {
				result[a-1][b-1] = bfs(a, b)
			}
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	result := knightlOnAChessboard(n)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, result[i][j])
		}
		fmt.Fprintln(writer)
	}
}
