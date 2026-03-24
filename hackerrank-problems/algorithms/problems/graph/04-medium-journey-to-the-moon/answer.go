package main

import (
	"bufio"
	"fmt"
	"os"
)

// journeyToMoon은 서로 다른 나라에서 2명을 선발하는 경우의 수를 반환한다.
//
// [매개변수]
//   - n: 우주비행사 수
//   - pairs: 같은 나라 출신 쌍 목록
//
// [반환값]
//   - int64: 가능한 조합 수
//
// [알고리즘 힌트]
//
//	Union-Find로 연결 요소를 구한 뒤,
//	전체 조합 C(N,2)에서 같은 그룹 내 조합을 빼서 답을 구한다.
func journeyToMoon(n int, pairs [][2]int) int64 {
	// Union-Find 초기화
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	// Find 함수 (경로 압축)
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// Union 함수 (랭크 기반)
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if rank[ra] < rank[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		if rank[ra] == rank[rb] {
			rank[ra]++
		}
	}

	// 같은 나라 쌍을 합치기
	for _, p := range pairs {
		union(p[0], p[1])
	}

	// 각 그룹의 크기 계산
	groupSize := make(map[int]int)
	for i := 0; i < n; i++ {
		groupSize[find(i)]++
	}

	// 전체 조합에서 같은 그룹 내 조합을 빼기
	total := int64(n) * int64(n-1) / 2
	for _, size := range groupSize {
		total -= int64(size) * int64(size-1) / 2
	}

	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, p int
	fmt.Fscan(reader, &n, &p)

	pairs := make([][2]int, p)
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &pairs[i][0], &pairs[i][1])
	}

	result := journeyToMoon(n, pairs)
	fmt.Fprintln(writer, result)
}
