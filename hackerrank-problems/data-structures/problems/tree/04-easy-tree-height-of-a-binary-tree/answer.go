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

// height는 이진 트리의 높이를 반환한다.
//
// [매개변수]
//   - root: 이진 트리의 루트 노드 포인터
//
// [반환값]
//   - int: 트리의 높이 (루트에서 가장 먼 리프까지의 간선 수)
//
// [알고리즘 힌트]
//
//	재귀적으로 왼쪽/오른쪽 서브트리의 높이를 구한다.
//	두 높이 중 큰 값에 1을 더하면 현재 서브트리의 높이가 된다.
//	nil 노드는 -1을 반환하여 리프 노드의 높이가 0이 되도록 한다.
func height(root *TreeNode) int {
	// 기저 조건: nil 노드의 높이는 -1
	if root == nil {
		return -1
	}
	// 왼쪽 서브트리의 높이를 재귀적으로 구함
	leftHeight := height(root.Left)
	// 오른쪽 서브트리의 높이를 재귀적으로 구함
	rightHeight := height(root.Right)
	// 두 높이 중 큰 값에 1을 더하여 반환
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
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

	fmt.Fprintln(writer, height(root))
}
