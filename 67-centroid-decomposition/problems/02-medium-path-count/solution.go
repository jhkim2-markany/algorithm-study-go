package main

import (
	"bufio"
	"fmt"
	"os"
)

// 특정 거리 경로 수 세기
// 센트로이드 분할 + 해시맵으로 정확히 거리 K인 쌍을 센다.

const MAXN = 50001

var (
	adj     [MAXN][]Edge // 가중치 포함 인접 리스트
	sz      [MAXN]int    // 서브트리 크기
	removed [MAXN]bool   // 센트로이드 제거 표시
	n, k    int
	answer  int64
)

// Edge는 가중치 간선을 나타낸다
type Edge struct {
	to, weight int
}

// calcSize는 현재 서브트리의 크기를 계산한다
func calcSize(v, p int) int {
	sz[v] = 1
	for _, e := range adj[v] {
		if e.to == p || removed[e.to] {
			continue
		}
		sz[v] += calcSize(e.to, v)
	}
	return sz[v]
}

// findCentroid는 센트로이드를 찾는다
func findCentroid(v, p, treeSize int) int {
	for _, e := range adj[v] {
		if e.to == p || removed[e.to] {
			continue
		}
		if sz[e.to] > treeSize/2 {
			return findCentroid(e.to, v, treeSize)
		}
	}
	return v
}

// getDists는 센트로이드에서 서브트리 내 모든 노드까지의 가중치 거리를 수집한다
func getDists(v, p, dist int, dists *[]int) {
	if dist > k {
		return // K를 초과하면 더 이상 탐색하지 않는다 (가지치기)
	}
	*dists = append(*dists, dist)
	for _, e := range adj[v] {
		if e.to == p || removed[e.to] {
			continue
		}
		getDists(e.to, v, dist+e.weight, dists)
	}
}

// countExact는 거리 배열에서 합이 정확히 K인 쌍의 수를 해시맵으로 센다
func countExact(dists []int) int64 {
	freq := make(map[int]int64)
	var cnt int64
	for _, d := range dists {
		need := k - d
		if need >= 0 {
			cnt += freq[need]
		}
		freq[d]++
	}
	return cnt
}

// decompose는 센트로이드 분할을 수행하며 쌍을 센다
func decompose(v int) {
	// 서브트리 크기 계산 및 센트로이드 찾기
	treeSize := calcSize(v, -1)
	centroid := findCentroid(v, -1, treeSize)
	removed[centroid] = true

	// 센트로이드를 지나는 경로 처리
	// 전체 거리 배열 (센트로이드 자신 = 거리 0)
	allDists := []int{0}

	for _, e := range adj[centroid] {
		if removed[e.to] {
			continue
		}
		// 각 서브트리의 거리 수집
		subDists := []int{}
		getDists(e.to, centroid, e.weight, &subDists)

		// 같은 서브트리 내의 쌍은 빼야 한다 (포함-배제)
		answer -= countExact(subDists)

		// 전체 배열에 추가
		allDists = append(allDists, subDists...)
	}

	// 전체에서 정확히 K인 쌍 세기
	answer += countExact(allDists)

	// 각 서브트리에 대해 재귀
	for _, e := range adj[centroid] {
		if removed[e.to] {
			continue
		}
		decompose(e.to)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수 N, 목표 거리 K
	fmt.Fscan(reader, &n, &k)

	// 입력: 가중치 간선 정보
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], Edge{v, w})
		adj[v] = append(adj[v], Edge{u, w})
	}

	// 센트로이드 분할로 쌍 세기
	answer = 0
	decompose(1)

	// 출력
	fmt.Fprintln(writer, answer)
}
