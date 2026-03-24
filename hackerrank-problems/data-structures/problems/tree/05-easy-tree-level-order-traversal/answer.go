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

// levelOrder는 이진 트리를 레벨 순서로 순회하여 결과를 출력한다.
//
// [매개변수]
//   - root: 이진 트리의 루트 노드 포인터
//
// [반환값]
//   - 없음 (표준 출력으로 레벨 순서 순회 결과를 공백 구분하여 출력)
//
// [알고리즘 힌트]
//
//	큐를 사용하여 BFS를 수행한다.
//	루트를 큐에 넣고, 큐에서 꺼낸 노드의 자식을 순서대로 큐에 넣는다.
//	큐가 빌 때까지 반복하면 레벨 순서로 방문하게 된다.
func levelOrder(root *TreeNode) {
	// 루트가 nil이면 반환
	if root == nil {
		return
	}
	// 큐 초기화 및 루트 삽입
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// 큐에서 노드를 꺼냄
		node := queue[0]
		queue = queue[1:]
		// 현재 노드의 데이터를 출력
		fmt.Printf("%d ", node.Data)
		// 왼쪽 자식이 있으면 큐에 추가
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		// 오른쪽 자식이 있으면 큐에 추가
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
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

	levelOrder(root)
	fmt.Fprintln(writer)
}
