package fastgotext

import (
	"fmt"
	"testing"
)

func TestCastResFastText(t *testing.T) {
	if CastResFastText(int(RES_OK)) != nil {
		t.Error("failed to cast", RES_OK)
	}

	if CastResFastText(int(RES_ERROR_NOT_OPEN)) == nil {
		t.Error("failed to cast error")
	}
}

func TestFastText(t *testing.T) {
	model := FastText()
	if model == nil {
		t.Error("failed to create fasttext model")
	} else {
		model.Free()
	}
}

func TestFastText_LoadModel(t *testing.T) {
	model := FastText()
	if model == nil {
		t.Error("failed to create fasttext model")
		return
	}
	defer model.Free()

	if err := model.LoadModel("test/model.bin"); err != nil {
		t.Error("failed to load fasttext model:", err)
	}
}

func TestFastText_LoadVector(t *testing.T) {
	model := FastText()
	if model == nil {
		t.Error("failed to create fast text vector")
		return
	}
	defer model.Free()

	if err := model.LoadVectors("test/model.vec"); err == nil {
		t.Error("load fasttext vector into uninit model")
	}

	if err := model.LoadModel("test/model.bin"); err != nil {
		t.Error("failed to load fasttext model:", err)
	}

	if err := model.LoadVectors("test/model.vec"); err != nil {
		t.Error("failed to load fasttext vector:", err)
	}
}

func TestFastText_GetDictionary(t *testing.T) {
	model := FastText()
	if model == nil {
		t.Error("failed to create fast text vector")
		return
	}
	defer model.Free()

	if err := model.LoadModel("test/model.bin"); err != nil {
		t.Error("failed to load fasttext model:", err)
		return
	}

	if err := model.LoadVectors("test/model.vec"); err != nil {
		t.Error("failed to load fasttext vector:", err)
		return
	}

	dict := model.GetDictionary()
	wordIndex := dict.Find("first")

	fmt.Println("wordIndex:", wordIndex)
}
