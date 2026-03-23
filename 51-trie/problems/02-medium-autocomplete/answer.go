package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TrieNode는 트라이 노드 구조체이다.
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

// insertWord는 트라이에 단어를 삽입한다.
//
// [매개변수]
//   - root: 트라이의 루트 노드
//   - word: 삽입할 단어 (소문자 알파벳)
//
// [알고리즘 힌트]
//
//	각 문자에 대해 자식 노드가 없으면 새로 생성하며, 마지막 노드에 isEnd를 표시한다.
func insertWord(root *TrieNode, word string) {
	node := root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// autocomplete는 주어진 접두사로 시작하는 단어를 사전 순으로 최대 k개 반환한다.
//
// [매개변수]
//   - root: 트라이의 루트 노드
//   - prefix: 검색할 접두사
//   - k: 최대 추천 개수
//
// [반환값]
//   - []string: 접두사로 시작하는 단어 목록 (사전 순, 최대 k개)
//
// [알고리즘 힌트]
//
//	접두사 노드까지 이동한 뒤, DFS로 알파벳 순서(a~z)로 자식을 탐색하면 사전 순이 보장된다.
//	k개를 수집하면 즉시 종료한다.
func autocomplete(root *TrieNode, prefix string, k int) []string {
	// 접두사 노드까지 이동
	node := root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}

	// DFS로 사전 순 최대 k개 수집
	var results []string
	var dfs func(nd *TrieNode, cur string)
	dfs = func(nd *TrieNode, cur string) {
		if len(results) >= k {
			return
		}
		if nd.isEnd {
			results = append(results, cur)
		}
		for i := 0; i < 26; i++ {
			if nd.children[i] != nil {
				dfs(nd.children[i], cur+string(rune('a'+i)))
				if len(results) >= k {
					return
				}
			}
		}
	}
	dfs(node, prefix)
	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	root := &TrieNode{}
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(reader, &word)
		insertWord(root, word)
	}

	for i := 0; i < m; i++ {
		var query string
		fmt.Fscan(reader, &query)
		results := autocomplete(root, query, k)
		if len(results) == 0 {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, strings.Join(results, " "))
		}
	}
}
