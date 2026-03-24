package main

import "fmt"

// 이진 탐색 트리 (Binary Search Tree) - 기본 연산 예시
// BST의 삽입, 탐색, 삭제, 높이 계산과 순회(전위, 중위, 후위, 레벨 순서) 연산을 구현한다.
//
// 시간 복잡도:
//   - 삽입 (Insert):           O(H) (H: 트리 높이, 균형 시 O(log N), 최악 O(N))
//   - 탐색 (SearchBST):       O(H)
//   - 삭제 (DeleteBST):       O(H)
//   - 높이 (Height):          O(N) (N: 노드 수, 모든 노드 방문)
//   - 전위 순회 (Preorder):    O(N)
//   - 중위 순회 (Inorder):     O(N)
//   - 후위 순회 (Postorder):   O(N)
//   - 레벨 순회 (LevelOrder):  O(N)
//
// 공간 복잡도:
//   - 트리 저장:               O(N)
//   - 재귀 순회 호출 스택:     O(H)
//   - 레벨 순회 큐:            O(W) (W: 트리의 최대 너비)

// TreeNode는 이진 트리의 노드를 나타낸다.
// Data 필드에 값을 저장하고, Left와 Right로 자식 노드를 가리킨다.
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// Insert 함수는 BST에 새 값을 삽입한다.
// 현재 노드보다 작으면 왼쪽, 크면 오른쪽 서브트리로 재귀 이동한다.
// 이미 같은 값이 있으면 삽입하지 않는다.
func Insert(root *TreeNode, data int) *TreeNode {
	// 빈 위치를 찾으면 새 노드를 생성하여 반환
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		// 삽입할 값이 현재 노드보다 작으면 왼쪽 서브트리로 이동
		root.Left = Insert(root.Left, data)
	} else if data > root.Data {
		// 삽입할 값이 현재 노드보다 크면 오른쪽 서브트리로 이동
		root.Right = Insert(root.Right, data)
	}
	// 같은 값이면 무시 (중복 허용하지 않음)

	return root
}

// Preorder 함수는 전위 순회를 수행한다.
// 방문 순서: 현재 노드 → 왼쪽 서브트리 → 오른쪽 서브트리
// 트리의 구조를 복사하거나 직렬화할 때 유용하다.
func Preorder(root *TreeNode) {
	if root == nil {
		return
	}
	// 현재 노드를 먼저 방문
	fmt.Printf("%d ", root.Data)
	// 왼쪽 서브트리 순회
	Preorder(root.Left)
	// 오른쪽 서브트리 순회
	Preorder(root.Right)
}

// Inorder 함수는 중위 순회를 수행한다.
// 방문 순서: 왼쪽 서브트리 → 현재 노드 → 오른쪽 서브트리
// BST에서 중위 순회를 하면 값이 오름차순으로 출력된다.
func Inorder(root *TreeNode) {
	if root == nil {
		return
	}
	// 왼쪽 서브트리 순회
	Inorder(root.Left)
	// 현재 노드 방문
	fmt.Printf("%d ", root.Data)
	// 오른쪽 서브트리 순회
	Inorder(root.Right)
}

// Postorder 함수는 후위 순회를 수행한다.
// 방문 순서: 왼쪽 서브트리 → 오른쪽 서브트리 → 현재 노드
// 트리를 삭제하거나 후위 표기식을 평가할 때 유용하다.
func Postorder(root *TreeNode) {
	if root == nil {
		return
	}
	// 왼쪽 서브트리 순회
	Postorder(root.Left)
	// 오른쪽 서브트리 순회
	Postorder(root.Right)
	// 현재 노드를 마지막에 방문
	fmt.Printf("%d ", root.Data)
}

// LevelOrder 함수는 레벨 순서(너비 우선) 순회를 수행한다.
// 큐를 사용하여 같은 깊이의 노드를 왼쪽에서 오른쪽으로 방문한다.
// 트리의 레벨별 구조를 파악할 때 유용하다.
func LevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	// 큐를 슬라이스로 구현 (FIFO)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// 큐의 맨 앞 노드를 꺼냄
		current := queue[0]
		queue = queue[1:]

		// 현재 노드 방문
		fmt.Printf("%d ", current.Data)

		// 왼쪽 자식이 있으면 큐에 추가
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		// 오른쪽 자식이 있으면 큐에 추가
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

// SearchBST 함수는 BST에서 target 값을 가진 노드를 탐색한다.
// BST 속성을 이용하여 현재 노드보다 작으면 왼쪽, 크면 오른쪽으로 이동한다.
// 값을 찾으면 해당 노드를, 찾지 못하면 nil을 반환한다.
func SearchBST(root *TreeNode, target int) *TreeNode {
	// 노드가 nil이거나 값을 찾은 경우
	if root == nil || root.Data == target {
		return root
	}

	if target < root.Data {
		// 찾는 값이 현재 노드보다 작으면 왼쪽 서브트리 탐색
		return SearchBST(root.Left, target)
	}
	// 찾는 값이 현재 노드보다 크면 오른쪽 서브트리 탐색
	return SearchBST(root.Right, target)
}

// DeleteBST 함수는 BST에서 key 값을 가진 노드를 삭제한다.
// 삭제에는 3가지 케이스가 있다:
//  1. 리프 노드 (자식 없음): 단순히 노드를 제거
//  2. 자식이 하나인 노드: 자식으로 대체
//  3. 자식이 둘인 노드: 중위 후속자(오른쪽 서브트리의 최솟값)로 대체 후 후속자 삭제
func DeleteBST(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Data {
		// 삭제할 값이 현재 노드보다 작으면 왼쪽 서브트리에서 삭제
		root.Left = DeleteBST(root.Left, key)
	} else if key > root.Data {
		// 삭제할 값이 현재 노드보다 크면 오른쪽 서브트리에서 삭제
		root.Right = DeleteBST(root.Right, key)
	} else {
		// 삭제할 노드를 찾은 경우

		// 케이스 1 & 2: 자식이 없거나 하나인 경우
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 케이스 3: 자식이 둘인 경우
		// 중위 후속자(오른쪽 서브트리의 최솟값)를 찾는다
		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}

		// 현재 노드의 값을 중위 후속자의 값으로 대체
		root.Data = minNode.Data
		// 중위 후속자를 오른쪽 서브트리에서 삭제
		root.Right = DeleteBST(root.Right, minNode.Data)
	}

	return root
}

// Height 함수는 트리의 높이를 계산한다.
// 트리의 높이는 루트에서 가장 깊은 리프까지의 간선(edge) 수이다.
// 빈 트리의 높이는 -1, 루트만 있는 트리의 높이는 0이다.
func Height(root *TreeNode) int {
	if root == nil {
		return -1
	}

	// 왼쪽과 오른쪽 서브트리의 높이를 재귀적으로 계산
	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)

	// 더 큰 높이에 1을 더하여 반환
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func main() {
	// --- BST 삽입 예시 ---
	fmt.Println("=== 이진 탐색 트리 기본 연산 ===")
	fmt.Println()

	// BST 구성: 50, 30, 70, 20, 40, 60, 80 순서로 삽입
	//
	//         50
	//        /  \
	//      30    70
	//     / \   / \
	//   20  40 60  80
	//
	var root *TreeNode
	values := []int{50, 30, 70, 20, 40, 60, 80}

	fmt.Println("--- 삽입 (Insert) ---")
	fmt.Printf("삽입 순서: %v\n", values)
	for _, v := range values {
		root = Insert(root, v)
	}
	fmt.Println("트리 구성 완료")

	// --- 순회 예시 ---
	fmt.Println("\n--- 전위 순회 (Preorder) ---")
	fmt.Print("결과: ")
	Preorder(root)
	fmt.Println() // 50 30 20 40 70 60 80

	fmt.Println("\n--- 중위 순회 (Inorder) ---")
	fmt.Print("결과: ")
	Inorder(root)
	fmt.Println() // 20 30 40 50 60 70 80 (오름차순)

	fmt.Println("\n--- 후위 순회 (Postorder) ---")
	fmt.Print("결과: ")
	Postorder(root)
	fmt.Println() // 20 40 30 60 80 70 50

	fmt.Println("\n--- 레벨 순회 (Level Order) ---")
	fmt.Print("결과: ")
	LevelOrder(root)
	fmt.Println() // 50 30 70 20 40 60 80

	// --- 추가 삽입 후 순회 ---
	fmt.Println("\n--- 추가 삽입 테스트 ---")
	root = Insert(root, 25)
	root = Insert(root, 35)
	fmt.Println("25, 35 삽입 후 중위 순회:")
	fmt.Print("결과: ")
	Inorder(root)
	fmt.Println() // 20 25 30 35 40 50 60 70 80

	// --- BST 탐색 예시 ---
	fmt.Println("\n--- BST 탐색 (SearchBST) ---")
	if node := SearchBST(root, 40); node != nil {
		fmt.Printf("값 40 탐색: 찾음 (노드 값: %d)\n", node.Data)
	}
	if node := SearchBST(root, 99); node == nil {
		fmt.Println("값 99 탐색: 찾지 못함 (nil)")
	}

	// --- 트리 높이 예시 ---
	fmt.Println("\n--- 트리 높이 (Height) ---")
	fmt.Printf("현재 트리 높이: %d\n", Height(root))

	// --- BST 삭제 예시 ---
	fmt.Println("\n--- BST 삭제 (DeleteBST) ---")

	// 케이스 1: 리프 노드 삭제
	fmt.Println("리프 노드 25 삭제:")
	root = DeleteBST(root, 25)
	fmt.Print("중위 순회: ")
	Inorder(root)
	fmt.Println()

	// 케이스 2: 자식이 하나인 노드 삭제
	fmt.Println("\n자식이 하나인 노드 35 삭제:")
	root = DeleteBST(root, 35)
	fmt.Print("중위 순회: ")
	Inorder(root)
	fmt.Println()

	// 케이스 3: 자식이 둘인 노드 삭제 (루트 노드)
	fmt.Println("\n자식이 둘인 노드 50 삭제 (루트):")
	root = DeleteBST(root, 50)
	fmt.Print("중위 순회: ")
	Inorder(root)
	fmt.Println()

	// 삭제 후 트리 높이 확인
	fmt.Printf("\n삭제 후 트리 높이: %d\n", Height(root))

	// 없는 값 삭제 시도
	fmt.Println("\n없는 값 100 삭제 시도:")
	root = DeleteBST(root, 100)
	fmt.Print("중위 순회: ")
	Inorder(root)
	fmt.Println()
}
