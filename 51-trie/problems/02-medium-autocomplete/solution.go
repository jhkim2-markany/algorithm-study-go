package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 트라이 노드 구조체
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool // 단어의 끝인지 여부
}

// 트라이에 단어를 삽입한다
func insert(root *TrieNode, word string) {
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

// 접두사 노드까지 이동한다. 경로가 없으면 nil을 반환한다.
func findPrefixNode(root *TrieNode, prefix string) *TrieNode {
	node := root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}

// DFS로 접두사 이후의 단어를 사전 순으로 최대 k개 수집한다
func collectWords(node *TrieNode, prefix string, k int, results *[]string) {
	if len(*results) >= k {
		return // 이미 k개를 수집했으면 종료
	}
	if node.isEnd {
		*results = append(*results, prefix)
	}
	// 알파벳 순서(a~z)로 자식을 탐색하면 사전 순이 보장된다
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			collectWords(node.children[i], prefix+string(rune('a'+i)), k, results)
			if len(*results) >= k {
				return
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 단어 수 N, 질의 수 M, 최대 추천 개수 K
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	// 트라이 생성 및 단어 삽입
	root := &TrieNode{}
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(reader, &word)
		insert(root, word)
	}

	// 각 질의에 대해 자동완성 결과 출력
	for i := 0; i < m; i++ {
		var query string
		fmt.Fscan(reader, &query)

		// 접두사 노드를 찾는다
		prefixNode := findPrefixNode(root, query)
		if prefixNode == nil {
			fmt.Fprintln(writer, "-1")
			continue
		}

		// DFS로 사전 순 최대 K개 단어를 수집한다
		var results []string
		collectWords(prefixNode, query, k, &results)

		if len(results) == 0 {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, strings.Join(results, " "))
		}
	}
}
