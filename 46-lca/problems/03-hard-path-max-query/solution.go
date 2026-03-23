package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 100001
const LG = 17

// 간선 정보
type edgeInfo struct {
	to, w int
}

var graph [MX][]edgeInfo
var anc [MX][LG]int  // anc[v][k] = v의 2^k번째 조상
var maxW [MX][LG]int // maxW[v][k] = v에서 2^k번째 조상까지 경로의 최대 가중치
var dep [MX]int

// DFS로 깊이, 조상, 경로 최대 가중치를 구성한다
func dfs(v, par, d, w int) {
	dep[v] = d
	anc[v][0] = par
	maxW[v][0] = w
	for k := 1; k < LG; k++ {
		anc[v][k] = anc[anc[v][k-1]][k-1]
		// 두 구간의 최대 가중치 중 큰 값을 취한다
		if maxW[v][k-1] > maxW[anc[v][k-1]][k-1] {
			maxW[v][k] = maxW[v][k-1]
		} else {
			maxW[v][k] = maxW[anc[v][k-1]][k-1]
		}
	}
	for _, e := range graph[v] {
		if e.to != par {
			dfs(e.to, v, d+1, e.w)
		}
	}
}

// 경로 위 최대 가중치를 구하며 LCA까지 올린다
func queryPathMax(u, v int) int {
	result := 0

	// u가 더 깊도록 보장한다
	if dep[u] < dep[v] {
		u, v = v, u
	}

	// 깊이를 맞추면서 최대 가중치를 갱신한다
	diff := dep[u] - dep[v]
	for k := 0; k < LG; k++ {
		if (diff>>k)&1 == 1 {
			if maxW[u][k] > result {
				result = maxW[u][k]
			}
			u = anc[u][k]
		}
	}

	if u == v {
		return result
	}

	// LCA 바로 아래까지 올리면서 최대 가중치를 갱신한다
	for k := LG - 1; k >= 0; k-- {
		if anc[u][k] != anc[v][k] {
			if maxW[u][k] > result {
				result = maxW[u][k]
			}
			if maxW[v][k] > result {
				result = maxW[v][k]
			}
			u = anc[u][k]
			v = anc[v][k]
		}
	}

	// 마지막 한 칸 (LCA까지)의 가중치를 확인한다
	if maxW[u][0] > result {
		result = maxW[u][0]
	}
	if maxW[v][0] > result {
		result = maxW[v][0]
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 간선 입력 (가중치 포함)
	for i := 0; i < n-1; i++ {
		var a, b, w int
		fmt.Fscan(reader, &a, &b, &w)
		graph[a] = append(graph[a], edgeInfo{b, w})
		graph[b] = append(graph[b], edgeInfo{a, w})
	}

	// 전처리: DFS로 깊이, 조상, 경로 최대 가중치 구성
	dfs(1, 0, 0, 0)

	// 쿼리 처리
	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		// 경로 위 최대 가중치를 출력한다
		fmt.Fprintln(writer, queryPathMax(u, v))
	}
}
