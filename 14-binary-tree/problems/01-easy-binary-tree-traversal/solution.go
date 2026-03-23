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
func binaryTreeTraversal(left, right []int, root int) ([]int, []int, []int) {
	// 여기에 코드를 작성하세요
	return nil, nil, nil
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
