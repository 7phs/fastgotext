package fastgotext

// #cgo LDFLAGS: -L${SRCDIR}/lib -lfasttext -lstdc++
//
// #include <stdlib.h>
//
// struct WrapperFastText;
// struct WrapperDicionary;
// struct WrapperVector;
// struct WrapperFastText* FastText();
// const int FT_LoadModel(struct WrapperFastText*, const char*);
// const int FT_LoadVectors(struct WrapperFastText*, const char*);
// struct WrapperDictionary* FT_GetDictionary(struct WrapperFastText*);
// struct WrapperVector* FT_GetVector(struct WrapperFastText*, const char*);
// void FT_Release(struct WrapperFastText*);
// int DICT_Find(struct WrapperDictionary*, const char*);
// const char* DICT_GetWord(struct WrapperDictionary*, int);
// int DICT_WordsCount(struct WrapperDictionary*);
// void DICT_Release(struct WrapperDictionary*);
// float VEC_Distance(struct WrapperVector* vec1, struct WrapperVector* vec2);
// int VEC_Size(struct WrapperVector* wrapper);
// float* VEC_GetData(struct WrapperVector* wrapper);
// void VEC_Release(struct WrapperVector* wrapper);
import "C"
import (
	"bitbucket.org/7phs/fastgotext/marshal"
	"bitbucket.org/7phs/fastgotext/vector"
	"math"
	"os"
	"sort"
	"unsafe"
)

const (
	RES_OK                   ResFastText = 0
	RES_ERROR_NOT_OPEN       ResFastText = 1
	RES_ERROR_WRONG_MODEL    ResFastText = 2
	RES_ERROR_MODEL_NOT_INIT ResFastText = 3
)

type ResFastText int

func CastResFastText(err int) error {
	switch ResFastText(err) {
	case RES_OK:
		return nil
	default:
		return ResFastText(err)
	}
}

func (v ResFastText) String() string {
	switch v {
	case RES_ERROR_NOT_OPEN:
		return "failed to open"
	case RES_ERROR_WRONG_MODEL:
		return "wrong format"
	case RES_ERROR_MODEL_NOT_INIT:
		return "model wasn't init"
	case RES_OK:
		return "ok"
	default:
		return "unknown"
	}
}

func (v ResFastText) Error() string {
	return v.String()
}

type dictionary struct {
	wrapper *C.struct_WrapperDictionary
}

func (w *dictionary) Find(word string) int {
	cWord := C.CString(word)
	defer C.free(unsafe.Pointer(cWord))

	return int(C.DICT_Find(w.wrapper, cWord))
}

func (w *dictionary) GetWord(id int) string {
	return C.GoString(C.DICT_GetWord(w.wrapper, C.int(id)))
}

func (w *dictionary) WordsCount() int {
	return int(C.DICT_WordsCount(w.wrapper))
}

func (w *dictionary) Free() {
	C.DICT_Release(w.wrapper)
	w.wrapper = nil
}

type fastText struct {
	wrapper *C.struct_WrapperFastText
}

func FastText() *fastText {
	return &fastText{
		wrapper: C.FastText(),
	}
}

func (w *fastText) LoadModel(modelPath string) error {
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))

	return CastResFastText(int(C.FT_LoadModel(w.wrapper, cModelPath)))
}

func (w *fastText) LoadVectors(vectorsPath string) error {
	cVectorsPath := C.CString(vectorsPath)
	defer C.free(unsafe.Pointer(cVectorsPath))

	return CastResFastText(int(C.FT_LoadVectors(w.wrapper, cVectorsPath)))
}

func (w *fastText) GetDictionary() *dictionary {
	return &dictionary{
		wrapper: C.FT_GetDictionary(w.wrapper),
	}
}

func (w *fastText) filterDoc(doc []string) (res []string) {
	dict := w.GetDictionary()
	res = make([]string, 0, len(doc))

	for _, word := range doc {
		if dict.Find(word) > 0 {
			res = append(res, word)
		}
	}

	return
}

func (w *fastText) WordToVector(word string) []float32 {
	cWord := C.CString(word)
	defer C.free(unsafe.Pointer(cWord))

	vec := C.FT_GetVector(w.wrapper, cWord)
	defer C.VEC_Release(vec)

	return vector.UnmarshalF32(unsafe.Pointer(C.VEC_GetData(vec)), int(C.VEC_Size(vec)))
}

func (w *fastText) DocToVectors(doc []string) [][]float32 {
	res := make([][]float32, 0, len(doc))

	for _, word := range doc {
		res = append(res, w.WordToVector(word))
	}

	return res
}

func (w *fastText) WordsDistance(word1, word2 string) float32 {
	cWord1 := C.CString(word1)
	defer marshal.FreePointer(unsafe.Pointer(cWord1))

	vec1 := C.FT_GetVector(w.wrapper, cWord1)
	defer C.VEC_Release(vec1)

	cWord2 := C.CString(word2)
	defer C.free(unsafe.Pointer(cWord2))

	vec2 := C.FT_GetVector(w.wrapper, cWord2)
	defer C.VEC_Release(vec2)

	return float32(C.VEC_Distance(vec1, vec2))
}

func (w *fastText) BowNormalize(bow []int, docLen int) []float32 {
	var (
		res        = make([]float32, len(bow))
		normalizer = float32(docLen)
	)

	for wordIndex, freq := range bow {
		res[wordIndex] = float32(freq) / normalizer
	}

	return res
}

func (w *fastText) WMDistance(doc1, doc2 []string) (float32, error) {
	doc1 = w.filterDoc(doc1)
	doc2 = w.filterDoc(doc2)

	if len(doc1) == 0 || len(doc2) == 0 {
		return float32(math.Inf(1)), os.ErrInvalid
	}

	sort.Strings(doc1)
	sort.Strings(doc2)

	dict := Dictionary(Join(doc1, doc2))
	dictLen := dict.Len()
	if dictLen == 1 {
		return 1., nil
	}

	distanceMatrix := &marshal.FloatArray{}

	dict1 := Dictionary(doc1)
	dict2 := Dictionary(doc2)

	var distance float32 = .0

	for _, word1 := range dict {
		for _, word2 := range dict {
			if dict1.WordIndex(word1) < 0 || dict2.WordIndex(word2) < 0 {
				distance = .0
			} else {
				distance = w.WordsDistance(word1, word2)
			}

			distanceMatrix.Push(distance)
		}
	}

	d1 := w.BowNormalize(dict.Doc2Bow(doc1), len(doc1))
	d2 := w.BowNormalize(dict.Doc2Bow(doc2), len(doc2))

	return Emd(d1, d2, uint(distanceMatrix.Len()), distanceMatrix.Pointer()), nil
}

func (w *fastText) Similarity(doc1, doc2 []string) (float32, error) {
	doc1 = w.filterDoc(doc1)
	doc2 = w.filterDoc(doc2)

	core1, err := vector.Mean(w.DocToVectors(doc1)...)
	if err != nil {
		// TODO error wrap
		return .0, err
	}

	core2, err := vector.Mean(w.DocToVectors(doc2)...)
	if err != nil {
		// TODO error wrap
		return .0, err
	}

	return vector.Dot(core1, core2), os.ErrInvalid
}

func (w *fastText) Free() {
	C.FT_Release(w.wrapper)
	w.wrapper = nil
}
