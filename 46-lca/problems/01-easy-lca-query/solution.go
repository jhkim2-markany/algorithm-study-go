package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxN = 100001
const logN = 17

var tree [maxN][]int
var up [maxN][logN]int
var dep [maxN]int

// buildLCA는 DFS로 깊이와 희소 배열을 구성한다.
func buildLCA(v, par, d int) {
	dep[v] = d
	up[v][0] = par
	for k := 1; k < logN; k++ {
		up[v][k] = up[up[v][k-1]][k-1]
	}
	for _, u := range tree[v] {
		if u != par {
			buildLCA(u, v, d+1)
		}
	}
}

// lca는 Binary Lifting으로 두 노드의 최소 공통 조상을 구한다.
//
// [매개변수]
//   - u: 첫 번째 노드 번호
//   - v: 두 번째 노드 번호
//
// [반환값]
//   - int: u와 v의 최소 공통 조상 (LCA) 노드 번호
func lca(u, v int) int {
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
		var a, b int
		fmt.Fscan(reader, &a, &b)
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}

	buildLCA(1, 0, 0)

	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		fmt.Fprintln(writer, lca(u, v))
	}
}
