package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기 N, 윈도우 크기 K
	var n, k int
	fmt.Fscan(reader, &n, &k)

	// 배열 입력
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 1단계: 모노톤 덱으로 각 윈도우의 최솟값을 구한다
	// windowMin[j] = j번째 윈도우(인덱스 j ~ j+K-1)의 최솟값
	numWindows := n - k + 1
	windowMin := make([]int, numWindows)

	// 단조 증가 덱: 앞쪽이 항상 최솟값의 인덱스
	deque := make([]int, 0, n)

	for i := 0; i < n; i++ {
		// 윈도우 범위를 벗어난 인덱스를 앞에서 제거한다
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 현재 원소보다 크거나 같은 원소의 인덱스를 뒤에서 제거한다
		for len(deque) > 0 && a[deque[len(deque)-1]] >= a[i] {
			deque = deque[:len(deque)-1]
		}

		// 현재 인덱스를 덱에 추가한다
		deque = append(deque, i)

		// 윈도우가 완성된 시점부터 최솟값을 기록한다
		if i >= k-1 {
			windowMin[i-k+1] = a[deque[0]]
		}
	}

	// 2단계: 각 인덱스 i에 대해 B[i]를 구한다
	// i가 포함되는 윈도우 범위: j = max(0, i-K+1) ~ min(N-K, i)
	// B[i] = max(windowMin[j]) for j in [max(0,i-K+1), min(N-K, i)]
	// 이는 windowMin 배열에 대한 슬라이딩 윈도우 최댓값 문제이다

	// windowMin 배열에서 크기 K인 슬라이딩 윈도우 최댓값을 구한다
	// 단, i번째 원소가 포함되는 윈도우 범위의 크기는 min(i+1, K, N-i, N-K+1)이다
	// 실제로 windowMin의 인덱스 j에 대해, j번째 윈도우는 원소 j ~ j+K-1을 포함한다
	// 따라서 원소 i가 포함되는 윈도우: j in [max(0, i-K+1), min(numWindows-1, i)]

	// windowMin 배열에 대해 슬라이딩 윈도우 최댓값을 구한다
	// 윈도우 크기는 min(K, numWindows)이다
	// B[i]는 windowMin[max(0,i-K+1) .. min(numWindows-1,i)] 구간의 최댓값이다

	// 모노톤 덱으로 windowMin 배열의 슬라이딩 윈도우 최댓값을 구한다
	// 윈도우 크기 W = min(K, numWindows)
	w := k
	if w > numWindows {
		w = numWindows
	}

	// maxOfMin[j] = windowMin[j..j+W-1] 구간의 최댓값
	maxOfMin := make([]int, 0, numWindows)
	deque = deque[:0]

	for j := 0; j < numWindows; j++ {
		for len(deque) > 0 && deque[0] <= j-w {
			deque = deque[1:]
		}
		for len(deque) > 0 && windowMin[deque[len(deque)-1]] <= windowMin[j] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, j)
		if j >= w-1 {
			maxOfMin = append(maxOfMin, windowMin[deque[0]])
		}
	}

	// 3단계: B[i]를 계산하고 합을 구한다
	// 원소 i가 포함되는 윈도우 범위: [lo, hi] where lo=max(0,i-K+1), hi=min(numWindows-1,i)
	// 이 구간의 길이는 hi - lo + 1이며, 이 구간의 최댓값이 B[i]이다
	// maxOfMin 배열의 인덱스 t는 windowMin[t..t+W-1] 구간의 최댓값이다
	// lo = max(0, i-K+1), 이 구간의 시작 윈도우 인덱스
	// 구간 [lo, hi]의 최댓값을 구하기 위해 다시 모노톤 덱을 사용한다

	// 직접 계산: 각 i에 대해 B[i] 구하기
	// 구간 [lo, hi]에서 windowMin의 최댓값을 구해야 한다
	// 이를 위해 windowMin에 대한 sparse table을 구축한다

	// Sparse Table 구축 (구간 최댓값 쿼리)
	logN := 0
	for (1 << logN) <= numWindows {
		logN++
	}

	sparse := make([][]int, logN)
	sparse[0] = make([]int, numWindows)
	copy(sparse[0], windowMin)

	for p := 1; p < logN; p++ {
		length := 1 << p
		sparse[p] = make([]int, numWindows)
		for j := 0; j+length-1 < numWindows; j++ {
			if sparse[p-1][j] > sparse[p-1][j+length/2] {
				sparse[p][j] = sparse[p-1][j]
			} else {
				sparse[p][j] = sparse[p-1][j+length/2]
			}
		}
	}

	// 구간 최댓값 쿼리 함수
	queryMax := func(l, r int) int {
		if l > r {
			return 0
		}
		length := r - l + 1
		p := 0
		for (1 << (p + 1)) <= length {
			p++
		}
		a := sparse[p][l]
		b := sparse[p][r-(1<<p)+1]
		if a > b {
			return a
		}
		return b
	}

	// B[i]의 합을 계산한다
	sum := int64(0)
	for i := 0; i < n; i++ {
		lo := i - k + 1
		if lo < 0 {
			lo = 0
		}
		hi := i
		if hi > numWindows-1 {
			hi = numWindows - 1
		}
		// windowMin[lo..hi] 구간의 최댓값이 B[i]이다
		sum += int64(queryMax(lo, hi))
	}

	fmt.Fprintln(writer, sum)
}
