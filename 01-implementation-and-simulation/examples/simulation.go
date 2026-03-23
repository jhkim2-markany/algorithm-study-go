package main

import "fmt"

// 로봇 이동 시뮬레이션 - 구현과 시뮬레이션의 기본 예시
// 2차원 격자에서 명령에 따라 로봇을 이동시키는 시뮬레이션이다.
// 시간 복잡도: O(K) (K: 명령 수)
// 공간 복잡도: O(1)

// 방향 상수 정의 (상, 우, 하, 좌)
const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

// 방향별 이동량 (dx: 행 변화, dy: 열 변화)
var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}

// simulate 함수는 N×N 격자에서 명령에 따라 로봇을 이동시킨다.
// 명령: 'G' = 전진, 'L' = 좌회전, 'R' = 우회전
// 격자 범위를 벗어나는 이동은 무시한다.
func simulate(n int, commands string) (int, int, int) {
	// 로봇의 초기 위치와 방향 설정
	x, y := 0, 0
	dir := RIGHT // 초기 방향: 오른쪽

	// 각 명령을 순서대로 처리
	for _, cmd := range commands {
		switch cmd {
		case 'G':
			// 전진: 현재 방향으로 한 칸 이동
			nx, ny := x+dx[dir], y+dy[dir]
			// 격자 범위 검사
			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				x, y = nx, ny
			}
		case 'L':
			// 좌회전: 방향을 반시계 방향으로 90도 회전
			dir = (dir + 3) % 4
		case 'R':
			// 우회전: 방향을 시계 방향으로 90도 회전
			dir = (dir + 1) % 4
		}
	}

	return x, y, dir
}

// 방향을 한국어 문자열로 변환
func dirName(dir int) string {
	names := [4]string{"위", "오른쪽", "아래", "왼쪽"}
	return names[dir]
}

func main() {
	// 5×5 격자에서 로봇 이동 시뮬레이션 실행
	n := 5
	commands := "GGRGGRGGLG"

	fmt.Printf("격자 크기: %d×%d\n", n, n)
	fmt.Printf("명령어: %s\n", commands)

	x, y, dir := simulate(n, commands)

	fmt.Printf("최종 위치: (%d, %d)\n", x, y)
	fmt.Printf("최종 방향: %s\n", dirName(dir))

	// 추가 예시: 경계에서의 동작 확인
	fmt.Println("\n--- 경계 테스트 ---")
	commands2 := "GGGGGGGGG" // 오른쪽으로 계속 전진
	x2, y2, dir2 := simulate(n, commands2)
	fmt.Printf("명령어: %s\n", commands2)
	fmt.Printf("최종 위치: (%d, %d) (경계에서 멈춤)\n", x2, y2)
	fmt.Printf("최종 방향: %s\n", dirName(dir2))
}
