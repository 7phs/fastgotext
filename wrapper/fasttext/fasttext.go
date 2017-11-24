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

	"bitbucket.org/7phs/fastgotext/vector"
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

type predictResult struct {
	wrapper *C.struct_PredictResult
}

func (p *predictResult) Unmarshal() []*Predict {
	var (
		data    = C.PRDCT_Records(p.wrapper)
		len     = int(C.PRDCT_Len(p.wrapper))
		records = (*[1 << 30]C.struct_PredictRecord)(unsafe.Pointer(data))[:len:len]
		result  = make([]*Predict, 0, len)
	)

	for _, rec := range records {
		result = append(result, ToPredic(&rec))
	}

	return result
}

func (p *predictResult) HasError() error {
	if err := C.PRDCT_Error(p.wrapper); err != nil {
		return errors.New(C.GoString(err))
	}

	return nil
}

func (p *predictResult) Free() {
	C.PRDCT_Release(p.wrapper)
	p.wrapper = nil
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

func (w *fastText) WordToVector(word string) []float32 {
	cWord := C.CString(word)
	defer C.free(unsafe.Pointer(cWord))

	vec := C.FT_GetVector(w.wrapper, cWord)
	defer C.VEC_Release(vec)

	return vector.UnmarshalF32(unsafe.Pointer(C.VEC_GetData(vec)), int(C.VEC_Size(vec)))
}

func (w *fastText) Predict(text string, count int) ([]*Predict, error) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	predict := &predictResult{
		wrapper: C.FT_Predict(w.wrapper, cText, C.int(count)),
	}
	defer predict.Free()

	if err := predict.HasError(); err != nil {
		return nil, err
	}

	return predict.Unmarshal(), nil
}

func (w *fastText) Free() {
	C.FT_Release(w.wrapper)
	w.wrapper = nil
}
