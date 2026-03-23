package main

import (
	"bufio"
	"fmt"
	"os"
)

// pathUpdate는 오일러 투어와 펜윅 트리(차분 배열)를 이용하여
// 루트→v 경로 갱신과 노드 값 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - edges: 간선 목록 (u, v 쌍)
//   - queries: 질의 목록 (타입, 인자들)
//
// [반환값]
//   - []int: 노드 값 질의(타입 2)의 결과 배열
//
// [알고리즘 힌트]
//   1. 오일러 투어로 각 노드의 in/out 시각을 구한다
//   2. 루트→v 경로에 x를 더하는 것은 평탄화 배열의 [in[v], out[v]]에 x를 더하는 것과 같다
//   3. 노드 u의 현재 값은 평탄화 배열에서 [0, in[u]]까지의 누적 합이다
//   4. 펜윅 트리에 차분 배열 기법을 적용하여 구간 갱신/점 질의를 O(log N)에 처리한다
func pathUpdate(n int, edges [][2]int, queries [][]int) []int {
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	in := make([]int, n+1)
	out := make([]int, n+1)
	timer := 0

	var dfs func(v, parent int)
	dfs = func(v, parent int) {
		in[v] = timer
		timer++
		for _, u := range adj[v] {
			if u == parent {
				continue
			}
			dfs(u, v)
		}
		out[v] = timer - 1
	}
	dfs(1, 0)

	bit := make([]int, n+2)
	bitUpdate := func(i, delta int) {
		for i++; i <= n; i += i & (-i) {
			bit[i] += delta
		}
	}
	bitQuery := func(i int) int {
		sum := 0
		for i++; i > 0; i -= i & (-i) {
			sum += bit[i]
		}
		return sum
	}
	rangeAdd := func(l, r, delta int) {
		bitUpdate(l, delta)
		if r+1 <= n-1 {
			bitUpdate(r+1, -delta)
		}
	}

	var results []int
	for _, q := range queries {
		if q[0] == 1 {
			v, x := q[1], q[2]
			rangeAdd(in[v], out[v], x)
		} else {
			v := q[1]
			results = append(results, bitQuery(in[v]))
		}
	}
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var v, x int
			fmt.Fscan(reader, &v, &x)
			queries[i] = []int{t, v, x}
		} else {
			var v int
			fmt.Fscan(reader, &v)
			queries[i] = []int{t, v}
		}
	}

	results := pathUpdate(n, edges, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
