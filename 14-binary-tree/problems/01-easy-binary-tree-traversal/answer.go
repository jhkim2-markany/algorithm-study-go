package main

import (
	"bufio"
	"fmt"
	"os"
)

// binaryTreeTraversal은 이진 트리의 전위, 중위, 후위 순회 결과를 반환한다.
//
// [매개변수]
//   - left: 각 노드의 왼쪽 자식 배열 (-1이면 자식 없음)
//   - right: 각 노드의 오른쪽 자식 배열 (-1이면 자식 없음)
//   - root: 루트 노드 번호
//
// [반환값]
//   - []int: 전위 순회 결과
//   - []int: 중위 순회 결과
//   - []int: 후위 순회 결과
//
// [알고리즘 힌트]
//
//	전위 순회: 현재 → 왼쪽 → 오른쪽 순서로 방문한다.
//	중위 순회: 왼쪽 → 현재 → 오른쪽 순서로 방문한다.
//	후위 순회: 왼쪽 → 오른쪽 → 현재 순서로 방문한다.
//	각 순회를 재귀 함수로 구현하며, 노드가 -1이면 반환한다.
func binaryTreeTraversal(left, right []int, root int) ([]int, []int, []int) {
	var pre, in, post []int

	var preorder func(node int)
	preorder = func(node int) {
		if node == -1 {
			return
		}
		pre = append(pre, node)
		preorder(left[node])
		preorder(right[node])
	}

	var inorder func(node int)
	inorder = func(node int) {
		if node == -1 {
			return
		}
		inorder(left[node])
		in = append(in, node)
		inorder(right[node])
	}

	var postorder func(node int)
	postorder = func(node int) {
		if node == -1 {
			return
		}
		postorder(left[node])
		postorder(right[node])
		post = append(post, node)
	}

	preorder(root)
	inorder(root)
	postorder(root)

	return pre, in, post
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 왼쪽/오른쪽 자식 배열 초기화
	left := make([]int, n+1)
	right := make([]int, n+1)

	// 각 노드의 자식 정보 입력
	for i := 0; i < n; i++ {
		var node, l, r int
		fmt.Fscan(reader, &node, &l, &r)
		left[node] = l
		right[node] = r
	}

	// 핵심 함수 호출
	pre, in, post := binaryTreeTraversal(left, right, 1)

	// 전위 순회 출력
	for i, v := range pre {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	// 중위 순회 출력
	for i, v := range in {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	// 후위 순회 출력
	for i, v := range post {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
