package main

import (
	"bufio"
	"fmt"
	"os"
)

// 색칠된 경로 쿼리
// 센트로이드 분할로 각 거리별 같은 색상 쌍의 수를 전처리한다.
// 쿼리마다 O(1)로 답을 구한다.

const MAXN = 50001

var (
	adj     [MAXN][]int // 인접 리스트
	color   [MAXN]int   // 각 노드의 색상
	sz      [MAXN]int   // 서브트리 크기
	removed [MAXN]bool  // 센트로이드 제거 표시
	ans     [MAXN]int64 // ans[d] = 거리가 d이고 같은 색상인 쌍의 수
	n, q    int
)

// NodeInfo는 센트로이드에서의 거리와 색상 정보를 저장한다
type NodeInfo struct {
	dist  int
	color int
}

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

// getNodes는 센트로이드에서 서브트리 내 모든 노드의 거리와 색상을 수집한다
func getNodes(v, p, dist int, nodes *[]NodeInfo) {
	*nodes = append(*nodes, NodeInfo{dist, color[v]})
	for _, u := range adj[v] {
		if u == p || removed[u] {
			continue
		}
		getNodes(u, v, dist+1, nodes)
	}
}

// countColorPairs는 노드 정보 배열에서 같은 색상이고 거리 합이 d인 쌍의 수를
// 거리별로 ans 배열에 더하거나 빼준다 (sign = +1 또는 -1)
func countColorPairs(nodes []NodeInfo, sign int64) {
	// colorDist[c] = 색상 c인 노드들의 거리 목록
	colorDist := make(map[int][]int)
	for _, nd := range nodes {
		colorDist[nd.color] = append(colorDist[nd.color], nd.dist)
	}

	// 같은 색상 내에서 거리 합별 쌍 수를 센다
	// distCount[d] = 거리 합이 d인 같은 색상 쌍의 수
	distCount := make(map[int]int64)
	for _, dists := range colorDist {
		if len(dists) < 2 {
			continue
		}
		// 같은 색상의 노드들 사이에서 모든 쌍의 거리 합을 구한다
		// freq[d] = 이 색상에서 거리가 d인 노드 수
		freq := make(map[int]int64)
		for _, d := range dists {
			freq[d]++
		}
		// 모든 쌍의 거리 합 계산
		for i := 0; i < len(dists); i++ {
			for j := i + 1; j < len(dists); j++ {
				distCount[dists[i]+dists[j]]++
			}
		}
	}

	// ans 배열에 반영
	for d, cnt := range distCount {
		if d < MAXN {
			ans[d] += sign * cnt
		}
	}
}

// decompose는 센트로이드 분할을 수행하며 거리별 같은 색상 쌍을 센다
func decompose(v int) {
	treeSize := calcSize(v, -1)
	centroid := findCentroid(v, -1, treeSize)
	removed[centroid] = true

	// 센트로이드를 포함한 전체 노드 정보
	allNodes := []NodeInfo{{0, color[centroid]}}

	for _, u := range adj[centroid] {
		if removed[u] {
			continue
		}
		subNodes := []NodeInfo{}
		getNodes(u, centroid, 1, &subNodes)

		// 같은 서브트리 내의 쌍은 빼야 한다 (포함-배제)
		countColorPairs(subNodes, -1)

		allNodes = append(allNodes, subNodes...)
	}

	// 전체에서 같은 색상 쌍 세기
	countColorPairs(allNodes, 1)

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

	// 입력: 노드 수 N, 쿼리 수 Q
	fmt.Fscan(reader, &n, &q)

	// 입력: 각 노드의 색상
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &color[i])
	}

	// 입력: 간선 정보
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 센트로이드 분할로 거리별 같은 색상 쌍 전처리
	decompose(1)

	// 쿼리 처리
	for i := 0; i < q; i++ {
		var k int
		fmt.Fscan(reader, &k)
		fmt.Fprintln(writer, ans[k])
	}
}
