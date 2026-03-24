package main

import (
	"bufio"
	"fmt"
	"os"
)

// quickestWayUp은 뱀과 사다리 게임에서 100번 칸에 도달하는 최소 주사위 횟수를 반환한다.
//
// [매개변수]
//   - ladders: 사다리 목록 (각 원소는 [2]int{시작, 끝})
//   - snakes: 뱀 목록 (각 원소는 [2]int{시작, 끝})
//
// [반환값]
//   - int: 최소 주사위 횟수 (-1이면 도달 불가)
//
// [알고리즘 힌트]
//
//	사다리와 뱀을 매핑 테이블에 저장한 뒤,
//	1번 칸에서 BFS를 수행하여 100번 칸까지의 최단 거리를 구한다.
func quickestWayUp(ladders [][2]int, snakes [][2]int) int {
	// 사다리와 뱀 매핑 테이블 구성
	move := make(map[int]int)
	for _, l := range ladders {
		move[l[0]] = l[1]
	}
	for _, s := range snakes {
		move[s[0]] = s[1]
	}

	// BFS 수행
	dist := make([]int, 101)
	for i := range dist {
		dist[i] = -1
	}
	dist[1] = 0
	queue := []int{1}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// 주사위 1~6
		for dice := 1; dice <= 6; dice++ {
			next := cur + dice
			if next > 100 {
				continue
			}
			// 사다리 또는 뱀이 있으면 이동
			if dest, ok := move[next]; ok {
				next = dest
			}
			if dist[next] == -1 {
				dist[next] = dist[cur] + 1
				if next == 100 {
					return dist[next]
				}
				queue = append(queue, next)
			}
		}
	}

	return dist[100]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(reader, &n)
		ladders := make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &ladders[i][0], &ladders[i][1])
		}

		var m int
		fmt.Fscan(reader, &m)
		snakes := make([][2]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &snakes[i][0], &snakes[i][1])
		}

		result := quickestWayUp(ladders, snakes)
		fmt.Fprintln(writer, result)
	}
}
