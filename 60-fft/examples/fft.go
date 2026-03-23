package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// FFT (고속 푸리에 변환) - 다항식 곱셈을 O(N log N)에 수행
// 시간 복잡도: O(N log N)
// 공간 복잡도: O(N)

// fft는 쿨리-튜키 알고리즘으로 DFT/IDFT를 수행한다.
// invert가 true이면 IDFT(역변환)를 수행한다.
func fft(a []complex128, invert bool) {
	n := len(a)
	if n == 1 {
		return
	}

	// 비트 반전 순열로 배열 재배치
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

	// 아래에서 위로 버터플라이 연산 수행
	for length := 2; length <= n; length <<= 1 {
		// 단위근 계산
		angle := 2 * math.Pi / float64(length)
		if invert {
			angle = -angle
		}
		wn := cmplx.Exp(complex(0, angle))

		for i := 0; i < n; i += length {
			w := complex(1, 0)
			for j := 0; j < length/2; j++ {
				// 버터플라이 연산: 짝수/홀수 항 결합
				u := a[i+j]
				v := a[i+j+length/2] * w
				a[i+j] = u + v
				a[i+j+length/2] = u - v
				w *= wn
			}
		}
	}

	// IDFT인 경우 N으로 나누기
	if invert {
		for i := range a {
			a[i] /= complex(float64(n), 0)
		}
	}
}

// multiply는 두 다항식의 계수 배열을 받아 곱한 결과를 반환한다.
func multiply(a, b []int) []int {
	// 결과 다항식의 최소 크기 계산
	resultLen := len(a) + len(b) - 1

	// N을 2의 거듭제곱으로 올림
	n := 1
	for n < resultLen {
		n <<= 1
	}

	// 복소수 배열로 변환 (0으로 패딩)
	fa := make([]complex128, n)
	fb := make([]complex128, n)
	for i := 0; i < len(a); i++ {
		fa[i] = complex(float64(a[i]), 0)
	}
	for i := 0; i < len(b); i++ {
		fb[i] = complex(float64(b[i]), 0)
	}

	// 정방향 FFT (계수 → 점값)
	fft(fa, false)
	fft(fb, false)

	// 점별 곱셈
	for i := 0; i < n; i++ {
		fa[i] *= fb[i]
	}

	// 역방향 FFT (점값 → 계수)
	fft(fa, true)

	// 실수부를 반올림하여 정수 계수 복원
	result := make([]int, resultLen)
	for i := 0; i < resultLen; i++ {
		result[i] = int(math.Round(real(fa[i])))
	}
	return result
}

func main() {
	// 예제: A(x) = 1 + 2x + 3x²,  B(x) = 4 + 5x
	a := []int{1, 2, 3}
	b := []int{4, 5}

	fmt.Println("다항식 A 계수:", a)
	fmt.Println("다항식 B 계수:", b)

	// 다항식 곱셈 수행
	result := multiply(a, b)
	fmt.Println("A × B 결과 계수:", result)
	// 기대 출력: [4 13 22 15]
	// C(x) = 4 + 13x + 22x² + 15x³

	// 예제 2: 큰 다항식 곱셈
	c := []int{1, 1, 1, 1} // 1 + x + x² + x³
	d := []int{1, 1, 1, 1} // 1 + x + x² + x³
	result2 := multiply(c, d)
	fmt.Println("\n(1+x+x²+x³)² 결과:", result2)
	// 기대 출력: [1 2 3 4 3 2 1]
}
