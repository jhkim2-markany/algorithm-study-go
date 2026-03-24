package main

import (
	"bufio"
	"fmt"
	"os"
)

// printShortestPath는 레드 나이트의 최단 경로를 출력한다.
//
// [매개변수]
//   - n: 체스판 크기
//   - rStart, cStart: 시작 위치
//   - rEnd, cEnd: 목표 위치
//   - writer: 출력 버퍼
//
// [반환값]
//   - 없음 (표준 출력으로 최소 이동 횟수와 경로를 출력)
//
// [알고리즘 힌트]
//
//	BFS로 최단 경로를 탐색한다. 6가지 이동 방향을 우선순위 순서
//	(UL, UR, R, LR, LL, L)로 탐색하여 사전순 최소 경로를 보장한다.
//	부모 정보를 기록하여 경로를 역추적한다.
func printShortestPath(n int, rStart, cStart, rEnd, cEnd int, writer *bufio.Writer) {
	// 6가지 이동 방향 (우선순위 순서)
	dirNames := []string{"UL", "UR", "R", "LR", "LL", "L"}
	dr := []int{-2, -2, 0, 2, 2, 0}
	dc := []int{-1, 1, 2, 1, -1, -2}

	// 방문 배열 및 부모 정보
	type cell struct {
		r, c int
		dir  string
	}
	visited := make([][]bool, n)
	parent := make([][]cell, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
		parent[i] = make([]cell, n)
		for j := 0; j < n; j++ {
			parent[i][j] = cell{-1, -1, ""}
		}
	}

	// BFS 시작
	type pos struct{ r, c int }
	queue := []pos{{rStart, cStart}}
	visited[rStart][cStart] = true

	found := false
	for len(queue) > 0 && !found {
		size := len(queue)
		for s := 0; s < size && !found; s++ {
			cur := queue[s]

			// 목표 도달 확인
			if cur.r == rEnd && cur.c == cEnd {
				found = true
				break
			}

			// 6방향 탐색 (우선순위 순서)
			for d := 0; d < 6; d++ {
				nr, nc := cur.r+dr[d], cur.c+dc[d]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
					visited[nr][nc] = true
					parent[nr][nc] = cell{cur.r, cur.c, dirNames[d]}
					queue = append(queue, pos{nr, nc})
				}
			}
		}
		queue = queue[size:]
	}

	if !found {
		fmt.Fprintln(writer, "Impossible")
		return
	}

	// 경로 역추적
	var path []string
	r, c := rEnd, cEnd
	for r != rStart || c != cStart {
		p := parent[r][c]
		path = append(path, p.dir)
		r, c = p.r, p.c
	}

	// 경로 뒤집기
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// 결과 출력
	fmt.Fprintln(writer, len(path))
	for i, d := range path {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, d)
	}
	fmt.Fprintln(writer)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var rStart, cStart, rEnd, cEnd int
	fmt.Fscan(reader, &rStart, &cStart)
	fmt.Fscan(reader, &rEnd, &cEnd)

	printShortestPath(n, rStart, cStart, rEnd, cEnd, writer)
}
