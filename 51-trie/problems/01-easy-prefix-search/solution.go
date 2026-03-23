package main

import (
	"bufio"
	"fmt"
	"os"
)

// 트라이 노드 구조체
type TrieNode struct {
	children [26]*TrieNode
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
}

// 주어진 문자열이 트라이에 저장된 단어의 접두사인지 확인한다
func hasPrefix(root *TrieNode, prefix string) bool {
	node := root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 단어 수 N, 질의 수 M
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 트라이 생성 및 단어 삽입
	root := &TrieNode{}
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(reader, &word)
		insert(root, word)
	}

	// 각 질의에 대해 접두사 존재 여부 판별
	for i := 0; i < m; i++ {
		var query string
		fmt.Fscan(reader, &query)
		if hasPrefix(root, query) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
