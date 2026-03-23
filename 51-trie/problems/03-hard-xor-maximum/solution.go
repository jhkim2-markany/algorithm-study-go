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
func insertBitTrie(root *BitTrieNode, num int) {
	// 여기에 코드를 작성하세요
}

// findMaxXOR는 주어진 숫자와 XOR 결과가 최대가 되는 값을 트라이에서 찾는다.
//
// [매개변수]
//   - root: 비트 트라이의 루트 노드
//   - num: XOR 대상 정수
//
// [반환값]
//   - int: num과 XOR했을 때 최대가 되는 XOR 값
func findMaxXOR(root *BitTrieNode, num int) int {
	// 여기에 코드를 작성하세요
	return 0
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
