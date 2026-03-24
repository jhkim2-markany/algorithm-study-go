package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// minimumLoss는 매입가 - 매도가의 최솟값(양수)을 반환한다.
// 매입은 매도보다 먼저 일어나야 하며, 매입가 > 매도가이다.
//
// [매개변수]
//   - price: 연도별 집값 배열
//
// [반환값]
//   - int64: 최소 손실
//
// [알고리즘 힌트]
//
//	가격을 정렬하고 인접한 값의 차이를 확인한다.
//	원래 인덱스에서 매입(앞)이 매도(뒤)보다 먼저인지 검증한다.
func minimumLoss(price []int64) int64 {
	n := len(price)

	// 인덱스와 함께 정렬하기 위한 구조체
	type indexed struct {
		val int64
		idx int
	}

	arr := make([]indexed, n)
	for i, v := range price {
		arr[i] = indexed{v, i}
	}

	// 가격 오름차순 정렬
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val
	})

	// 인접한 값들의 차이 중 최소 손실 탐색
	minLoss := int64(math.MaxInt64)
	for i := 0; i < n-1; i++ {
		diff := arr[i+1].val - arr[i].val
		// 차이가 양수이고, 큰 값(매입)이 원래 배열에서 앞에 있어야 함
		if diff > 0 && diff < minLoss && arr[i+1].idx < arr[i].idx {
			minLoss = diff
		}
	}

	return minLoss
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	price := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &price[i])
	}

	fmt.Fprintln(writer, minimumLoss(price))
}
