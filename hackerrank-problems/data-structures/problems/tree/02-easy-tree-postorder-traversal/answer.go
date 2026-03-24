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

// postorder는 이진 트리를 후위 순회하여 결과를 출력한다.
//
// [매개변수]
//   - root: 이진 트리의 루트 노드 포인터
//
// [반환값]
//   - 없음 (표준 출력으로 후위 순회 결과를 공백 구분하여 출력)
//
// [알고리즘 힌트]
//
//	후위 순회는 "왼쪽 → 오른쪽 → 루트" 순서이다.
//	왼쪽/오른쪽 서브트리를 먼저 재귀 호출한 후 현재 노드를 출력한다.
func postorder(root *TreeNode) {
	// 기저 조건: 노드가 nil이면 반환
	if root == nil {
		return
	}
	// 왼쪽 서브트리를 후위 순회
	postorder(root.Left)
	// 오른쪽 서브트리를 후위 순회
	postorder(root.Right)
	// 현재 노드의 데이터를 출력
	fmt.Printf("%d ", root.Data)
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

	postorder(root)
	fmt.Fprintln(writer)
}
