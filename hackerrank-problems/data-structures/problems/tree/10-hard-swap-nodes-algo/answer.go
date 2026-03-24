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
//
// [알고리즘 힌트]
//
//	DFS로 트리를 순회하면서 현재 깊이를 추적한다.
//	깊이가 k의 배수이면 왼쪽/오른쪽 자식을 교환한다.
//	교환 후에도 양쪽 서브트리를 계속 순회한다.
func swapNodes(root *TreeNode, k int) {
	// 헬퍼 함수: 현재 깊이를 추적하며 DFS
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		// 기저 조건: nil이면 반환
		if node == nil {
			return
		}
		// 깊이가 k의 배수이면 자식 교환
		if depth%k == 0 {
			node.Left, node.Right = node.Right, node.Left
		}
		// 왼쪽/오른쪽 서브트리를 계속 순회
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	// 루트의 깊이는 1
	dfs(root, 1)
}

// inorder는 이진 트리를 중위 순회하여 결과를 출력한다.
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	// 왼쪽 서브트리 순회
	inorder(root.Left)
	// 현재 노드 출력
	fmt.Printf("%d ", root.Data)
	// 오른쪽 서브트리 순회
	inorder(root.Right)
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
		// 교환 수행
		swapNodes(root, k)
		// 중위 순회 결과 출력
		inorder(root)
		fmt.Fprintln(writer)
	}
}
