package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// 오일러 투어 + Mo's Algorithm으로 서브트리 색상 종류 쿼리를 처리한다.
//
// 핵심 아이디어:
//   오일러 투어로 서브트리를 연속 구간으로 변환한 뒤,
//   Mo's Algorithm을 적용하여 구간 내 서로 다른 색상 수를 구한다.
//   갱신 질의가 있으므로 Mo's Algorithm with Updates를 사용한다.
//
// 시간 복잡도: O((N + Q) * N^(2/3))

const MAXN = 50001

var (
	adj    [MAXN][]int
	in     [MAXN]int
	out    [MAXN]int
	euler  [MAXN]int // euler[i] = i번째 방문한 노드
	color  [MAXN]int // 현재 색상
	timer  int
	n, c   int
	q      int
	block  int
	cnt    [MAXN]int // 각 색상의 등장 횟수
	curAns int       // 현재 구간의 서로 다른 색상 수
)

// 질의 구조체
type query struct {
	l, r, t int // 구간 [l, r], 시간 t
	idx     int // 원래 질의 번호
}

// 갱신 구조체
type update struct {
	pos     int // 오일러 투어 배열에서의 위치
	oldCol  int // 이전 색상
	newCol  int // 새 색상
	nodeIdx int // 노드 번호
}

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

// 색상을 추가한다
func add(pos int) {
	col := color[euler[pos]]
	cnt[col]++
	if cnt[col] == 1 {
		curAns++
	}
}

// 색상을 제거한다
func remove(pos int) {
	col := color[euler[pos]]
	cnt[col]--
	if cnt[col] == 0 {
		curAns--
	}
}

// 시간 갱신을 적용한다
func applyUpdate(upd update, curL, curR int) {
	// 현재 구간에 포함된 위치라면 색상 교체 반영
	if curL <= upd.pos && upd.pos <= curR {
		remove(upd.pos)
		color[upd.nodeIdx] = upd.newCol
		add(upd.pos)
	} else {
		color[upd.nodeIdx] = upd.newCol
	}
}

// 시간 갱신을 되돌린다
func rollbackUpdate(upd update, curL, curR int) {
	if curL <= upd.pos && upd.pos <= curR {
		remove(upd.pos)
		color[upd.nodeIdx] = upd.oldCol
		add(upd.pos)
	} else {
		color[upd.nodeIdx] = upd.oldCol
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력
	fmt.Fscan(reader, &n, &c, &q)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &color[i])
	}

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 오일러 투어 수행
	timer = 0
	dfs(1, 0)

	// 블록 크기 설정 (Mo's with Updates: N^(2/3))
	block = int(math.Max(1, math.Cbrt(float64(n))))

	// 질의와 갱신 분리
	queries := make([]query, 0)
	updates := make([]update, 0)
	tCnt := 0 // 현재까지의 갱신 횟수

	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var v, newC int
			fmt.Fscan(reader, &v, &newC)
			updates = append(updates, update{
				pos:     in[v],
				oldCol:  color[v],
				newCol:  newC,
				nodeIdx: v,
			})
			color[v] = newC // 색상 갱신 (나중에 되돌릴 것)
			tCnt++
		} else {
			var v int
			fmt.Fscan(reader, &v)
			queries = append(queries, query{
				l:   in[v],
				r:   out[v],
				t:   tCnt,
				idx: len(queries),
			})
		}
	}

	// 모든 갱신을 되돌려서 초기 상태로 복원
	for i := len(updates) - 1; i >= 0; i-- {
		color[updates[i].nodeIdx] = updates[i].oldCol
	}

	// Mo's Algorithm 정렬
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

	// Mo's Algorithm 실행
	ans := make([]int, len(queries))
	curL, curR, curT := 0, -1, 0

	for _, qr := range queries {
		// 시간 갱신 적용/되돌리기
		for curT < qr.t {
			applyUpdate(updates[curT], curL, curR)
			curT++
		}
		for curT > qr.t {
			curT--
			rollbackUpdate(updates[curT], curL, curR)
		}

		// 구간 확장/축소
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

	// 결과 출력
	for _, a := range ans {
		fmt.Fprintln(writer, a)
	}
}
