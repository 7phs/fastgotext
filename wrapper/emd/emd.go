package emd

// #cgo LDFLAGS: -L${SRCDIR}/lib -lemd
// #include <stdlib.h>
// #include "lib/src/emd.h"
import "C"
import (
	"bitbucket.org/7phs/fastgotext/wrapper/native"
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
		words:      native.NewIntArray(uint(len(docBow))),
		weights:    native.NewFloatArray(uint(len(docBow))),
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
	o.words.Free()
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

func dumbEmd(*C.signature_t, *C.signature_t, *C.dist_features_t) float32 {
	return .0
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
	//res := dumbEmd(signature1.C, signature2.C, distance.C)

	return float32(res)
}
