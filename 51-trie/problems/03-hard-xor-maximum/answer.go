package main

import (
	"bufio"
	"fmt"
	"os"
)

// BitTrieNode는 비트 트라이 노드 구조체이다.
type BitTrieNode struct {
	children [2]*BitTrieNode
}

// insertBitTrie는 비트 트라이에 숫자를 삽입한다.
//
// [매개변수]
//   - root: 비트 트라이의 루트 노드
//   - num: 삽입할 정수
//
// [알고리즘 힌트]
//
//	최상위 비트(30번째)부터 최하위 비트(0번째)까지 순서대로 삽입한다.
func insertBitTrie(root *BitTrieNode, num int) {
	node := root
	for i := 30; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &BitTrieNode{}
		}
		node = node.children[bit]
	}
}

// findMaxXOR는 주어진 숫자와 XOR 결과가 최대가 되는 값을 트라이에서 찾는다.
//
// [매개변수]
//   - root: 비트 트라이의 루트 노드
//   - num: XOR 대상 정수
//
// [반환값]
//   - int: num과 XOR했을 때 최대가 되는 XOR 값
//
// [알고리즘 힌트]
//
//	각 비트에서 현재 비트와 반대되는 경로를 우선 선택하여 XOR을 최대화한다.
//	반대 비트 경로가 없으면 같은 비트 경로로 진행한다.
func findMaxXOR(root *BitTrieNode, num int) int {
	node := root
	result := 0
	for i := 30; i >= 0; i-- {
		bit := (num >> i) & 1
		want := 1 - bit
		if node.children[want] != nil {
			result |= (1 << i)
			node = node.children[want]
		} else {
			node = node.children[bit]
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	root := &BitTrieNode{}
	for _, num := range a {
		insertBitTrie(root, num)
	}

	maxXOR := 0
	for _, num := range a {
		xorVal := findMaxXOR(root, num)
		if xorVal > maxXOR {
			maxXOR = xorVal
		}
	}

	fmt.Fprintln(writer, maxXOR)
}
