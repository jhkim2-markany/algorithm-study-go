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

	// 각 값의 빈도를 다항식 계수로 변환
	// freqA[v] = A에서 값 v의 등장 횟수
	maxVal := 0
	for i := 0; i < n; i++ {
		if a[i] > maxVal {
			maxVal = a[i]
		}
		if b[i] > maxVal {
			maxVal = b[i]
		}
	}

	freqA := make([]int, maxVal+1)
	freqB := make([]int, maxVal+1)
	for i := 0; i < n; i++ {
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

	// A의 최솟값 + B의 최솟값 ~ A의 최댓값 + B의 최댓값 범위에서 출력
	// 실제 가능한 합의 범위 계산
	minA, maxA := a[0], a[0]
	minB, maxB := b[0], b[0]
	for i := 1; i < n; i++ {
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

	// 결과 출력: 합성곱 결과에서 유효 범위만 출력
	for s := minSum; s <= maxSum; s++ {
		if s > minSum {
			fmt.Fprint(writer, " ")
		}
		cnt := int(math.Round(real(fa[s])))
		fmt.Fprint(writer, cnt)
	}
	fmt.Fprintln(writer)
}
