package columns

import "unsafe"

type Column struct {
	x int64
	y int8
	z float64
}

const SizeColumn = int(unsafe.Sizeof(Column{}))

type Columns struct {
	data []byte
}

func New(n int) *Columns {
	return &Columns{make([]byte, n*SizeColumn)}
}

func (c *Columns) BadAt1(i int) Column {
	data := c.data[i*SizeColumn : (i+1)*SizeColumn]
	return *(*Column)(unsafe.Pointer(&data[0]))
}

func (c *Columns) BadAt2(i int) Column {
	return *(*Column)(unsafe.Pointer(&c.data[i*SizeColumn]))
}

func (c *Columns) GoodAt1(i int) Column {
	var col [SizeColumn]byte
	copy(col[:], c.data[i*SizeColumn:(i+1)*SizeColumn])
	return *(*Column)(unsafe.Pointer(&col))
}

func (c *Columns) GoodAt2(i int) Column {
	a := [SizeColumn]byte(c.data[i*SizeColumn : (i+1)*SizeColumn])
	return *(*Column)(unsafe.Pointer(&a))
}

func (c *Columns) GoodAt3(i int) Column {
	p := (*[SizeColumn]byte)(c.data[i*SizeColumn : (i+1)*SizeColumn])
	return *(*Column)(unsafe.Pointer(p))
}
