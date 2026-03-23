package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
	"os"
)

// fft는 비재귀 쿨리-튜키 알고리즘으로 FFT를 수행한다
func fft(a []complex128, invert bool) {
	n := len(a)
	if n == 1 {
		return
	}

	// 비트 반전 순열
	for i, j := 1, 0; i < n; i++ {
		bit := n >> 1
		for ; j&bit != 0; bit >>= 1 {
			j ^= bit
		}
		j ^= bit
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	// 버터플라이 연산
	for length := 2; length <= n; length <<= 1 {
		angle := 2 * math.Pi / float64(length)
		if invert {
			angle = -angle
		}
		wn := cmplx.Exp(complex(0, angle))

		for i := 0; i < n; i += length {
			w := complex(1, 0)
			for j := 0; j < length/2; j++ {
				u := a[i+j]
				v := a[i+j+length/2] * w
				a[i+j] = u + v
				a[i+j+length/2] = u - v
				w *= wn
			}
		}
	}

	if invert {
		for i := range a {
			a[i] /= complex(float64(n), 0)
		}
	}
}

// convolutionCount는 두 배열의 값 빈도를 합성곱하여 각 합에 대한 쌍의 개수를 반환한다.
// 배열 A와 B에서 각각 하나씩 원소를 골라 합이 s가 되는 쌍의 수를 구한다.
//
// [매개변수]
//   - a: 첫 번째 정수 배열
//   - b: 두 번째 정수 배열
//
// [반환값]
//   - []int: 인덱스 s에 대해 a[i]+b[j]==s인 쌍의 개수 배열 (유효 범위만)
//
// [알고리즘 힌트]
//
//	각 배열의 값 빈도를 다항식 계수로 변환한다 (freqA[v] = A에서 값 v의 등장 횟수).
//	두 빈도 다항식을 FFT로 곱하면, 결과의 s번째 계수가 a[i]+b[j]==s인 쌍의 수가 된다.
//	유효 범위(minA+minB ~ maxA+maxB)만 추출하여 반환한다.
func convolutionCount(a, b []int) []int {
	// 각 값의 빈도를 다항식 계수로 변환
	maxVal := 0
	for i := 0; i < len(a); i++ {
		if a[i] > maxVal {
			maxVal = a[i]
		}
		if b[i] > maxVal {
			maxVal = b[i]
		}
	}

	freqA := make([]int, maxVal+1)
	freqB := make([]int, maxVal+1)
	for i := 0; i < len(a); i++ {
		freqA[a[i]]++
		freqB[b[i]]++
	}

	// 결과 다항식 크기 계산
	resultLen := 2*maxVal + 1

	sz := 1
	for sz < resultLen {
		sz <<= 1
	}

	// 복소수 배열로 변환
	fa := make([]complex128, sz)
	fb := make([]complex128, sz)
	for i := 0; i <= maxVal; i++ {
		fa[i] = complex(float64(freqA[i]), 0)
		fb[i] = complex(float64(freqB[i]), 0)
	}

	// FFT → 점별 곱셈 → IFFT
	fft(fa, false)
	fft(fb, false)
	for i := 0; i < sz; i++ {
		fa[i] *= fb[i]
	}
	fft(fa, true)

	// 유효 범위 계산
	minA, maxA := a[0], a[0]
	minB, maxB := b[0], b[0]
	for i := 1; i < len(a); i++ {
		if a[i] < minA {
			minA = a[i]
		}
		if a[i] > maxA {
			maxA = a[i]
		}
		if b[i] < minB {
			minB = b[i]
		}
		if b[i] > maxB {
			maxB = b[i]
		}
	}

	minSum := minA + minB
	maxSum := maxA + maxB

	// 유효 범위만 추출
	result := make([]int, maxSum-minSum+1)
	for s := minSum; s <= maxSum; s++ {
		result[s-minSum] = int(math.Round(real(fa[s])))
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 A, B 입력
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	// 핵심 함수 호출
	result := convolutionCount(a, b)

	// 결과 출력
	for i := 0; i < len(result); i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, result[i])
	}
	fmt.Fprintln(writer)
}
