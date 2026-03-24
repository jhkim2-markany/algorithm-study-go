package main

import (
	"bufio"
	"fmt"
	"os"
)

// twoStacks는 두 스택에서 꺼낼 수 있는 최대 원소 개수를 반환한다.
//
// [매개변수]
//   - maxSum: 최대 허용 합
//   - a: 스택 A의 원소 배열 (위에서 아래 순서)
//   - b: 스택 B의 원소 배열 (위에서 아래 순서)
//
// [반환값]
//   - int: 꺼낼 수 있는 최대 원소 개수
//
// [알고리즘 힌트]
//
//	먼저 A에서 가능한 만큼 꺼낸 뒤, A에서 하나씩 되돌리면서
//	B에서 추가로 꺼낼 수 있는 개수를 확인한다.
func twoStacks(maxSum int, a, b []int) int {
	// 스택 A에서 가능한 만큼 꺼냄
	sumA := 0
	i := 0
	for i < len(a) && sumA+a[i] <= maxSum {
		sumA += a[i]
		i++
	}

	// 현재 최대 개수
	best := i

	// A에서 하나씩 되돌리면서 B에서 추가로 꺼냄
	j := 0
	for i >= 0 {
		// B에서 가능한 만큼 꺼냄
		for j < len(b) && sumA+b[j] <= maxSum {
			sumA += b[j]
			j++
		}

		// 최대 개수 갱신
		count := i + j
		if count > best {
			best = count
		}

		// A에서 하나 되돌림
		if i > 0 {
			i--
			sumA -= a[i]
		} else {
			break
		}
	}

	return best
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 게임 횟수 입력
	var g int
	fmt.Fscan(reader, &g)

	for ; g > 0; g-- {
		var n, m, maxSum int
		fmt.Fscan(reader, &n, &m, &maxSum)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		b := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &b[i])
		}

		// 핵심 함수 호출
		fmt.Fprintln(writer, twoStacks(maxSum, a, b))
	}
}
