package fasttext

// #cgo LDFLAGS: -L${SRCDIR}/lib -lfasttext -lstdc++
//
// #include <stdlib.h>
//
// struct PredictRecord {
//     float       predict;
//     const char* word;
// };
//
// struct WrapperFastText* FastText();
// const int FT_LoadModel(struct WrapperFastText*, const char*);
// const int FT_LoadVectors(struct WrapperFastText*, const char*);
// struct WrapperDictionary* FT_GetDictionary(struct WrapperFastText*);
// struct WrapperVector* FT_GetVector(struct WrapperFastText*, const char*);
// struct PredictResult* FT_Predict(struct WrapperFastText*, const char*, int);
// void FT_Release(struct WrapperFastText*);
// const int DICT_Find(struct WrapperDictionary*, const char*);
// const char* DICT_GetWord(struct WrapperDictionary*, int);
// const int DICT_WordsCount(struct WrapperDictionary*);
// const int VEC_Size(struct WrapperVector*);
// const float* VEC_GetData(struct WrapperVector*);
// void VEC_Release(struct WrapperVector*);
// const int PRDCT_Len(struct PredictResult*);
// const char* PRDCT_Error(struct PredictResult*);
// struct PredictRecord* PRDCT_Records(struct PredictResult* result);
// void PRDCT_Release(struct PredictResult*);
import "C"
import (
	"errors"
	"fmt"
	"strings"
	"unsafe"

	"github.com/7phs/fastgotext/vector"
)

const (
	RES_OK ResFastText = iota
	RES_ERROR_NOT_OPEN
	RES_ERROR_WRONG_MODEL
	RES_ERROR_MODEL_NOT_INIT
	RES_ERROR_EXECUTION
)

const (
	FASTTEXT_LABEL_PREFIX = "__label__"
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
	case RES_ERROR_EXECUTION:
		return "execution error"
	case RES_OK:
		return "ok"
	default:
		return "unknown"
	}
}

func (v ResFastText) Error() string {
	return v.String()
}

type Dictionary interface {
	Find(word string) int
	GetWord(id int) string
	WordsCount() int
}

type dictionary = C.struct_WrapperDictionary

func (o *dictionary) Find(word string) int {
	cWord := C.CString(word)
	defer C.free(unsafe.Pointer(cWord))

	return int(C.DICT_Find(o, cWord))
}

func (o *dictionary) GetWord(id int) string {
	return C.GoString(C.DICT_GetWord(o, C.int(id)))
}

func (o *dictionary) WordsCount() int {
	return int(C.DICT_WordsCount(o))
}

type Predict struct {
	Probability float32
	Word        string
}

func (o *Predict) String() string {
	return fmt.Sprint("word: '", o.Word, "'; probability: ", o.Probability)
}

func ToPredic(rec *C.struct_PredictRecord) *Predict {
	return &Predict{
		Probability: float32(rec.predict),
		Word:        strings.TrimPrefix(C.GoString(rec.word), FASTTEXT_LABEL_PREFIX),
	}
}

type predictResult = C.struct_PredictResult

func (o *predictResult) Unmarshal() []*Predict {
	var (
		data    = C.PRDCT_Records(o)
		len     = int(C.PRDCT_Len(o))
		records = (*[1 << 30]C.struct_PredictRecord)(unsafe.Pointer(data))[:len:len]
		result  = make([]*Predict, 0, len)
	)

	for _, rec := range records {
		result = append(result, ToPredic(&rec))
	}

	return result
}

func (o *predictResult) HasError() error {
	if err := C.PRDCT_Error(o); err != nil {
		return errors.New(C.GoString(err))
	}

	return nil
}

func (o *predictResult) Free() {
	C.PRDCT_Release(o)
}

type FastText interface {
	LoadModel(modelPath string) error
	LoadVectors(vectorsPath string) error
	GetDictionary() Dictionary
	WordToVector(word string) []float32
	Predict(text string, count int) ([]*Predict, error)
	Free()
}

type fastText = C.struct_WrapperFastText

func NewFastText() FastText {
	return C.FastText()
}

func (o *fastText) LoadModel(modelPath string) error {
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))

	return CastResFastText(int(C.FT_LoadModel(o, cModelPath)))
}

func (o *fastText) LoadVectors(vectorsPath string) error {
	cVectorsPath := C.CString(vectorsPath)
	defer C.free(unsafe.Pointer(cVectorsPath))

	return CastResFastText(int(C.FT_LoadVectors(o, cVectorsPath)))
}

func (o *fastText) GetDictionary() Dictionary {
	return C.FT_GetDictionary(o)
}

func (o *fastText) WordToVector(word string) []float32 {
	cWord := C.CString(word)
	defer C.free(unsafe.Pointer(cWord))

	vec := C.FT_GetVector(o, cWord)
	defer C.VEC_Release(vec)

	return vector.UnmarshalF32(unsafe.Pointer(C.VEC_GetData(vec)), int(C.VEC_Size(vec)))
}

func (o *fastText) Predict(text string, count int) ([]*Predict, error) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	predict := C.FT_Predict(o, cText, C.int(count))
	defer predict.Free()

	if err := predict.HasError(); err != nil {
		return nil, err
	}

	return predict.Unmarshal(), nil
}

func (o *fastText) Free() {
	C.FT_Release(o)
}
