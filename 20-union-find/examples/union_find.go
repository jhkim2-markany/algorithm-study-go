package main

import "fmt"

// 유니온 파인드(Disjoint Set Union) 기본 구현
// 경로 압축(Path Compression)과 랭크 기반 합치기(Union by Rank)를 적용
// 시간 복잡도: Find/Union 모두 O(α(N)) ≈ O(1)
// 공간 복잡도: O(N)

// parent[i]: i번 노드의 부모 노드
// rank[i]: i번 노드를 루트로 하는 트리의 랭크(높이 상한)
var parent []int
var rank_ []int

// init 함수는 N개의 원소를 각각 독립된 집합으로 초기화한다
func initialize(n int) {
	parent = make([]int, n)
	rank_ = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 자기 자신이 루트
		rank_[i] = 0  // 초기 랭크는 0
	}
}

// find 함수는 x가 속한 집합의 루트를 반환한다 (경로 압축 적용)
func find(x int) int {
	if parent[x] != x {
		// 경로 압축: 재귀적으로 루트를 찾으면서 부모를 루트로 갱신
		parent[x] = find(parent[x])
	}
	return parent[x]
}

// union 함수는 x와 y가 속한 두 집합을 하나로 합친다 (랭크 기반)
func union(x, y int) bool {
	rootX := find(x)
	rootY := find(y)

	// 이미 같은 집합이면 합치지 않음
	if rootX == rootY {
		return false
	}

	// 랭크가 낮은 트리를 높은 트리 아래에 붙인다
	if rank_[rootX] < rank_[rootY] {
		parent[rootX] = rootY
	} else if rank_[rootX] > rank_[rootY] {
		parent[rootY] = rootX
	} else {
		// 랭크가 같으면 한쪽을 다른 쪽 아래에 붙이고 랭크 증가
		parent[rootY] = rootX
		rank_[rootX]++
	}
	return true
}

// connected 함수는 x와 y가 같은 집합에 속하는지 판별한다
func connected(x, y int) bool {
	return find(x) == find(y)
}

func main() {
	// 예제: 6개의 노드(0~5)와 간선들
	n := 6
	initialize(n)

	// 간선 추가: (0,1), (1,2), (3,4)
	edges := [][2]int{{0, 1}, {1, 2}, {3, 4}}

	fmt.Println("=== 유니온 파인드 기본 예제 ===")
	fmt.Printf("노드 수: %d\n\n", n)

	// 간선을 하나씩 추가하며 Union 수행
	for _, e := range edges {
		union(e[0], e[1])
		fmt.Printf("Union(%d, %d) 수행\n", e[0], e[1])
	}

	// 연결 상태 확인
	fmt.Println("\n=== 연결 상태 확인 ===")
	fmt.Printf("0과 2가 연결됨? %v\n", connected(0, 2)) // true (0-1-2)
	fmt.Printf("0과 3이 연결됨? %v\n", connected(0, 3)) // false
	fmt.Printf("3과 4가 연결됨? %v\n", connected(3, 4)) // true
	fmt.Printf("4와 5가 연결됨? %v\n", connected(4, 5)) // false

	// 연결 요소 개수 세기
	roots := make(map[int]bool)
	for i := 0; i < n; i++ {
		roots[find(i)] = true
	}
	fmt.Printf("\n연결 요소 개수: %d\n", len(roots)) // 3개: {0,1,2}, {3,4}, {5}

	// 추가 Union 후 확인
	fmt.Println("\n=== Union(2, 4) 수행 후 ===")
	union(2, 4)
	fmt.Printf("0과 3이 연결됨? %v\n", connected(0, 3)) // true (0-1-2-4-3)

	roots = make(map[int]bool)
	for i := 0; i < n; i++ {
		roots[find(i)] = true
	}
	fmt.Printf("연결 요소 개수: %d\n", len(roots)) // 2개: {0,1,2,3,4}, {5}
}
