package main

import (
	"bufio"
	"fmt"
	"os"
)

// getPrimes는 에라토스테네스의 체로 소수 목록을 생성한다.
func getPrimes(count int) []int {
	primes := []int{}
	sieve := make([]bool, 50000)
	for i := 2; len(primes) < count; i++ {
		if i >= len(sieve) {
			break
		}
		if !sieve[i] {
			primes = append(primes, i)
			for j := i * i; j < len(sieve); j += i {
				sieve[j] = true
			}
		}
	}
	return primes
}

// waiter는 접시 분류 결과를 반환한다.
//
// [매개변수]
//   - number: 접시 번호 배열 (스택의 아래에서 위 순서)
//   - q: 반복 횟수
//
// [반환값]
//   - []int: 결과 접시 번호 목록
//
// [알고리즘 힌트]
//
//	에라토스테네스의 체로 소수를 구한 뒤, 각 반복에서
//	스택의 최상위부터 소수 나눗셈 여부에 따라 분류한다.
func waiter(number []int, q int) []int {
	// 필요한 소수 생성
	primes := getPrimes(q)

	result := []int{}

	// 초기 스택 A: 입력 배열을 스택으로 사용 (마지막 원소가 최상위)
	a := make([]int, len(number))
	copy(a, number)

	for i := 0; i < q; i++ {
		// B_i 스택과 새로운 A 스택
		b := []int{}
		newA := []int{}

		// 현재 스택의 최상위부터 확인
		for j := len(a) - 1; j >= 0; j-- {
			if a[j]%primes[i] == 0 {
				// 소수로 나누어지면 B에 추가
				b = append(b, a[j])
			} else {
				// 아니면 새 A에 추가
				newA = append(newA, a[j])
			}
		}

		// B_i의 최상위부터 결과에 추가
		for j := len(b) - 1; j >= 0; j-- {
			result = append(result, b[j])
		}

		// 다음 반복을 위해 A 갱신
		a = newA
	}

	// 마지막 A의 최상위부터 결과에 추가
	for j := len(a) - 1; j >= 0; j-- {
		result = append(result, a[j])
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 접시 개수와 반복 횟수 입력
	var n, q int
	fmt.Fscan(reader, &n, &q)

	// 접시 번호 입력
	number := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &number[i])
	}

	// 핵심 함수 호출
	result := waiter(number, q)

	// 결과 출력
	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}
