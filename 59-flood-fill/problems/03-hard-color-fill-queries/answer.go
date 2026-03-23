package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

// colorFillQueries는 색상이 칠해진 격자에서 플러드 필 방식의
// 색상 변경 쿼리들을 유니온 파인드를 이용하여 효율적으로 처리한다.
// 각 쿼리는 (r, c) 위치와 연결된 같은 색상 영역 전체를 새 색상으로 변경한다.
//
// [매개변수]
//   - grid: n×m 크기의 격자 (각 칸에 색상 번호)
//   - n: 격자의 행 수
//   - m: 격자의 열 수
//   - queries: 쿼리 목록 ([r, c, newColor], 1-indexed 좌표)
//
// [반환값]
//   - [][]int: 모든 쿼리 처리 후의 최종 격자 상태
//
// [알고리즘 힌트]
//   1. 유니온 파인드로 같은 색상의 인접 칸들을 하나의 집합으로 관리한다
//   2. 초기 상태에서 인접한 같은 색상 칸들을 모두 합친다
//   3. 쿼리마다 대표 노드의 색상을 변경한 뒤, 영역에 속한 칸들의 경계에서
//      새 색상과 같은 인접 영역을 합친다
//   4. 최종 격자는 각 칸의 대표 노드 색상으로 복원한다
func colorFillQueries(grid [][]int, n, m int, queries [][3]int) [][]int {
	total := n * m

	// 유니온 파인드
	parent := make([]int, total)
	rank_ := make([]int, total)
	for i := 0; i < total; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		a, b = find(a), find(b)
		if a == b {
			return
		}
		if rank_[a] < rank_[b] {
			a, b = b, a
		}
		parent[b] = a
		if rank_[a] == rank_[b] {
			rank_[a]++
		}
	}

	// 색상 배열: 각 대표 노드의 색상
	color := make([]int, total)
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			color[r*m+c] = grid[r][c]
		}
	}

	// 초기 상태에서 인접한 같은 색상 칸을 합친다
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			idx := r*m + c
			if c+1 < m && grid[r][c] == grid[r][c+1] {
				union(idx, idx+1)
			}
			if r+1 < n && grid[r][c] == grid[r+1][c] {
				union(idx, (r+1)*m+c)
			}
		}
	}

	// 쿼리 처리
	for _, qr := range queries {
		r, c, x := qr[0]-1, qr[1]-1, qr[2]

		root := find(r*m + c)
		oldColor := color[root]
		if oldColor == x {
			continue
		}

		color[root] = x

		// 영역에 속한 칸들을 수집
		members := []int{}
		for rr := 0; rr < n; rr++ {
			for cc := 0; cc < m; cc++ {
				if find(rr*m+cc) == root {
					members = append(members, rr*m+cc)
				}
			}
		}

		// 경계에서 인접한 같은 색상 영역과 합친다
		for _, idx := range members {
			rr, cc := idx/m, idx%m
			for d := 0; d < 4; d++ {
				nr, nc := rr+dr[d], cc+dc[d]
				if nr >= 0 && nr < n && nc >= 0 && nc < m {
					nRoot := find(nr*m + nc)
					if color[nRoot] == x {
						union(root, nr*m+nc)
						newRoot := find(root)
						color[newRoot] = x
						root = newRoot
					}
				}
			}
		}
	}

	// 최종 격자 복원
	result := make([][]int, n)
	for r := 0; r < n; r++ {
		result[r] = make([]int, m)
		for c := 0; c < m; c++ {
			result[r][c] = color[find(r*m+c)]
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	queries := make([][3]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}

	result := colorFillQueries(grid, n, m, queries)

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if c > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, result[r][c])
		}
		fmt.Fprintln(writer)
	}
}
