package fastgotext

import (
	"bitbucket.org/7phs/fastgotext/vector"
	"strings"
	"testing"
)

func TestCastResFastText(t *testing.T) {
	if CastResFastText(int(RES_OK)) != nil {
		t.Error("failed to marshal", RES_OK)
	}

	if CastResFastText(int(RES_ERROR_NOT_OPEN)) == nil {
		t.Error("failed to marshal error")
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

func TestFastText_WordToVector(t *testing.T) {
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
	expected := []float32{
		-0.00028115525, -0.002524364, 0.0043603373, -0.001954828, -0.003957717, 0.0014963509, -0.0023059447, 0.0051522283, -0.005889909, 0.0010086608,
		0.00048281666, 0.003591048, -0.0026444907, 0.0019374192, -8.312254e-06, 0.0021863126, -0.005283905, -0.0038253884, 0.00011087136, 0.0006882718,
		-0.00073126255, 0.0011861239, -0.00024638764, -0.0029747693, -0.0052905437, -0.00011283379, 0.0004090964, -0.004917168, -0.0020880806, -0.00021160375,
		-0.0020581782, 0.001400112, 0.004406866, 0.0013573511, -0.0021093662, -0.0035325664, -0.0003895998, 0.0020886585, -0.0021902653, -0.0045223925,
		0.0013979502, -0.0016400857, -0.0015675564, 0.0037028338, 0.0009960134, 0.004231611, -5.1880575e-06, 0.003360684, 0.0034160935, -0.0041204593,
		-0.0009076344, -0.0005216633, -0.004358193, 0.0011702557, 0.00036213157, -0.0002627358, -0.0005607174, -0.000797773, -0.0033221773, -0.0019280508,
		-0.002110209, 0.0013653629, -0.0024630232, 0.0032884185, 0.0021542606, 0.0023555662, 2.8422353e-05, 0.0020822124, 0.0012163108, 0.0006428569,
		0.0025403555, 0.0001870439, 0.0015540848, 0.0018884136, 0.0047199456, -0.0016869125, 0.0002191428, -0.00010784376, -0.00046199455, -0.0018992046,
		-0.0014797813, 0.0005875507, 0.00279407, -0.004776223, 0.0022434208, -0.004802419, 0.00016729049, -0.0023424146, 0.001429804, 0.003817638,
		0.003346058, 0.0035668253, -0.0010157845, -0.0023282776, -0.003215957, 0.00042506386, 0.0002837536, 0.004089029, -0.00035011096, 0.0012651555,
	}

	exist := model.WordToVector(word)

	if !vector.IsEqual(exist, expected) {
		t.Error("failed to get vector of word '", word, "'")
	}
}
