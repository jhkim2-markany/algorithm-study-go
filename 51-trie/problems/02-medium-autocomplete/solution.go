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
func insertWord(root *TrieNode, word string) {
	// 여기에 코드를 작성하세요
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
func autocomplete(root *TrieNode, prefix string, k int) []string {
	// 여기에 코드를 작성하세요
	return nil
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
