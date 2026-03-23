package main

import (
	"bufio"
	"fmt"
	"os"
)

// reconstructPostorder는 전위 순회와 중위 순회 결과로 후위 순회 결과를 복원한다.
//
// [매개변수]
//   - preorder: 전위 순회 결과 배열
//   - inorder: 중위 순회 결과 배열
//
// [반환값]
//   - []int: 후위 순회 결과 배열
func reconstructPostorder(preorder, inorder []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 전위 순회 결과 입력
	preorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &preorder[i])
	}

	// 중위 순회 결과 입력
	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &inorder[i])
	}

	// 핵심 함수 호출
	result := reconstructPostorder(preorder, inorder)

	// 결과 출력
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
