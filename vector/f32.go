package vector

import "C"
import (
	"math"
	"os"
	"unsafe"
)

func F32Compare(f1, f2 float32) int {
	res := f1 - f2

	switch {
	case math.Abs(float64(res)) < F32_EPS:
		return 0
	case res < 0:
		return -1
	case res > 0:
		return 1
	}

	return 2
}

type F32Vector []float32

func MakeF32(ln int) F32Vector {
	return make(F32Vector, ln)
}

func CopyF32(vec F32Vector) F32Vector {
	res := make(F32Vector, 0, vec.Len())

	return append(res, vec...)
}

func UnmarshalF32(ptr unsafe.Pointer, sz int) F32Vector {
	var (
		arrP    = uintptr(ptr)
		floatSz = uintptr(C.sizeof_float)
		i       uintptr
		res     = make(F32Vector, 0, sz)
	)

	for i = 0; i < uintptr(sz); i++ {
		value := unsafe.Pointer(arrP + i*floatSz)
		res = append(res, float32(*(*C.float)(value)))
	}

	return res
}

func (v F32Vector) Len() int {
	return len(v)
}

func (v *F32Vector) Add(vec F32Vector) error {
	ref := (*F32Vector)(v)

	if (*v).Len() > vec.Len() {
		return os.ErrInvalid
	}

	for i := range *ref {
		(*ref)[i] += vec[i]
	}

	return nil
}

func (v *F32Vector) Mul(vec F32Vector) error {
	ref := (*F32Vector)(v)

	if (*v).Len() > vec.Len() {
		return os.ErrInvalid
	}

	for i := range *ref {
		(*ref)[i] *= vec[i]
	}

	return nil
}

func (v F32Vector) Sum() (res float32) {
	for _, value := range v {
		res += value
	}

	return
}

func (v *F32Vector) Normalize(normalizer float32) {
	ref := (*F32Vector)(v)

	for i := range *ref {
		(*ref)[i] /= normalizer
	}
}

func IsF32Equal(vec1, vec2 F32Vector) bool {
	if vec1.Len() != vec2.Len() {
		return false
	}

	for i := 0; i < vec1.Len(); i++ {
		if F32Compare(vec1[i], vec2[i]) != 0 {
			return false
		}
	}

	return true
}

func F32Mean(vectors ...F32Vector) (F32Vector, error) {
	if len(vectors) == 0 {
		return nil, os.ErrInvalid
	}

	result := MakeF32(vectors[0].Len())

	for _, vec := range vectors {
		if err := result.Add(vec); err != nil {
			return nil, err
		}
	}

	result.Normalize(float32(result.Len()))

	return result, nil
}

func F32Dot(vectors ...F32Vector) float32 {
	if len(vectors) == 0 {
		return .0
	}

	result := CopyF32(vectors[0])

	for i := 1; i < len(vectors); i++ {
		if err := result.Mul(vectors[i]); err != nil {
			return .0
		}
	}

	return result.Sum()
}
