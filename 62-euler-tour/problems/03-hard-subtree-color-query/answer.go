package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// subtreeColorQuery는 오일러 투어와 Mo's Algorithm(with Updates)을 이용하여
// 서브트리 내 서로 다른 색상 수 질의를 처리한다.
//
// [매개변수]
//   - n: 노드 수
//   - c: 색상 종류 수
//   - initColor: 각 노드의 초기 색상 (1-indexed)
//   - edges: 간선 목록 (u, v 쌍)
//   - ops: 연산 목록 (타입1: 색상 변경, 타입2: 서브트리 색상 수 질의)
//
// [반환값]
//   - []int: 서브트리 색상 수 질의(타입 2)의 결과 배열
//
// [알고리즘 힌트]
//   1. 오일러 투어로 서브트리를 연속 구간 [in[v], out[v]]로 변환한다
//   2. 질의와 갱신을 분리하고, Mo's Algorithm with Updates를 적용한다
//   3. 블록 크기를 N^(2/3)으로 설정하여 (l/block, r/block, t) 순으로 정렬한다
//   4. 색상 추가/제거 시 빈도 배열로 서로 다른 색상 수를 관리한다
//   5. 시간 갱신 시 현재 구간에 포함된 위치면 색상 교체를 반영한다
func subtreeColorQuery(n, c int, initColor []int, edges [][2]int, ops [][]int) []int {
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

	color := make([]int, n+1)
	copy(color, initColor)

	block := int(math.Max(1, math.Cbrt(float64(n))))

	type query struct {
		l, r, t, idx int
	}
	type update struct {
		pos, oldCol, newCol, nodeIdx int
	}

	var queries []query
	var updates []update
	tCnt := 0

	for _, op := range ops {
		if op[0] == 1 {
			v, newC := op[1], op[2]
			updates = append(updates, update{in[v], color[v], newC, v})
			color[v] = newC
			tCnt++
		} else {
			v := op[1]
			queries = append(queries, query{in[v], out[v], tCnt, len(queries)})
		}
	}

	// 초기 상태로 복원
	for i := len(updates) - 1; i >= 0; i-- {
		color[updates[i].nodeIdx] = updates[i].oldCol
	}

	sort.Slice(queries, func(i, j int) bool {
		bi, bj := queries[i].l/block, queries[j].l/block
		if bi != bj {
			return bi < bj
		}
		ri, rj := queries[i].r/block, queries[j].r/block
		if ri != rj {
			return ri < rj
		}
		return queries[i].t < queries[j].t
	})

	cnt := make([]int, c+1)
	curAns := 0

	add := func(pos int) {
		col := color[euler[pos]]
		cnt[col]++
		if cnt[col] == 1 {
			curAns++
		}
	}
	remove := func(pos int) {
		col := color[euler[pos]]
		cnt[col]--
		if cnt[col] == 0 {
			curAns--
		}
	}

	ans := make([]int, len(queries))
	curL, curR, curT := 0, -1, 0

	for _, qr := range queries {
		for curT < qr.t {
			upd := updates[curT]
			if curL <= upd.pos && upd.pos <= curR {
				remove(upd.pos)
				color[upd.nodeIdx] = upd.newCol
				add(upd.pos)
			} else {
				color[upd.nodeIdx] = upd.newCol
			}
			curT++
		}
		for curT > qr.t {
			curT--
			upd := updates[curT]
			if curL <= upd.pos && upd.pos <= curR {
				remove(upd.pos)
				color[upd.nodeIdx] = upd.oldCol
				add(upd.pos)
			} else {
				color[upd.nodeIdx] = upd.oldCol
			}
		}
		for curR < qr.r {
			curR++
			add(curR)
		}
		for curL > qr.l {
			curL--
			add(curL)
		}
		for curR > qr.r {
			remove(curR)
			curR--
		}
		for curL < qr.l {
			remove(curL)
			curL++
		}
		ans[qr.idx] = curAns
	}

	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, c, q int
	fmt.Fscan(reader, &n, &c, &q)

	initColor := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &initColor[i])
	}

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	ops := make([][]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var v, newC int
			fmt.Fscan(reader, &v, &newC)
			ops[i] = []int{t, v, newC}
		} else {
			var v int
			fmt.Fscan(reader, &v)
			ops[i] = []int{t, v}
		}
	}

	results := subtreeColorQuery(n, c, initColor, edges, ops)
	for _, r := range results {
		fmt.Fprintln(writer, r)
	}
}
