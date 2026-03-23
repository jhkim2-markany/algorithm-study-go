package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4방향 이동: 상, 하, 좌, 우
var dr = [4]int{-1, 1, 0, 0}
var dc = [4]int{0, 0, -1, 1}

// 유니온 파인드 구조체
var parent []int
var rank_ []int

// find는 x가 속한 집합의 대표를 반환한다 (경로 압축)
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// union은 두 집합을 합친다 (랭크 기반)
func union(a, b int) {
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 격자 크기와 쿼리 수
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)

	// 입력: 격자 정보
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	// 유니온 파인드 초기화
	total := n * m
	parent = make([]int, total)
	rank_ = make([]int, total)
	for i := 0; i < total; i++ {
		parent[i] = i
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
			// 오른쪽, 아래쪽만 확인 (중복 방지)
			if c+1 < m && grid[r][c] == grid[r][c+1] {
				union(idx, idx+1)
			}
			if r+1 < n && grid[r][c] == grid[r+1][c] {
				union(idx, (r+1)*m+c)
			}
		}
	}

	// 쿼리 처리
	for i := 0; i < q; i++ {
		var r, c, x int
		fmt.Fscan(reader, &r, &c, &x)
		r-- // 0-indexed로 변환
		c--

		root := find(r*m + c)
		oldColor := color[root]

		// 이미 같은 색이면 건너뛴다
		if oldColor == x {
			continue
		}

		// BFS로 같은 영역(같은 대표)에 속한 칸들의 경계를 찾아 합친다
		// 먼저 대표의 색상을 변경
		color[root] = x

		// 영역에 속한 칸들을 BFS로 수집
		visited := map[int]bool{root: true}
		queue := []int{root}

		// 대표가 같은 모든 칸을 찾기 위해 격자 전체를 스캔
		members := []int{}
		for rr := 0; rr < n; rr++ {
			for cc := 0; cc < m; cc++ {
				if find(rr*m+cc) == root {
					members = append(members, rr*m+cc)
				}
			}
		}

		// 영역의 경계에서 인접한 같은 색상(x) 영역과 합친다
		queue = members
		visited = map[int]bool{}
		for _, idx := range members {
			visited[idx] = true
		}

		for _, idx := range queue {
			rr, cc := idx/m, idx%m
			for d := 0; d < 4; d++ {
				nr, nc := rr+dr[d], cc+dc[d]
				if nr >= 0 && nr < n && nc >= 0 && nc < m {
					nIdx := nr*m + nc
					nRoot := find(nIdx)
					if !visited[nRoot] && color[nRoot] == x {
						union(root, nIdx)
						// 합친 후 대표의 색상을 유지
						newRoot := find(root)
						color[newRoot] = x
						root = newRoot
					}
				}
			}
		}
	}

	// 출력: 최종 격자 상태
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if c > 0 {
				fmt.Fprint(writer, " ")
			}
			root := find(r*m + c)
			fmt.Fprint(writer, color[root])
		}
		fmt.Fprintln(writer)
	}
}
