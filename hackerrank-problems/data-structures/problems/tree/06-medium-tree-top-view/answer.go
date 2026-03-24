package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// TreeNode는 이진 트리의 노드를 나타낸다.
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// queueItem은 BFS 큐에 저장할 노드와 수평 거리 쌍이다.
type queueItem struct {
	node *TreeNode
	hd   int
}

// topView는 이진 트리의 탑 뷰를 출력한다.
//
// [매개변수]
//   - root: 이진 트리의 루트 노드 포인터
//
// [반환값]
//   - 없음 (표준 출력으로 탑 뷰 결과를 공백 구분하여 출력)
//
// [알고리즘 힌트]
//
//	BFS로 레벨 순서 순회하면서 각 노드에 수평 거리를 부여한다.
//	각 수평 거리에서 처음 방문한 노드만 맵에 저장한다.
//	수평 거리를 정렬하여 왼쪽에서 오른쪽 순서로 출력한다.
func topView(root *TreeNode) {
	// 루트가 nil이면 반환
	if root == nil {
		return
	}
	// 수평 거리별 첫 번째 노드 데이터를 저장하는 맵
	hdMap := make(map[int]int)
	// BFS 큐 초기화
	queue := []queueItem{{node: root, hd: 0}}

	for len(queue) > 0 {
		// 큐에서 항목을 꺼냄
		item := queue[0]
		queue = queue[1:]
		// 해당 수평 거리에 처음 방문한 경우에만 저장
		if _, exists := hdMap[item.hd]; !exists {
			hdMap[item.hd] = item.node.Data
		}
		// 왼쪽 자식을 큐에 추가 (수평 거리 -1)
		if item.node.Left != nil {
			queue = append(queue, queueItem{node: item.node.Left, hd: item.hd - 1})
		}
		// 오른쪽 자식을 큐에 추가 (수평 거리 +1)
		if item.node.Right != nil {
			queue = append(queue, queueItem{node: item.node.Right, hd: item.hd + 1})
		}
	}
	// 수평 거리를 정렬
	keys := make([]int, 0, len(hdMap))
	for k := range hdMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// 왼쪽에서 오른쪽 순서로 출력
	for _, k := range keys {
		fmt.Printf("%d ", hdMap[k])
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

	topView(root)
	fmt.Fprintln(writer)
}
