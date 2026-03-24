package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// TreeNode는 이진 트리의 노드를 나타낸다.
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// isBSTHelper는 범위를 이용하여 BST 유효성을 재귀적으로 검증한다.
func isBSTHelper(node *TreeNode, min, max int) bool {
	// 기저 조건: nil 노드는 유효
	if node == nil {
		return true
	}
	// 현재 노드의 값이 허용 범위를 벗어나면 유효하지 않음
	if node.Data <= min || node.Data >= max {
		return false
	}
	// 왼쪽 서브트리는 (min, 현재 값) 범위, 오른쪽은 (현재 값, max) 범위
	return isBSTHelper(node.Left, min, node.Data) &&
		isBSTHelper(node.Right, node.Data, max)
}

// isBST는 주어진 이진 트리가 유효한 BST인지 판별한다.
//
// [매개변수]
//   - root: 이진 트리의 루트 노드 포인터
//
// [반환값]
//   - bool: 유효한 BST이면 true, 아니면 false
//
// [알고리즘 힌트]
//
//	각 노드에 허용되는 값의 범위(min, max)를 전달하여 검증한다.
//	왼쪽 자식은 (min, 현재 값), 오른쪽 자식은 (현재 값, max) 범위를 가진다.
//	모든 노드가 범위 내에 있으면 유효한 BST이다.
func isBST(root *TreeNode) bool {
	return isBSTHelper(root, math.MinInt64, math.MaxInt64)
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

	if isBST(root) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
