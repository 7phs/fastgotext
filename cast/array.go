package cast

// #include <stdlib.h>
import "C"
import (
	"bytes"
	"encoding/binary"
	"sync/atomic"
	"unsafe"
)

func FreePointer(ptr unsafe.Pointer) {
	C.free(ptr)
}

type Array struct {
	buf  bytes.Buffer
	size int
}

func (o *Array) Len() int {
	return o.size
}

func (o *Array) Pointer() unsafe.Pointer {
	return C.CBytes(o.buf.Bytes())
}

type FloatArray struct {
	Array
}

func (o *FloatArray) Push(v float32) {
	binary.Write(&o.buf, binary.LittleEndian, C.float(v))

	o.size++
}

func (o *FloatArray) Cast(array []float32) *FloatArray {
	for _, value := range array {
		o.Push(value)
	}

	return o
}

type IntArray struct {
	Array
}

func (o *IntArray) Push(v int) {
	binary.Write(&o.buf, binary.LittleEndian, C.int(v))

	o.size++
}

type FloatMatrix struct {
	rowCount int32

	FloatArray
}

func (o *FloatMatrix) RowLen() int {
	return int(o.rowCount)
}

func (o *FloatMatrix) Row() {
	atomic.AddInt32(&o.rowCount, 1)
}

func (o *FloatMatrix) Cast(matrix [][]float32) *FloatMatrix {
	for _, row := range matrix {
		o.Row()

		for _, value := range row {
			o.Push(value)
		}
	}

	return o
}
