package sumofeven

import "testing"

var benchResult int
var _ = &benchResult

func benchmarkSumEven(b *testing.B, fn func([]int) int) {
	for _, size := range []int{16, 1024, 65536} {
		data := makeBenchmarkData(size)
		b.Run("size="+itoa(size), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				benchResult = fn(data)
			}
		})
	}
}

func BenchmarkSumOfEven(b *testing.B) {
	benchmarkSumEven(b, SumOfEven)
}

func BenchmarkSumOfEvenBit(b *testing.B) {
	benchmarkSumEven(b, SumOfEvenBit)
}

func BenchmarkSumOfEvenIfPrev(b *testing.B) {
	benchmarkSumEven(b, SumOfEvenIfPrev)
}

func makeBenchmarkData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = i - size/2
	}
	return data
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var digits [20]byte
	idx := len(digits)
	for n > 0 {
		idx--
		digits[idx] = byte('0' + n%10)
		n /= 10
	}
	return string(digits[idx:])
}
