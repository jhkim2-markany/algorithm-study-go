package main

import (
	"bufio"
	"fmt"
	"os"
)

// Robot은 로봇의 위치와 방향 정보를 저장하는 구조체이다
type Robot struct {
	r, c, d int
}

// simulateRobots는 R×C 격자 위의 로봇들에게 명령 문자열을 순서대로
// 수행시킨 후 최종 상태를 반환한다.
//
// [매개변수]
//   - robots: 로봇 배열. 각 로봇은 (r, c, d)로 행, 열, 방향을 가진다
//   - commands: 명령 문자열. 'L'(좌회전), 'R'(우회전), 'F'(전진)로 구성
//   - R: 격자의 행 수 (1~R)
//   - C: 격자의 열 수 (1~C)
//
// [반환값]
//   - []Robot: 모든 명령 수행 후 각 로봇의 최종 상태 (위치 + 방향)
func simulateRobots(robots []Robot, commands string, R, C int) []Robot {
	// 여기에 코드를 작성하세요
	return nil
}

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

	// 로봇 초기 정보 입력
	robots := make([]Robot, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &robots[i].r, &robots[i].c, &robots[i].d)
	}

	// 명령 문자열 입력
	var commands string
	fmt.Fscan(reader, &commands)

	// 핵심 함수 호출
	result := simulateRobots(robots, commands, R, C)

	// 각 로봇의 최종 상태 출력
	for i := 0; i < len(result); i++ {
		fmt.Fprintf(writer, "%d %d %d\n", result[i].r, result[i].c, result[i].d)
	}
}
