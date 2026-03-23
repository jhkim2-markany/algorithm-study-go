package main

import (
	"bufio"
	"fmt"
	"os"
)

// subtreeSum은 오일러 투어와 펜윅 트리를 이용하여 트리에서
// 서브트리 합 질의와 노드 값 갱신을 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - val: 각 노드의 초기 값 (1-indexed)
//   - edges: 간선 목록 (u, v 쌍)
//   - queries: 질의 목록 (타입, 인자들)
//
// [반환값]
//   - []int: 서브트리 합 질의(타입 2)의 결과 배열
//
// [알고리즘 힌트]
//   1. 오일러 투어로 각 노드의 방문 시작/종료 시각(in, out)을 구한다
//   2. 서브트리는 평탄화 배열에서 [in[v], out[v]] 연속 구간이 된다
//   3. 펜윅 트리로 구간 합 질의와 점 갱신을 O(log N)에 처리한다
//   4. 노드 값 변경 시 기존 값과의 차이만큼 펜윅 트리를 갱신한다
func subtreeSum(n int, val []int, edges [][2]int, queries [][]int) []int {
	adj := make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	in := make([]int, n+1)
	out := make([]int, n+1)
	euler := make([]int, n)
	timer := 0

	var dfs func(v, parent int)
	dfs = func(v, parent int) {
		in[v] = timer
		euler[timer] = v
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
	rangeQuery := func(l, r int) int {
		if l == 0 {
			return bitQuery(r)
		}
		return bitQuery(r) - bitQuery(l-1)
	}

	for i := 0; i < n; i++ {
		bitUpdate(i, val[euler[i]])
	}

	var results []int
	for _, q := range queries {
		if q[0] == 1 {
			v, x := q[1], q[2]
			bitUpdate(in[v], x-val[v])
			val[v] = x
		} else {
			v := q[1]
			results = append(results, rangeQuery(in[v], out[v]))
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

	val := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &val[i])
	}

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

	results := subtreeSum(n, val, edges, queries)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
