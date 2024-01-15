package columns_test

import (
	"testing"

	columns "github.com/qsliu2017/wrong-usage-of-unsafe"
)

func BenchmarkColumnsAt(b *testing.B) {
	n := 1000
	cols := columns.New(n)
	for _, tc := range []struct {
		atFunc func(int) columns.Column
		name   string
	}{
		{cols.BadAt1, "BadAt1"},
		{cols.BadAt2, "BadAt2"},
		{cols.GoodAt1, "GoodAt1"},
		{cols.GoodAt2, "GoodAt2"},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.atFunc(i % n)
			}
		})
	}
}
