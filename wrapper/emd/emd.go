package emd

// #cgo LDFLAGS: -L${SRCDIR}/lib -lemd
// #include <stdlib.h>
// #include "lib/src/emd.h"
import "C"
import (
	"bitbucket.org/7phs/fastgotext/wrapper/native"
)

var (
	iPool = native.NewIntPoolManager()
	fPool = native.NewFloatPoolManager()
)

type signatureT struct {
	wordsCount int
	words      *native.IntArray
	weights    *native.FloatArray
	C          *C.signature_t
}

func NewSignatureT(docBow []float32) *signatureT {
	return (&signatureT{
		wordsCount: 0,
		words:      iPool.Get(uint(len(docBow))),
		weights:    fPool.Get(uint(len(docBow))),
	}).init(docBow)
}

func (o *signatureT) init(docBow []float32) *signatureT {
	var (
		wordsSlice   = o.words.Slice()
		weightsSlice = o.weights.Slice()
	)

	for wordIndex, wordWeight := range docBow {
		if wordWeight > 0. {
			wordsSlice[o.wordsCount] = (native.Cint)(wordIndex)
			weightsSlice[o.wordsCount] = (native.Cfloat)(wordWeight)

			o.wordsCount++
		}
	}

	o.C = &C.signature_t{
		n:        C.int(o.wordsCount),
		Weights:  (*C.float)(o.weights.Pointer()),
		Features: (*C.int)(o.words.Pointer()),
	}

	return o
}

func (o *signatureT) Free() {
	o.words.Put()
	o.weights.Put()
}

type distFeaturesT struct {
	C *C.dist_features_t
}

func NewDistFeatureT(distanceMatrix *native.FloatMatrix) *distFeaturesT {
	return &distFeaturesT{
		C: &C.dist_features_t{
			dim:            (C.uint)(distanceMatrix.LenRow()),
			distanceMatrix: (*C.float)(distanceMatrix.Pointer()),
		},
	}
}

func Emd(docBow1, docBow2 []float32, distanceMatrix *native.FloatMatrix) float32 {
	var (
		signature1 = NewSignatureT(docBow1)
		signature2 = NewSignatureT(docBow2)
		distance   = NewDistFeatureT(distanceMatrix)
	)
	defer signature1.Free()
	defer signature2.Free()

	res := C.emd(signature1.C, signature2.C, distance.C, nil, nil)
	// res := C.emd_dumb(signature1.C, signature2.C, distance.C, nil, nil)

	return float32(res)
}
