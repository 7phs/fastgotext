package fastgotext

import (
	"strings"
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

	word := "```"
	wordIndex := 1

	dict := model.GetDictionary()

	if existWordIndex := dict.Find(word); existWordIndex != wordIndex {
		t.Error("failed to find exist word '", word, "'. Got index", existWordIndex, ", but expected", wordIndex)
	}

	wordsCount := dict.WordsCount()
	if wordsCount <= 0 {
		t.Error("wrong word count", wordsCount, ", but expected great than 0")
	} else {
		if existWord := dict.GetWord(wordIndex); strings.Compare(word, existWord) != 0 {
			t.Error("failed to find exist word by index ", wordIndex, ". Got '", existWord, "', but expected is '", 1, "'")
		}
	}
}
