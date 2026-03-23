package main

import (
	"bufio"
	"fmt"
	"os"
)

// 유니온 파인드로 연결 요소 개수를 구하는 풀이
var parent []int
var rank_ []int

// find 함수는 x의 루트를 반환한다 (경로 압축 적용)
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// union 함수는 x와 y가 속한 집합을 합친다 (랭크 기반)
func union(x, y int) {
	rootX := find(x)
	rootY := find(y)
	if rootX == rootY {
		return
	}
	// 랭크가 낮은 트리를 높은 트리 아래에 붙인다
	if rank_[rootX] < rank_[rootY] {
		parent[rootX] = rootY
	} else if rank_[rootX] > rank_[rootY] {
		parent[rootY] = rootX
	} else {
		parent[rootY] = rootX
		rank_[rootX]++
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 초기화: 각 노드를 독립된 집합으로 설정
	parent = make([]int, n+1)
	rank_ = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	// 간선 입력 처리: 두 노드를 같은 집합으로 합친다
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		union(u, v)
	}

	// 연결 요소 개수 세기: 서로 다른 루트의 수를 센다
	count := 0
	for i := 1; i <= n; i++ {
		if find(i) == i {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}
