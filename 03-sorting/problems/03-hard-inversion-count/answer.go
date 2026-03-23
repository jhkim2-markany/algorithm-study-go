package main

import (
	"bufio"
	"fmt"
	"os"
)

// countInversions는 배열에서 역전(Inversion)의 총 개수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열 (길이 N, 1 ≤ N ≤ 500,000)
//
// [반환값]
//   - int64: 역전의 총 개수 (i < j이면서 arr[i] > arr[j]인 쌍의 수)
//
// [알고리즘 힌트]
//
//	병합 정렬(Merge Sort)을 변형하여 역전 개수를 센다.
//	배열을 반으로 나누어 재귀적으로 정렬하면서,
//	병합(merge) 단계에서 왼쪽 배열의 원소가 오른쪽 배열의 원소보다
//	클 때 역전이 발생한다.
//
//	왼쪽 배열의 i번째 원소 > 오른쪽 배열의 j번째 원소이면,
//	왼쪽 배열의 i부터 mid까지 모든 원소가 arr[j]보다 크므로
//	(mid - i + 1)개의 역전이 추가된다.
//
//	시간복잡도: O(N log N), 공간복잡도: O(N)
//
//	예시: arr=[3, 1, 2, 5, 4]
//	  → 역전 쌍: (3,1), (3,2), (5,4) → 총 3개
func countInversions(arr []int) int64 {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	temp := make([]int, n)
	return mergeSortCount(arr, temp, 0, n-1)
}

// mergeSortCount는 병합 정렬을 수행하면서 역전 개수를 반환한다
func mergeSortCount(arr, temp []int, left, right int) int64 {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	var count int64
	count += mergeSortCount(arr, temp, left, mid)
	count += mergeSortCount(arr, temp, mid+1, right)
	count += mergeCount(arr, temp, left, mid, right)
	return count
}

// mergeCount는 두 정렬된 부분 배열을 병합하면서 역전 개수를 반환한다
func mergeCount(arr, temp []int, left, mid, right int) int64 {
	i := left
	j := mid + 1
	k := left
	var count int64

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			// 왼쪽 배열의 i번째 원소가 오른쪽보다 크면 역전 발생
			count += int64(mid - i + 1)
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	for idx := left; idx <= right; idx++ {
		arr[idx] = temp[idx]
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := countInversions(arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
