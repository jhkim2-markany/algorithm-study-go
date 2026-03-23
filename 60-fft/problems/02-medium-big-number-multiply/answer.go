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

// bigNumberMultiply는 두 큰 수 문자열을 곱한 결과 문자열을 반환한다.
// FFT를 이용하여 각 자릿수를 다항식 계수로 취급하고 곱셈을 수행한다.
//
// [매개변수]
//   - sa: 첫 번째 큰 수 문자열
//   - sb: 두 번째 큰 수 문자열
//
// [반환값]
//   - string: 두 수를 곱한 결과 문자열
//
// [알고리즘 힌트]
//
//	각 자릿수를 역순으로 다항식 계수로 변환한 뒤 FFT 곱셈을 수행한다.
//	IFFT 결과를 정수로 반올림한 후, 낮은 자릿수부터 올림(carry) 처리를 하고
//	선행 0을 제거하여 역순으로 문자열을 구성한다.
func bigNumberMultiply(sa, sb string) string {
	na, nb := len(sa), len(sb)
	resultLen := na + nb - 1

	sz := 1
	for sz < na+nb {
		sz <<= 1
	}

	// 복소수 배열 생성
	fa := make([]complex128, sz)
	fb := make([]complex128, sz)

	// 문자열을 역순으로 복소수 배열에 저장
	for i := 0; i < na; i++ {
		fa[i] = complex(float64(sa[na-1-i]-'0'), 0)
	}
	for i := 0; i < nb; i++ {
		fb[i] = complex(float64(sb[nb-1-i]-'0'), 0)
	}

	// FFT → 점별 곱셈 → IFFT
	fft(fa, false)
	fft(fb, false)
	for i := 0; i < sz; i++ {
		fa[i] *= fb[i]
	}
	fft(fa, true)

	// 결과를 정수 배열로 변환하고 올림(carry) 처리
	result := make([]int, resultLen+1)
	for i := 0; i < resultLen; i++ {
		result[i] += int(math.Round(real(fa[i])))
	}

	// 자릿수 올림 처리
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= 10 {
			result[i+1] += result[i] / 10
			result[i] %= 10
		}
	}

	// 선행 0 제거 후 역순 출력
	top := len(result) - 1
	for top > 0 && result[top] == 0 {
		top--
	}

	buf := make([]byte, top+1)
	for i := top; i >= 0; i-- {
		buf[top-i] = byte(result[i]) + '0'
	}
	return string(buf)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 두 큰 수를 문자열로 입력
	var sa, sb string
	fmt.Fscan(reader, &sa)
	fmt.Fscan(reader, &sb)

	// 핵심 함수 호출
	fmt.Fprintln(writer, bigNumberMultiply(sa, sb))
}
