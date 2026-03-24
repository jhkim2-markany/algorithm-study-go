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
//
// [알고리즘 힌트]
//
//	BST의 성질을 이용한다: 왼쪽 < 루트 < 오른쪽.
//	두 값이 모두 현재 노드보다 작으면 왼쪽으로, 모두 크면 오른쪽으로 이동한다.
//	두 값이 현재 노드의 양쪽에 걸치면 현재 노드가 LCA이다.
func lca(root *TreeNode, v1, v2 int) *TreeNode {
	// 현재 노드가 nil이면 반환
	if root == nil {
		return nil
	}
	// 두 값이 모두 현재 노드보다 작으면 왼쪽으로 이동
	if v1 < root.Data && v2 < root.Data {
		return lca(root.Left, v1, v2)
	}
	// 두 값이 모두 현재 노드보다 크면 오른쪽으로 이동
	if v1 > root.Data && v2 > root.Data {
		return lca(root.Right, v1, v2)
	}
	// 두 값이 양쪽에 걸치거나 하나가 현재 노드와 같으면 현재 노드가 LCA
	return root
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
