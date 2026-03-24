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

// preorder는 이진 트리를 전위 순회하여 결과를 출력한다.
//
// [매개변수]
//   - root: 이진 트리의 루트 노드 포인터
//
// [반환값]
//   - 없음 (표준 출력으로 전위 순회 결과를 공백 구분하여 출력)
func preorder(root *TreeNode) {
	// 여기에 코드를 작성하세요
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

	preorder(root)
	fmt.Fprintln(writer)
}
