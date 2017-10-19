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
// struct WrapperVector* FT_GetVector(struct WrapperFastText* wrapper, const char* word);
// void FT_Release(struct WrapperFastText*);
// int DICT_Find(struct WrapperDictionary*, const char*);
// void DICT_Release(struct WrapperDictionary*);
// float VEC_Distance(struct WrapperVector* vec1, struct WrapperVector* vec2);
// void VEC_Release(struct WrapperVector* wrapper);
import "C"
import (
	"math"
	"os"
	"sort"
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
	return int(C.DICT_Find(w.wrapper, C.CString(word)))
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
	return CastResFastText(int(C.FT_LoadModel(w.wrapper, C.CString(modelPath))))
}

func (w *fastText) LoadVectors(vectorsPath string) error {
	return CastResFastText(int(C.FT_LoadVectors(w.wrapper, C.CString(vectorsPath))))
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

func (w *fastText) WMDistance(doc1 []string, doc2 []string) (float64, error) {
	doc1 = w.filterDoc(doc1)
	doc2 = w.filterDoc(doc2)

	if len(doc1) == 0 || len(doc2) == 0 {
		return math.Inf(1), os.ErrInvalid
	}

	sort.Strings(doc1)
	sort.Strings(doc2)

	wordsSet := Join(doc1, doc2)
	vocabLen := len(wordsSet)
	if vocabLen == 1 {
		return 1., nil
	}

	distanceMatrix := make([][]float64, vocabLen*vocabLen)

	for i, word1 := range wordsSet {
		for j, word2 := range wordsSet {
			if sort.SearchStrings(doc1, word1) < 0 || sort.SearchStrings(doc2, word2) < 0 {
				continue
			}

			func() {
				vec1 := C.FT_GetVector(w.wrapper, C.CString(word1))
				vec2 := C.FT_GetVector(w.wrapper, C.CString(word1))
				defer C.VEC_Release(vec1)
				defer C.VEC_Release(vec2)

				distanceMatrix[i][j] = float64(C.VEC_Distance(vec1, vec2))
			}()
		}
	}

	return 0., os.ErrInvalid
}

func (w *fastText) Free() {
	C.FT_Release(w.wrapper)
	w.wrapper = nil
}
