package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 거리가 K 이하인 노드 쌍 세기
// 센트로이드 분할 + 정렬 + 투 포인터로 해결한다.

const MAXN = 100001

var (
	adj     [MAXN][]int // 인접 리스트
	sz      [MAXN]int   // 서브트리 크기
	removed [MAXN]bool  // 센트로이드 제거 표시
	n, k    int
	answer  int64
)

// calcSize는 현재 서브트리의 크기를 계산한다
func calcSize(v, p int) int {
	sz[v] = 1
	for _, u := range adj[v] {
		if u == p || removed[u] {
			continue
		}
		sz[v] += calcSize(u, v)
	}
	return sz[v]
}

// findCentroid는 센트로이드를 찾는다
func findCentroid(v, p, treeSize int) int {
	for _, u := range adj[v] {
		if u == p || removed[u] {
			continue
		}
		if sz[u] > treeSize/2 {
			return findCentroid(u, v, treeSize)
		}
	}
	return v
}

// getDists는 센트로이드에서 서브트리 내 모든 노드까지의 거리를 수집한다
func getDists(v, p, dist int, dists *[]int) {
	*dists = append(*dists, dist)
	for _, u := range adj[v] {
		if u == p || removed[u] {
			continue
		}
		getDists(u, v, dist+1, dists)
	}
}

// countPairs는 정렬된 거리 배열에서 합이 K 이하인 쌍의 수를 센다
func countPairs(dists []int) int64 {
	sort.Ints(dists)
	var cnt int64
	lo, hi := 0, len(dists)-1
	for lo < hi {
		if dists[lo]+dists[hi] <= k {
			// lo와 lo+1, lo+2, ..., hi 모두 조건 만족
			cnt += int64(hi - lo)
			lo++
		} else {
			hi--
		}
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
	// 전체 거리 배열 (센트로이드 자신 포함)
	allDists := []int{0}

	for _, u := range adj[centroid] {
		if removed[u] {
			continue
		}
		// 각 서브트리의 거리 수집
		subDists := []int{}
		getDists(u, centroid, 1, &subDists)

		// 같은 서브트리 내의 쌍은 빼야 한다 (포함-배제)
		answer -= countPairs(subDists)

		// 전체 배열에 추가
		allDists = append(allDists, subDists...)
	}

	// 전체에서 K 이하인 쌍 세기
	answer += countPairs(allDists)

	// 각 서브트리에 대해 재귀
	for _, u := range adj[centroid] {
		if removed[u] {
			continue
		}
		decompose(u)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수 N, 거리 제한 K
	fmt.Fscan(reader, &n, &k)

	// 입력: 간선 정보
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 센트로이드 분할로 쌍 세기
	answer = 0
	decompose(1)

	// 출력
	fmt.Fprintln(writer, answer)
}
