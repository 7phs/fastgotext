package fasttext

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
// int VEC_Size(struct WrapperVector* wrapper);
// float* VEC_GetData(struct WrapperVector* wrapper);
// void VEC_Release(struct WrapperVector* wrapper);
import "C"
import (
	"unsafe"

	"bitbucket.org/7phs/fastgotext/vector"
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

func (w *fastText) Free() {
	C.FT_Release(w.wrapper)
	w.wrapper = nil
}
