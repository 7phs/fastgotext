package native

// #include <stdlib.h>
// #include "lib/src/allocator.h"
import "C"
import (
	"unsafe"
)

type Cint = C.int
type Cfloat = C.float

func ToCint(v int) C.int {
	return C.int(v)
}

func ToCfloat(v float32) C.float {
	return C.float(v)
}

type IntArray struct {
	data *C.int
	len  uint
}

func ToIntArray(data []int) *IntArray {
	arr := NewIntArray(uint(len(data)))

	slice := arr.Slice()
	for i, v := range data {
		slice[i] = C.int(v)
	}

	return arr
}

func NewIntArray(ln uint) *IntArray {
	return &IntArray{
		data: (*C.int)(C.malloc_int(C.ulong(ln))),
		len:  ln,
	}
}

func (o *IntArray) Pointer() unsafe.Pointer {
	return unsafe.Pointer(o.data)
}

func (o *IntArray) Len() uint {
	return o.len
}

func (o *IntArray) Slice() []C.int {
	if o.data == nil {
		return []C.int{}
	}

	return (*[1 << 30]C.int)(unsafe.Pointer(o.data))[:o.len:o.len]
}

func (o *IntArray) Marshal() []int {
	res := make([]int, 0, o.len)

	for _, v := range o.Slice() {
		res = append(res, int(v))
	}

	return res
}

func (o *IntArray) Free() {
	C.free_(unsafe.Pointer(o.data))

	o.data = nil
	o.len = 0
}

type FloatArray struct {
	data *C.float
	len  uint
}

func ToFloatArray(data []float32) *FloatArray {
	arr := NewFloatArray(uint(len(data)))

	slice := arr.Slice()
	for i, v := range data {
		slice[i] = C.float(v)
	}

	return arr
}

func NewFloatArray(ln uint) *FloatArray {
	return &FloatArray{
		data: (*C.float)(C.malloc_float(C.ulong(ln))),
		len:  ln,
	}
}

func (o *FloatArray) Pointer() unsafe.Pointer {
	return unsafe.Pointer(o.data)
}

func (o *FloatArray) Len() uint {
	return o.len
}

func (o *FloatArray) Slice() []C.float {
	if o.data == nil {
		return []C.float{}
	}

	return (*[1 << 30]C.float)(unsafe.Pointer(o.data))[:o.len:o.len]
}

func (o *FloatArray) Marshal() []float32 {
	res := make([]float32, 0, o.len)

	for _, v := range o.Slice() {
		res = append(res, float32(v))
	}

	return res
}

func (o *FloatArray) Free() {
	C.free_(unsafe.Pointer(o.data))

	o.data = nil
	o.len = 0
}

type FloatMatrix struct {
	data   *C.float
	lenRow uint
	lenCol uint

	slices [][]C.float
}

func ToFloatMatrix(data [][]float32) *FloatMatrix {
	var lenCol uint = 0
	if len(data) > 0 {
		lenCol = uint(len(data[0]))
	}

	arr := NewFloatMatrix(uint(len(data)), lenCol)

	slice := arr.Slice()
	for i, row := range data {
		for j, v := range row {
			slice[i][j] = C.float(v)
		}
	}

	return arr
}

func NewFloatMatrix(row, col uint) *FloatMatrix {
	return &FloatMatrix{
		data:   (*C.float)(C.malloc_float(C.ulong(row * col))),
		lenRow: row,
		lenCol: col,
	}
}

func (o *FloatMatrix) Pointer() unsafe.Pointer {
	return unsafe.Pointer(o.data)
}

func (o *FloatMatrix) LenRow() uint {
	return o.lenRow
}

func (o *FloatMatrix) LenCol() uint {
	return o.lenCol
}

func (o *FloatMatrix) Len() uint {
	return o.lenRow * o.lenCol
}

func (o *FloatMatrix) Slice() [][]C.float {
	if o.data == nil {
		return [][]C.float{}
	}

	if o.slices == nil {
		data := (*[1 << 30]C.float)(unsafe.Pointer(o.data))[: o.lenRow*o.lenCol : o.lenRow*o.lenCol]

		o.slices = make([][]C.float, o.lenRow)
		lenCol := int(o.lenCol)
		for i := range o.slices {
			o.slices[i] = data[i*lenCol : (i+1)*lenCol]
		}
	}

	return o.slices
}

func (o *FloatMatrix) Marshal() [][]float32 {
	res := make([][]float32, 0, o.lenRow)

	for _, rowV := range o.Slice() {
		row := make([]float32, 0, o.lenRow)

		for _, v := range rowV {
			row = append(row, float32(v))
		}

		res = append(res, row)
	}

	return res
}

func (o *FloatMatrix) Free() {
	C.free_(unsafe.Pointer(o.data))

	o.data = nil
	o.slices = nil
	o.lenRow = 0
	o.lenCol = 0
}
