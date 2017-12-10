package emd

// #cgo LDFLAGS: -L${SRCDIR}/lib -lemd
// #include <stdlib.h>
// #include "lib/src/emd.h"
import "C"
import (
	"bitbucket.org/7phs/fastgotext/wrapper/array"
)

var (
	iPool = array.NewIntArrayPool()
	fPool = array.NewFloatArrayPool()
)

type signatureT struct {
	wordsCount int
	words      *array.IntArray
	weights    *array.FloatArray
	C          *C.signature_t
}

func newSignatureT(docBow []float32) *signatureT {
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
			wordsSlice[o.wordsCount] = array.IntArrayRec(wordIndex)
			weightsSlice[o.wordsCount] = array.FloatArrayRec(wordWeight)

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

func (o *signatureT) free() {
	o.words.Free()
	o.weights.Free()
}

type distFeaturesT struct {
	C *C.dist_features_t
}

func newDistFeatureT(distanceMatrix *array.FloatMatrix) *distFeaturesT {
	return &distFeaturesT{
		C: &C.dist_features_t{
			dim:            (C.uint)(distanceMatrix.Dim()[0]),
			distanceMatrix: (*C.float)(distanceMatrix.Pointer()),
		},
	}
}

type emdFunc func(*signatureT, *signatureT, *distFeaturesT) float32

func emdCalc(signature1, signature2 *signatureT, distance *distFeaturesT) float32 {
	return float32(C.emd(signature1.C, signature2.C, distance.C, nil, nil))
}

func emdDumb(signature1, signature2 *signatureT, distance *distFeaturesT) float32 {
	return float32(C.emd_dumb(signature1.C, signature2.C, distance.C, nil, nil))
}
