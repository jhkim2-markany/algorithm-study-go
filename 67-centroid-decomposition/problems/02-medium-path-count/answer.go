package main

import (
	"bufio"
	"fmt"
	"os"
)

// countExactDistancePairs는 센트로이드 분할과 해시맵을 이용하여
// 트리에서 정확히 거리 K인 노드 쌍의 수를 구한다.
//
// [매개변수]
//   - n: 노드 수
//   - k: 목표 거리
//   - edges: 가중치 간선 목록 (u, v, w)
//
// [반환값]
//   - int64: 정확히 거리 K인 노드 쌍의 수
//
// [알고리즘 힌트]
//   1. 센트로이드를 찾아 제거하고, 센트로이드를 지나는 경로를 처리한다
//   2. 각 서브트리에서 센트로이드까지의 가중치 거리를 수집한다 (K 초과 시 가지치기)
//   3. 해시맵으로 거리 합이 정확히 K인 쌍을 센다
//   4. 같은 서브트리 내의 쌍은 포함-배제로 빼준다
//   5. 각 서브트리에 대해 재귀적으로 분할한다
func countExactDistancePairs(n, k int, edges [][3]int) int64 {
	type Edge struct {
		to, weight int
	}
	adj := make([][]Edge, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], Edge{e[0], e[2]})
	}

	sz := make([]int, n+1)
	removed := make([]bool, n+1)
	var answer int64

	var calcSize func(v, p int) int
	calcSize = func(v, p int) int {
		sz[v] = 1
		for _, e := range adj[v] {
			if e.to == p || removed[e.to] {
				continue
			}
			sz[v] += calcSize(e.to, v)
		}
		return sz[v]
	}

	var findCentroid func(v, p, treeSize int) int
	findCentroid = func(v, p, treeSize int) int {
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

	var getDists func(v, p, dist int, dists *[]int)
	getDists = func(v, p, dist int, dists *[]int) {
		if dist > k {
			return
		}
		*dists = append(*dists, dist)
		for _, e := range adj[v] {
			if e.to == p || removed[e.to] {
				continue
			}
			getDists(e.to, v, dist+e.weight, dists)
		}
	}

	countExact := func(dists []int) int64 {
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

	var decompose func(v int)
	decompose = func(v int) {
		treeSize := calcSize(v, -1)
		centroid := findCentroid(v, -1, treeSize)
		removed[centroid] = true

		allDists := []int{0}
		for _, e := range adj[centroid] {
			if removed[e.to] {
				continue
			}
			subDists := []int{}
			getDists(e.to, centroid, e.weight, &subDists)
			answer -= countExact(subDists)
			allDists = append(allDists, subDists...)
		}
		answer += countExact(allDists)

		for _, e := range adj[centroid] {
			if removed[e.to] {
				continue
			}
			decompose(e.to)
		}
	}

	decompose(1)
	return answer
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	edges := make([][3]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	fmt.Fprintln(writer, countExactDistancePairs(n, k, edges))
}
