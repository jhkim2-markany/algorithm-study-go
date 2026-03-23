package main

import (
	"bufio"
	"fmt"
	"os"
)

// 오일러 투어 + 펜윅 트리로 서브트리 합 질의를 처리한다
// 시간 복잡도: 전처리 O(N log N), 질의/갱신 O(log N)

const MAXN = 100001

var (
	adj   [MAXN][]int // 인접 리스트
	in    [MAXN]int   // 방문 시작 시각
	out   [MAXN]int   // 방문 종료 시각
	euler [MAXN]int   // 오일러 투어 순서
	bit   [MAXN]int   // 펜윅 트리 (Binary Indexed Tree)
	val   [MAXN]int   // 노드 값
	timer int
	n, q  int
)

// dfs는 오일러 투어를 수행한다
func dfs(v, parent int) {
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

// 펜윅 트리: i번 위치에 delta를 더한다 (1-indexed)
func update(i, delta int) {
	for i++; i <= n; i += i & (-i) {
		bit[i] += delta
	}
}

// 펜윅 트리: [0, i] 구간 합을 반환한다 (0-indexed → 내부 1-indexed)
func query(i int) int {
	sum := 0
	for i++; i > 0; i -= i & (-i) {
		sum += bit[i]
	}
	return sum
}

// rangeQuery는 [l, r] 구간 합을 반환한다
func rangeQuery(l, r int) int {
	if l == 0 {
		return query(r)
	}
	return query(r) - query(l-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 노드 수, 질의 수
	fmt.Fscan(reader, &n, &q)

	// 입력: 노드 값
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &val[i])
	}

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

	// 펜윅 트리 초기화: 평탄화 배열에 노드 값 삽입
	for i := 0; i < n; i++ {
		update(i, val[euler[i]])
	}

	// 질의 처리
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			// 노드 v의 값을 x로 변경
			var v, x int
			fmt.Fscan(reader, &v, &x)
			// 기존 값과의 차이만큼 갱신
			update(in[v], x-val[v])
			val[v] = x
		} else {
			// 노드 v의 서브트리 합 출력
			var v int
			fmt.Fscan(reader, &v)
			fmt.Fprintln(writer, rangeQuery(in[v], out[v]))
		}
	}
}
