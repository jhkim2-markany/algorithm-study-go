package main

import (
	"bufio"
	"fmt"
	"os"
)

// 비트 트라이 노드 구조체
// 각 노드는 비트 0과 1에 대한 두 개의 자식을 가진다
type BitTrieNode struct {
	children [2]*BitTrieNode
}

// 비트 트라이에 숫자를 삽입한다
// 최상위 비트(MSB)부터 최하위 비트(LSB)까지 순서대로 삽입한다
func insert(root *BitTrieNode, num int) {
	node := root
	// 30번째 비트부터 0번째 비트까지 (10⁹ < 2³⁰)
	for i := 30; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &BitTrieNode{}
		}
		node = node.children[bit]
	}
}

// 주어진 숫자와 XOR 결과가 최대가 되는 값을 트라이에서 찾는다
// 각 비트에서 현재 비트와 반대되는 경로를 우선 선택한다
func findMaxXOR(root *BitTrieNode, num int) int {
	node := root
	result := 0
	for i := 30; i >= 0; i-- {
		bit := (num >> i) & 1
		// XOR을 최대화하려면 반대 비트 경로를 선택한다
		want := 1 - bit
		if node.children[want] != nil {
			result |= (1 << i) // 해당 비트 위치에서 XOR 결과가 1이 된다
			node = node.children[want]
		} else {
			// 반대 비트 경로가 없으면 같은 비트 경로로 간다
			node = node.children[bit]
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기 N
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 비트 트라이 생성 및 모든 수 삽입
	root := &BitTrieNode{}
	for _, num := range a {
		insert(root, num)
	}

	// 각 수에 대해 XOR 최댓값을 구하고 전체 최댓값을 갱신한다
	maxXOR := 0
	for _, num := range a {
		xorVal := findMaxXOR(root, num)
		if xorVal > maxXOR {
			maxXOR = xorVal
		}
	}

	fmt.Fprintln(writer, maxXOR)
}
