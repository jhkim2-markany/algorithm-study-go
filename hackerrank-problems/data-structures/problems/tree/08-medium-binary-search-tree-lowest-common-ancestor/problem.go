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

// lca는 BST에서 두 값의 최소 공통 조상을 찾아 반환한다.
//
// [매개변수]
//   - root: 이진 탐색 트리의 루트 노드 포인터
//   - v1: 첫 번째 값
//   - v2: 두 번째 값
//
// [반환값]
//   - *TreeNode: 최소 공통 조상 노드 포인터
func lca(root *TreeNode, v1, v2 int) *TreeNode {
	// 여기에 코드를 작성하세요
	return nil
}

// insert는 BST에 새 노드를 삽입한다.
func insert(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = insert(root.Left, data)
	} else {
		root.Right = insert(root.Right, data)
	}
	return root
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var root *TreeNode
	for i := 0; i < n; i++ {
		var data int
		fmt.Fscan(reader, &data)
		root = insert(root, data)
	}

	var v1, v2 int
	fmt.Fscan(reader, &v1, &v2)

	result := lca(root, v1, v2)
	if result != nil {
		fmt.Fprintln(writer, result.Data)
	}
}
