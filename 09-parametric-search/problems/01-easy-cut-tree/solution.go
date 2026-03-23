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

	// 나무 수 N과 필요한 나무 길이 M 입력
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 각 나무의 높이 입력
	trees := make([]int, n)
	maxH := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &trees[i])
		if trees[i] > maxH {
			maxH = trees[i]
		}
	}

	// 파라메트릭 서치: 절단 높이의 최댓값을 이진 탐색으로 찾는다
	lo, hi := 0, maxH
	result := 0

	for lo <= hi {
		mid := (lo + hi) / 2

		// 결정 함수: 높이 mid로 잘랐을 때 M 이상을 얻을 수 있는가?
		total := 0
		for _, h := range trees {
			if h > mid {
				total += h - mid
			}
		}

		if total >= m {
			// 조건 만족: 더 높은 절단 높이도 가능한지 확인
			result = mid
			lo = mid + 1
		} else {
			// 조건 불만족: 절단 높이를 낮춘다
			hi = mid - 1
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, result)
}
