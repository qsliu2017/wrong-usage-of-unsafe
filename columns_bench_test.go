package columns_test

import (
	"testing"

	columns "github.com/qsliu2017/wrong-usage-of-unsafe"
)

func BenchmarkBadAt1(b *testing.B) {
	n := 1000
	cols := columns.New(n)
	b.StartTimer()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		_ = cols.BadAt1(i % n)
	}
}
func BenchmarkBadAt2(b *testing.B) {
	n := 1000
	cols := columns.New(n)
	b.StartTimer()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		_ = cols.BadAt2(i % n)
	}
}
func BenchmarkGoodAt1(b *testing.B) {
	n := 1000
	cols := columns.New(n)
	b.StartTimer()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		_ = cols.GoodAt1(i % n)
	}
}
func BenchmarkGoodAt2(b *testing.B) {
	n := 1000
	cols := columns.New(n)
	b.StartTimer()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		_ = cols.GoodAt2(i % n)
	}
}

func BenchmarkGoodAt3(b *testing.B) {
	n := 1000
	cols := columns.New(n)
	b.StartTimer()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		_ = cols.GoodAt3(i % n)
	}
}
