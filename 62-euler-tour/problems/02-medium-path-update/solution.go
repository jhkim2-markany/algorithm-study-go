package main

import (
	"bufio"
	"fmt"
	"os"
)

// 오일러 투어 + 펜윅 트리(차분 배열)로 경로 갱신 / 노드 값 질의를 처리한다.
//
// 핵심 아이디어:
//   루트→v 경로에 x를 더하는 것은, 오일러 투어 배열에서
//   v의 서브트리 구간 [in[v], out[v]]에 x를 더하는 것과 같다.
//   (서브트리 내 모든 노드는 v를 조상으로 가지므로, v를 지나는 경로에 포함된다)
//
//   노드 u의 현재 값 = 조상 중 갱신된 값의 합
//                    = 평탄화 배열에서 [0, in[u]] 구간의 합 (차분 배열 기법)
//
// 시간 복잡도: 전처리 O(N), 갱신 O(log N), 질의 O(log N)

const MAXN = 100001

var (
	adj   [MAXN][]int
	in    [MAXN]int
	out   [MAXN]int
	bit   [MAXN + 1]int // 펜윅 트리 (차분 배열용)
	timer int
	n, q  int
)

// dfs는 오일러 투어를 수행한다
func dfs(v, parent int) {
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

// 펜윅 트리: i번 위치에 delta를 더한다 (0-indexed → 내부 1-indexed)
func update(i, delta int) {
	for i++; i <= n; i += i & (-i) {
		bit[i] += delta
	}
}

// 펜윅 트리: [0, i] 구간 합 (0-indexed)
func query(i int) int {
	sum := 0
	for i++; i > 0; i -= i & (-i) {
		sum += bit[i]
	}
	return sum
}

// rangeAdd는 [l, r] 구간에 delta를 더한다 (차분 배열 기법)
func rangeAdd(l, r, delta int) {
	update(l, delta)
	if r+1 <= n-1 {
		update(r+1, -delta)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수, 질의 수
	fmt.Fscan(reader, &n, &q)

	// 입력: 간선 정보
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 오일러 투어 수행
	timer = 0
	dfs(1, 0)

	// 질의 처리
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			// 루트→v 경로에 x를 더한다
			// = 오일러 투어 배열에서 v의 서브트리 구간에 x를 더한다
			var v, x int
			fmt.Fscan(reader, &v, &x)
			rangeAdd(in[v], out[v], x)
		} else {
			// 노드 v의 현재 값 = 평탄화 배열에서 in[v]까지의 누적 합
			var v int
			fmt.Fscan(reader, &v)
			fmt.Fprintln(writer, query(in[v]))
		}
	}
}
