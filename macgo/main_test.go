package macgo_test

import (
	"testing"

	"github.com/hehecoi222/go-project/macgo"
)

func TestPrintHello(t *testing.T) {
	macgo.PrintHello()
}

func BenchmarkPrintHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		macgo.PrintHello()
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCases := []struct {
			desc   string
			N      int
			Output string
		}{
			{
				desc:   "123",
				N:      123,
				Output: "No",
			}, {
				desc:   "555",
				N:      555,
				Output: "Yes",
			},
			{
				desc:   "404",
				N:      404,
				Output: "Yes",
			},
		}
		b.ResetTimer()
		for _, tC := range testCases {
			b.Run(tC.desc, func(b *testing.B) {
				re := macgo.IsPalindrome(tC.N)
				if re != tC.Output {
					b.Error("expected ", tC.Output, ", got ", re)
				}
			})
		}
	}
}