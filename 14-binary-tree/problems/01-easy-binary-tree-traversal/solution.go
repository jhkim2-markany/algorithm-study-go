package main

import (
	"bufio"
	"fmt"
	"os"
)

// 이진 트리의 왼쪽/오른쪽 자식을 저장하는 배열
var left []int
var right []int

// preorder 함수는 전위 순회를 수행한다 (현재 → 왼쪽 → 오른쪽)
func preorder(node int, result *[]int) {
	if node == -1 {
		return
	}
	// 현재 노드를 먼저 방문
	*result = append(*result, node)
	preorder(left[node], result)
	preorder(right[node], result)
}

// inorder 함수는 중위 순회를 수행한다 (왼쪽 → 현재 → 오른쪽)
func inorder(node int, result *[]int) {
	if node == -1 {
		return
	}
	inorder(left[node], result)
	// 왼쪽을 모두 방문한 뒤 현재 노드 방문
	*result = append(*result, node)
	inorder(right[node], result)
}

// postorder 함수는 후위 순회를 수행한다 (왼쪽 → 오른쪽 → 현재)
func postorder(node int, result *[]int) {
	if node == -1 {
		return
	}
	postorder(left[node], result)
	postorder(right[node], result)
	// 자식을 모두 방문한 뒤 현재 노드 방문
	*result = append(*result, node)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 왼쪽/오른쪽 자식 배열 초기화
	left = make([]int, n+1)
	right = make([]int, n+1)

	// 각 노드의 자식 정보 입력
	for i := 0; i < n; i++ {
		var node, l, r int
		fmt.Fscan(reader, &node, &l, &r)
		left[node] = l
		right[node] = r
	}

	// 전위 순회 출력
	pre := []int{}
	preorder(1, &pre)
	for i, v := range pre {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	// 중위 순회 출력
	in := []int{}
	inorder(1, &in)
	for i, v := range in {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	// 후위 순회 출력
	post := []int{}
	postorder(1, &post)
	for i, v := range post {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
