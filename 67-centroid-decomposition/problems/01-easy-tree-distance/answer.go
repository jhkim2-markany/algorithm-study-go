package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// countPairsWithinK는 센트로이드 분할을 이용하여 트리에서
// 거리가 K 이하인 노드 쌍의 수를 구한다.
//
// [매개변수]
//   - n: 노드 수
//   - k: 거리 제한
//   - edges: 간선 목록 (u, v 쌍)
//
// [반환값]
//   - int64: 거리가 K 이하인 노드 쌍의 수
//
// [알고리즘 힌트]
//   1. 센트로이드를 찾아 제거하고, 센트로이드를 지나는 경로를 처리한다
//   2. 각 서브트리에서 센트로이드까지의 거리를 수집한다
//   3. 전체 거리 배열에서 합이 K 이하인 쌍을 투 포인터로 센다
//   4. 같은 서브트리 내의 쌍은 포함-배제로 빼준다
//   5. 각 서브트리에 대해 재귀적으로 분할한다
func countPairsWithinK(n, k int, edges [][2]int) int64 {
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	sz := make([]int, n+1)
	removed := make([]bool, n+1)
	var answer int64

	var calcSize func(v, p int) int
	calcSize = func(v, p int) int {
		sz[v] = 1
		for _, u := range adj[v] {
			if u == p || removed[u] {
				continue
			}
			sz[v] += calcSize(u, v)
		}
		return sz[v]
	}

	var findCentroid func(v, p, treeSize int) int
	findCentroid = func(v, p, treeSize int) int {
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

	var getDists func(v, p, dist int, dists *[]int)
	getDists = func(v, p, dist int, dists *[]int) {
		*dists = append(*dists, dist)
		for _, u := range adj[v] {
			if u == p || removed[u] {
				continue
			}
			getDists(u, v, dist+1, dists)
		}
	}

	countPairs := func(dists []int) int64 {
		sort.Ints(dists)
		var cnt int64
		lo, hi := 0, len(dists)-1
		for lo < hi {
			if dists[lo]+dists[hi] <= k {
				cnt += int64(hi - lo)
				lo++
			} else {
				hi--
			}
		}
		return cnt
	}

	var decompose func(v int)
	decompose = func(v int) {
		treeSize := calcSize(v, -1)
		centroid := findCentroid(v, -1, treeSize)
		removed[centroid] = true

		allDists := []int{0}
		for _, u := range adj[centroid] {
			if removed[u] {
				continue
			}
			subDists := []int{}
			getDists(u, centroid, 1, &subDists)
			answer -= countPairs(subDists)
			allDists = append(allDists, subDists...)
		}
		answer += countPairs(allDists)

		for _, u := range adj[centroid] {
			if removed[u] {
				continue
			}
			decompose(u)
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

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	fmt.Fprintln(writer, countPairsWithinK(n, k, edges))
}
