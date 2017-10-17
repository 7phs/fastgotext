package fast_go_text

// #cgo LDFLAGS: -L${SRCDIR}/lib -lfasttext -lstdc++
//
// #include <stdlib.h>
//
// struct WrapperFastText;
// struct WrapperFastText* FastText();
// const int FT_LoadModel(struct WrapperFastText*, const char*);
// const int FT_LoadVectors(struct WrapperFastText*, const char*);
// void FT_Release(struct WrapperFastText*);
import "C"
import "os"

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

func (w *fastText) NearestNeighbor(word string) error {
	return os.ErrNotExist
}

func (w *fastText) Free() {
	C.FT_Release(w.wrapper)
	w.wrapper = nil
}
