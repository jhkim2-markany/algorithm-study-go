package main

import (
	"bufio"
	"fmt"
	"os"
)

// 유니온 파인드로 여분의 간선(사이클을 형성하는 간선)을 찾는 풀이
// 간선을 순서대로 추가하며, 처음으로 사이클을 형성하는 간선이 답이다

var parent []int
var rank_ []int

// find 함수는 x의 루트를 반환한다 (경로 압축 적용)
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// union 함수는 x와 y가 속한 집합을 합친다
// 이미 같은 집합이면 false를 반환한다 (사이클 탐지)
func union(x, y int) bool {
	rootX := find(x)
	rootY := find(y)
	if rootX == rootY {
		// 같은 집합 → 이 간선을 추가하면 사이클 발생
		return false
	}
	// 랭크 기반 합치기
	if rank_[rootX] < rank_[rootY] {
		parent[rootX] = rootY
	} else if rank_[rootX] > rank_[rootY] {
		parent[rootY] = rootX
	} else {
		parent[rootY] = rootX
		rank_[rootX]++
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	// 유니온 파인드 초기화
	parent = make([]int, n+1)
	rank_ = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	// 간선을 순서대로 읽으며 Union 수행
	// 사이클을 형성하는 간선이 여분의 간선이다
	var ansU, ansV int
	for i := 0; i < n; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)

		if !union(u, v) {
			// Union 실패 = 이미 같은 집합 = 사이클 형성
			// 가장 마지막에 발견된 사이클 간선이 답
			ansU = u
			ansV = v
		}
	}

	fmt.Fprintf(writer, "%d %d\n", ansU, ansV)
}
