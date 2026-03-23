package main

import "fmt"

// 이진 트리 기본 구현 - 트리 구성, 순회, 높이 계산
// 시간 복잡도: O(N) (N: 노드 수)
// 공간 복잡도: O(N)

// 이진 트리 노드를 표현하는 구조체
type Node struct {
	value int
	left  *Node
	right *Node
}

// newNode 함수는 새로운 노드를 생성한다
func newNode(val int) *Node {
	return &Node{value: val}
}

// preorder 함수는 전위 순회를 수행한다 (현재 → 왼쪽 → 오른쪽)
func preorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	// 현재 노드를 먼저 방문
	*result = append(*result, node.value)
	preorder(node.left, result)
	preorder(node.right, result)
}

// inorder 함수는 중위 순회를 수행한다 (왼쪽 → 현재 → 오른쪽)
func inorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inorder(node.left, result)
	// 왼쪽 서브트리를 모두 방문한 뒤 현재 노드를 방문
	*result = append(*result, node.value)
	inorder(node.right, result)
}

// postorder 함수는 후위 순회를 수행한다 (왼쪽 → 오른쪽 → 현재)
func postorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	postorder(node.left, result)
	postorder(node.right, result)
	// 자식을 모두 방문한 뒤 현재 노드를 방문
	*result = append(*result, node.value)
}

// height 함수는 이진 트리의 높이를 반환한다
func height(node *Node) int {
	if node == nil {
		// 빈 트리의 높이는 -1
		return -1
	}
	// 왼쪽과 오른쪽 서브트리 높이 중 큰 값 + 1
	leftH := height(node.left)
	rightH := height(node.right)
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

// countNodes 함수는 이진 트리의 노드 수를 반환한다
func countNodes(node *Node) int {
	if node == nil {
		return 0
	}
	// 왼쪽 서브트리 노드 수 + 오른쪽 서브트리 노드 수 + 현재 노드
	return countNodes(node.left) + countNodes(node.right) + 1
}

func main() {
	// 이진 트리 구성
	// 구조:
	//        1
	//       / \
	//      2   3
	//     / \   \
	//    4   5   6
	//   /
	//  7
	root := newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.left.left = newNode(4)
	root.left.right = newNode(5)
	root.right.right = newNode(6)
	root.left.left.left = newNode(7)

	// 전위 순회
	pre := []int{}
	preorder(root, &pre)
	fmt.Printf("전위 순회: %v\n", pre)

	// 중위 순회
	in := []int{}
	inorder(root, &in)
	fmt.Printf("중위 순회: %v\n", in)

	// 후위 순회
	post := []int{}
	postorder(root, &post)
	fmt.Printf("후위 순회: %v\n", post)

	// 트리 높이
	fmt.Printf("트리 높이: %d\n", height(root))

	// 노드 수
	fmt.Printf("노드 수: %d\n", countNodes(root))
}
