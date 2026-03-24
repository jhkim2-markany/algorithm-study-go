package main

import (
	"bufio"
	"fmt"
	"os"
)

// TreeNode는 이진 트리의 노드를 나타낸다.
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// swapNodes는 깊이가 k의 배수인 모든 노드에서 자식을 교환한다.
//
// [매개변수]
//   - root: 이진 트리의 루트 노드 포인터
//   - k: 교환 대상 깊이의 배수
//
// [반환값]
//   - 없음 (트리 구조를 직접 변경)
func swapNodes(root *TreeNode, k int) {
	// 여기에 코드를 작성하세요
}

// inorder는 이진 트리를 중위 순회하여 결과를 출력한다.
func inorder(root *TreeNode) {
	// 여기에 코드를 작성하세요
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	// 인덱스 기반 트리 구성
	nodes := make([]*TreeNode, n+1)
	for i := 1; i <= n; i++ {
		nodes[i] = &TreeNode{Data: i}
	}

	for i := 1; i <= n; i++ {
		var left, right int
		fmt.Fscan(reader, &left, &right)
		if left != -1 {
			nodes[i].Left = nodes[left]
		}
		if right != -1 {
			nodes[i].Right = nodes[right]
		}
	}

	root := nodes[1]

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var k int
		fmt.Fscan(reader, &k)
		swapNodes(root, k)
		inorder(root)
		fmt.Fprintln(writer)
	}
}
