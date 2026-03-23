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

	// 역변환 시 N으로 나누기
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

	// 다항식 A 입력
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// 다항식 B 입력
	var m int
	fmt.Fscan(reader, &m)
	b := make([]int, m+1)
	for i := 0; i <= m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	// 결과 다항식 크기 계산
	resultLen := len(a) + len(b) - 1

	// N을 2의 거듭제곱으로 올림
	sz := 1
	for sz < resultLen {
		sz <<= 1
	}

	// 복소수 배열로 변환
	fa := make([]complex128, sz)
	fb := make([]complex128, sz)
	for i := 0; i < len(a); i++ {
		fa[i] = complex(float64(a[i]), 0)
	}
	for i := 0; i < len(b); i++ {
		fb[i] = complex(float64(b[i]), 0)
	}

	// FFT → 점별 곱셈 → IFFT
	fft(fa, false)
	fft(fb, false)
	for i := 0; i < sz; i++ {
		fa[i] *= fb[i]
	}
	fft(fa, true)

	// 결과 출력
	for i := 0; i < resultLen; i++ {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, int(math.Round(real(fa[i]))))
	}
	fmt.Fprintln(writer)
}
