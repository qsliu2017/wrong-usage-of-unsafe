package columns

import (
	"testing"
	"unsafe"
)

func TestColumnsAt(t *testing.T) {
	n := 1000
	underlying := make([]byte, n*SizeColumn)
	cols := &Columns{underlying}
	for _, tc := range []struct {
		atFunc func(int) Column
		name   string
	}{
		{cols.BadAt1, "BadAt1"},
		{cols.BadAt2, "BadAt2"},
		{cols.GoodAt1, "GoodAt1"},
		{cols.GoodAt2, "GoodAt2"},
		{cols.GoodAt3, "GoodAt3"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// re-initialize underlying data
			for i := 0; i < n; i++ {
				col := (*Column)(unsafe.Pointer(&underlying[i*SizeColumn]))
				col.x = int64(3*i + 1)
				col.y = int8(i + 1)
				col.z = 1.23
			}
			for i := 0; i < n; i++ {
				{
					// check if atFunc returns the correct value
					col := tc.atFunc(i)
					if col.x != int64(3*i+1) || col.y != int8(i+1) || col.z != 1.23 {
						t.Fatalf("col[%d] = %+v", i, col)
					}
					// change returned value
					col.x = int64(2*i + 3)
					col.y = int8(i + 3)
					col.z = 3.21
				}
				{
					// check if underlying data is changed
					col := tc.atFunc(i)
					if col.x != int64(3*i+1) || col.y != int8(i+1) || col.z != 1.23 {
						t.Fatalf("%s change underlying data", tc.name)
					}
				}
			}
		})
	}
}
