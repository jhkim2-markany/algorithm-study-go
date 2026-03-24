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

// insertBST는 BST에 새 값을 삽입하고 루트를 반환한다.
//
// [매개변수]
//   - root: 이진 탐색 트리의 루트 노드 포인터
//   - data: 삽입할 값
//
// [반환값]
//   - *TreeNode: 삽입 후 트리의 루트 노드 포인터
//
// [알고리즘 힌트]
//
//	nil 노드에 도달하면 새 노드를 생성하여 반환한다.
//	삽입할 값이 현재 노드보다 작으면 왼쪽, 크면 오른쪽으로 재귀 호출한다.
//	재귀 반환값을 자식 포인터에 연결하여 트리 구조를 유지한다.
func insertBST(root *TreeNode, data int) *TreeNode {
	// 기저 조건: nil이면 새 노드 생성
	if root == nil {
		return &TreeNode{Data: data}
	}
	// 삽입할 값이 현재 노드보다 작으면 왼쪽으로
	if data < root.Data {
		root.Left = insertBST(root.Left, data)
	} else {
		// 삽입할 값이 현재 노드보다 크면 오른쪽으로
		root.Right = insertBST(root.Right, data)
	}
	return root
}

// printLevelOrder는 트리를 레벨 순서로 출력한다.
func printLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.Data)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
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
		root = insertBST(root, data)
	}

	// 새로 삽입할 값
	var val int
	fmt.Fscan(reader, &val)
	root = insertBST(root, val)

	printLevelOrder(root)
	fmt.Fprintln(writer)
}
