package main

import (
	"bufio"
	"fmt"
	"os"
)

// TrieNode는 트라이 노드 구조체이다.
type TrieNode struct {
	children [26]*TrieNode
}

// insertWord는 트라이에 단어를 삽입한다.
//
// [매개변수]
//   - root: 트라이의 루트 노드
//   - word: 삽입할 단어 (소문자 알파벳)
func insertWord(root *TrieNode, word string) {
	// 여기에 코드를 작성하세요
}

// hasPrefix는 주어진 문자열이 트라이에 저장된 단어의 접두사인지 확인한다.
//
// [매개변수]
//   - root: 트라이의 루트 노드
//   - prefix: 확인할 접두사 문자열
//
// [반환값]
//   - bool: 접두사가 존재하면 true, 아니면 false
func hasPrefix(root *TrieNode, prefix string) bool {
	// 여기에 코드를 작성하세요
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	root := &TrieNode{}
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(reader, &word)
		insertWord(root, word)
	}

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
