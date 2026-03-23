package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 100001
const LG = 17

type edgeInfo struct {
	to, w int
}

var graph [MX][]edgeInfo
var anc [MX][LG]int
var maxW [MX][LG]int
var dep [MX]int

// buildTree는 DFS로 깊이, 조상, 경로 최대 가중치를 구성한다.
func buildTree(v, par, d, w int) {
	dep[v] = d
	anc[v][0] = par
	maxW[v][0] = w
	for k := 1; k < LG; k++ {
		anc[v][k] = anc[anc[v][k-1]][k-1]
		if maxW[v][k-1] > maxW[anc[v][k-1]][k-1] {
			maxW[v][k] = maxW[v][k-1]
		} else {
			maxW[v][k] = maxW[anc[v][k-1]][k-1]
		}
	}
	for _, e := range graph[v] {
		if e.to != par {
			buildTree(e.to, v, d+1, e.w)
		}
	}
}

// queryPathMax는 트리에서 두 노드 사이 경로의 최대 간선 가중치를 구한다.
//
// [매개변수]
//   - u: 첫 번째 노드 번호
//   - v: 두 번째 노드 번호
//
// [반환값]
//   - int: u에서 v까지 경로 위 간선 가중치의 최댓값
func queryPathMax(u, v int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n-1; i++ {
		var a, b, w int
		fmt.Fscan(reader, &a, &b, &w)
		graph[a] = append(graph[a], edgeInfo{b, w})
		graph[b] = append(graph[b], edgeInfo{a, w})
	}

	buildTree(1, 0, 0, 0)

	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		fmt.Fprintln(writer, queryPathMax(u, v))
	}
}
