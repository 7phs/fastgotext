package emd

// #cgo LDFLAGS: -L${SRCDIR}/lib -lemd
// #include <stdlib.h>
// #include "lib/src/emd.h"
import "C"
import (
	"unsafe"

	"bitbucket.org/7phs/fastgotext/wrapper/marshal"
)

func bowToWordsWeights(docBow []float32) (C.int, unsafe.Pointer, unsafe.Pointer) {
	var (
		words           = &marshal.IntArray{}
		weights         = &marshal.FloatArray{}
		wordCount C.int = 0
	)

	for wordIndex, wordWeight := range docBow {
		if wordWeight > 0. {
			words.Push(wordIndex)
			weights.Push(wordWeight)

			wordCount++
		}
	}

	return wordCount, words.Pointer(), weights.Pointer()
}

func Emd(docBow1, docBow2 []float32, dm uint, distanceMatrix unsafe.Pointer) float32 {
	count1, words1, weights1 := bowToWordsWeights(docBow1)
	count2, words2, weights2 := bowToWordsWeights(docBow2)
	defer marshal.FreePointer(words1)
	defer marshal.FreePointer(weights1)
	defer marshal.FreePointer(words2)
	defer marshal.FreePointer(weights2)

	sign1 := &C.signature_t{
		n:        count1,
		Weights:  (*C.float)(weights1),
		Features: (*C.int)(words1),
	}

	sign2 := &C.signature_t{
		n:        count2,
		Weights:  (*C.float)(weights2),
		Features: (*C.int)(words2),
	}

	distance := &C.dist_features_t{
		dim:            (C.uint)(dm),
		distanceMatrix: (*C.float)(distanceMatrix),
	}

	res := C.emd(sign1, sign2, distance, nil, nil)

	return float32(res)
}