package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 격자 크기 입력
	var R, C int
	fmt.Fscan(reader, &R, &C)

	// 로봇 수와 명령 수 입력
	var N, M int
	fmt.Fscan(reader, &N, &M)

	// 방향별 이동량 (1:상, 2:우, 3:하, 4:좌)
	// 인덱스 0은 사용하지 않음
	dx := [5]int{0, -1, 0, 1, 0}
	dy := [5]int{0, 0, 1, 0, -1}

	// 로봇 정보 저장 (행, 열, 방향)
	type Robot struct {
		r, c, d int
	}
	robots := make([]Robot, N)

	// 격자에 로봇 위치 기록 (충돌 감지용)
	grid := make([][]int, R+1)
	for i := 0; i <= R; i++ {
		grid[i] = make([]int, C+1)
	}

	// 로봇 초기 정보 입력
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &robots[i].r, &robots[i].c, &robots[i].d)
		// 격자에 로봇 번호 기록 (1-indexed, 0은 빈 칸)
		grid[robots[i].r][robots[i].c] = i + 1
	}

	// 명령 문자열 입력
	var commands string
	fmt.Fscan(reader, &commands)

	// 각 로봇에 대해 순서대로 명령 수행
	for i := 0; i < N; i++ {
		for _, cmd := range commands {
			switch cmd {
			case 'L':
				// 좌회전: 상→좌→하→우→상 (1→4→3→2→1)
				robots[i].d = (robots[i].d+2)%4 + 1
			case 'R':
				// 우회전: 상→우→하→좌→상 (1→2→3→4→1)
				robots[i].d = robots[i].d%4 + 1
			case 'F':
				// 전진: 현재 방향으로 한 칸 이동
				nr := robots[i].r + dx[robots[i].d]
				nc := robots[i].c + dy[robots[i].d]

				// 격자 범위 검사
				if nr < 1 || nr > R || nc < 1 || nc > C {
					continue
				}
				// 다른 로봇과 충돌 검사
				if grid[nr][nc] != 0 {
					continue
				}

				// 이동 수행: 이전 위치 비우고 새 위치에 기록
				grid[robots[i].r][robots[i].c] = 0
				robots[i].r = nr
				robots[i].c = nc
				grid[nr][nc] = i + 1
			}
		}
	}

	// 각 로봇의 최종 상태 출력
	for i := 0; i < N; i++ {
		fmt.Fprintf(writer, "%d %d %d\n", robots[i].r, robots[i].c, robots[i].d)
	}
}
