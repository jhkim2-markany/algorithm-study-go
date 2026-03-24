package main

import (
	"bufio"
	"fmt"
	"os"
)

const modVal = 1000000007

// constructArray는 조건을 만족하는 배열의 개수를 반환한다.
//
// [매개변수]
//   - n: 배열의 길이
//   - k: 원소의 최댓값
//   - x: 마지막 원소의 값
//
// [반환값]
//   - int: 조건을 만족하는 배열의 수 (mod 10^9+7)
//
// [알고리즘 힌트]
//
//	현재 값이 1인 경우와 아닌 경우를 분리하여 DP를 수행한다.
//	대칭성을 이용하여 "1이 아닌 특정 값"의 개수를 구한다.
func constructArray(n, k, x int) int {
	// one: 현재 위치의 값이 1인 배열의 수
	// notOne: 현재 위치의 값이 1이 아닌 배열의 수
	one := 1    // 첫 번째 위치는 1
	notOne := 0 // 첫 번째 위치에서 1이 아닌 경우는 없음

	for i := 2; i <= n; i++ {
		// 새로운 상태 계산
		newOne := notOne % modVal
		newNotOne := (one%modVal*int64ToInt(int64(k-1)) + notOne%modVal*int64ToInt(int64(k-2))) % modVal

		one = newOne
		notOne = newNotOne
	}

	// 마지막 원소가 x인 경우
	if x == 1 {
		return one
	}
	// x != 1: 대칭성에 의해 notOne을 (K-1)로 나눔
	// notOne / (K-1)을 모듈러 역원으로 계산
	return modDiv(notOne, k-1)
}

// int64ToInt는 int64를 int로 변환한다.
func int64ToInt(v int64) int {
	return int(v % int64(modVal))
}

// modPow는 밑^지수 mod m을 계산한다.
func modPow(base, exp, m int) int {
	result := 1
	base %= m
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % m
		}
		exp /= 2
		base = base * base % m
	}
	return result
}

// modDiv는 a / b mod modVal을 계산한다 (페르마 소정리 이용).
func modDiv(a, b int) int {
	return a % modVal * modPow(b, modVal-2, modVal) % modVal
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: n, k, x
	var n, k, x int
	fmt.Fscan(reader, &n, &k, &x)

	// 핵심 함수 호출 및 결과 출력
	result := constructArray(n, k, x)
	fmt.Fprintln(writer, result)
}
