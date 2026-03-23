package main

import "fmt"

// 트라이(Trie) - 접두사 트리 기본 구현
// 시간 복잡도: 삽입/검색/접두사 검색 O(L) (L: 문자열 길이)
// 공간 복잡도: O(N × L × 26) (N: 문자열 개수)

// TrieNode는 트라이의 각 노드를 나타낸다
type TrieNode struct {
	children [26]*TrieNode // 알파벳 소문자 a~z에 대한 자식 노드
	isEnd    bool          // 이 노드에서 단어가 끝나는지 여부
}

// Trie는 트라이 자료구조이다
type Trie struct {
	root *TrieNode
}

// NewTrie는 빈 트라이를 생성한다
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

// Insert는 트라이에 단어를 삽입한다
func (t *Trie) Insert(word string) {
	node := t.root
	// 단어의 각 문자를 따라가며 노드를 생성한다
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	// 마지막 노드에 단어 끝 표시
	node.isEnd = true
}

// Search는 트라이에 단어가 존재하는지 확인한다
func (t *Trie) Search(word string) bool {
	node := t.root
	// 단어의 각 문자를 따라간다
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false // 경로가 끊기면 단어가 없다
		}
		node = node.children[idx]
	}
	// 마지막 노드가 단어의 끝인지 확인한다
	return node.isEnd
}

// StartsWith는 주어진 접두사로 시작하는 단어가 있는지 확인한다
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	// 접두사의 각 문자를 따라간다
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false // 경로가 끊기면 해당 접두사가 없다
		}
		node = node.children[idx]
	}
	// 접두사 경로가 존재하면 true (단어 끝 여부는 무관)
	return true
}

// CollectWords는 주어진 접두사로 시작하는 모든 단어를 수집한다
func (t *Trie) CollectWords(prefix string) []string {
	node := t.root
	// 접두사 노드까지 이동한다
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	// DFS로 접두사 이후의 모든 단어를 수집한다
	var results []string
	var dfs func(n *TrieNode, path string)
	dfs = func(n *TrieNode, path string) {
		if n.isEnd {
			results = append(results, path)
		}
		for i := 0; i < 26; i++ {
			if n.children[i] != nil {
				dfs(n.children[i], path+string(rune('a'+i)))
			}
		}
	}
	dfs(node, prefix)
	return results
}

func main() {
	trie := NewTrie()

	// 단어 삽입
	words := []string{"apple", "app", "apt", "bat", "bad", "banana"}
	for _, w := range words {
		trie.Insert(w)
		fmt.Printf("삽입: %s\n", w)
	}
	fmt.Println()

	// 단어 검색
	searchWords := []string{"app", "apple", "ap", "bat", "ban"}
	fmt.Println("=== 단어 검색 ===")
	for _, w := range searchWords {
		fmt.Printf("Search(\"%s\") = %v\n", w, trie.Search(w))
	}
	fmt.Println()

	// 접두사 검색
	prefixes := []string{"ap", "ba", "cat", "app"}
	fmt.Println("=== 접두사 검색 ===")
	for _, p := range prefixes {
		fmt.Printf("StartsWith(\"%s\") = %v\n", p, trie.StartsWith(p))
	}
	fmt.Println()

	// 접두사로 시작하는 모든 단어 수집
	fmt.Println("=== 접두사로 시작하는 단어 수집 ===")
	collectPrefixes := []string{"ap", "ba", "b"}
	for _, p := range collectPrefixes {
		result := trie.CollectWords(p)
		fmt.Printf("CollectWords(\"%s\") = %v\n", p, result)
	}
}
