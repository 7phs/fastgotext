package fasttext

import (
	"strings"
	"testing"

	"bitbucket.org/7phs/fastgotext/vector"
)

const (
	testDataPath                 = "../../test-data/"
	testSupervisedDataModelBin   = testDataPath + "/supervised_model.bin"
	testSupervisedDataModelVec   = testDataPath + "/supervised_model.vec"
	testUnsupervisedDataModelBin = testDataPath + "/unsupervised_model.bin"
	testUnsupervisedDataModelVec = testDataPath + "/unsupervised_model.vec"
)

func TestCastResFastText(t *testing.T) {
	if CastResFastText(int(RES_OK)) != nil {
		t.Error("failed to marshal", RES_OK)
	}

	if CastResFastText(int(RES_ERROR_NOT_OPEN)) == nil {
		t.Error("failed to marshal error")
	}

	for _, v := range []ResFastText{
		RES_OK,
		RES_ERROR_NOT_OPEN,
		RES_ERROR_WRONG_MODEL,
		RES_ERROR_MODEL_NOT_INIT,
		ResFastText(10000000),
	} {
		if v.Error() == "" {
			t.Error("failed to implement error interface")
		}
	}
}

func TestFastText(t *testing.T) {
	model := NewFastText()
	if model == nil {
		t.Error("failed to create fasttext model")
	} else {
		model.Free()
	}
}

func TestFastText_LoadModel(t *testing.T) {
	model := NewFastText()
	if model == nil {
		t.Error("failed to create fasttext model")
		return
	}
	defer model.Free()

	if err := model.LoadModel(testUnsupervisedDataModelBin); err != nil {
		t.Error("failed to load fasttext model:", err)
	}
}

func TestFastText_LoadVector(t *testing.T) {
	model := NewFastText()
	if model == nil {
		t.Error("failed to create fast text vector")
		return
	}
	defer model.Free()

	if err := model.LoadVectors(testUnsupervisedDataModelVec); err == nil {
		t.Error("load fasttext vector into uninit model")
	}

	if err := model.LoadModel(testUnsupervisedDataModelBin); err != nil {
		t.Error("failed to load fasttext model:", err)
	}

	if err := model.LoadVectors(testUnsupervisedDataModelVec); err != nil {
		t.Error("failed to load fasttext vector:", err)
	}
}

func TestFastText_GetDictionary(t *testing.T) {
	model := NewFastText()
	if model == nil {
		t.Error("failed to create fast text vector")
		return
	}
	defer model.Free()

	if err := model.LoadModel(testUnsupervisedDataModelBin); err != nil {
		t.Error("failed to load fasttext model:", err)
		return
	}

	if err := model.LoadVectors(testUnsupervisedDataModelVec); err != nil {
		t.Error("failed to load fasttext vector:", err)
		return
	}

	word := "златом"
	wordIndex := 22

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
	model := NewFastText()
	if model == nil {
		t.Error("failed to create fast text vector")
		return
	}
	defer model.Free()

	if err := model.LoadModel(testUnsupervisedDataModelBin); err != nil {
		t.Error("failed to load fasttext model:", err)
		return
	}

	if err := model.LoadVectors(testUnsupervisedDataModelVec); err != nil {
		t.Error("failed to load fasttext vector:", err)
		return
	}

	word := "златом"
	expected := []float32{
		-0.00093162473, -0.00064210495, -0.0009707032, 0.00031974318, -0.0002315091, 0.002165826, 0.0018783867, 0.0021941606, 0.0014265366, -0.0004747169,
		0.00042756647, 0.002485942, -3.925598e-05, 0.0019237917, -0.0005152924, 0.00038866568, 0.0009962652, 1.5210654e-05, -0.0002595325, -0.00020531945,
		0.00032480282, -0.0004953152, 0.0011179452, 0.00047672325, -0.001567682, 0.0011290247, -0.001348938, 0.0013085243, 0.00067846774, -0.0006734705,
		-5.052518e-05, -0.0015279177, -0.00070518075, 0.00029840754, 0.00085332844, 0.00039165834, 0.0004000477, 0.00023921908, 0.0010118099, -0.00033629616,
		0.00089929136, 0.0023271097, -0.0025798841, 0.0011896471, 0.0010260209, -0.00046896562, -5.8911148e-05, 0.0013999777, 0.002503657, -0.00042554422,
		0.0006372212, 0.0020875847, -0.0013018411, 0.00032486717, 0.0005404552, 0.000112594964, 0.0009362643, -0.0021627878, 0.0002152612, -0.00052340183,
		-0.00093787076, -0.0017059175, 0.00025377644, 0.00017843576, -0.0014451658, -0.0015405576, 0.00087846775, 0.0005363762, 0.00039133723, 0.002590361,
		0.00010595202, 0.001551432, 0.0008010692, -0.00014902855, -0.0018039232, -0.0013080842, -0.00027898262, 0.0011848757, 0.0003816138, 0.00019682113,
		0.002425637, 0.0033935579, 0.00022636236, 0.00014700758, -0.00064028014, -0.00031854503, -0.00012771421, 0.00016985959, 0.0013882271, -0.0021626558,
		0.0010302362, 0.0005801475, -0.0024974225, -0.00091070565, -0.00015146619, -0.0018224111, -0.0011335025, -9.8327204e-05, 0.00093681796, 0.00021071793,
	}

	exist := model.WordToVector(word)

	if !vector.IsEqualExt(exist, expected, 0.0001) {
		t.Error("failed to got vector of word '", word, "'.\nGot ", exist, ",\nbut expected is ", expected)
	}
}

func TestFastText_PredictUnsipervised(t *testing.T) {
	model := NewFastText()
	if model == nil {
		t.Error("failed to create fast text vector")
		return
	}
	defer model.Free()

	if err := model.LoadModel(testUnsupervisedDataModelBin); err != nil {
		t.Error("failed to load fasttext model:", err)
		return
	}

	if err := model.LoadVectors(testUnsupervisedDataModelVec); err != nil {
		t.Error("failed to load fasttext vector:", err)
		return
	}

	_, err := model.Predict("русский", 10)
	if err == nil {
		t.Error("failed to catch error for unsupervised model")
	}
}

func TestFastText_PredictSupervised(t *testing.T) {
	model := NewFastText()
	if model == nil {
		t.Error("failed to create fast text vector")
		return
	}
	defer model.Free()

	if err := model.LoadModel(testSupervisedDataModelBin); err != nil {
		t.Error("failed to load fasttext model:", err)
		return
	}

	if err := model.LoadVectors(testSupervisedDataModelVec); err != nil {
		t.Error("failed to load fasttext vector:", err)
		return
	}

	// TODO: check result
	expected := []*Predict{
		{Probability: -1.1025261, Word: "вопрос"},
		{Probability: -1.1025261, Word: "пожелание"},
		{Probability: -1.1025261, Word: "приветствие"},
	}

	result, err := model.Predict("когда вы уедите?", 10)
	if err != nil {
		t.Error("failed to make a prediction with error", err)
	}

	exist := make(map[string]bool)
	for _, predict := range result {
		for _, expected_result := range expected {
			if strings.Compare(predict.Word, expected_result.Word) == 0 {
				exist[predict.Word] = true

				break
			}
		}
	}

	if len(exist) != len(expected) {
		t.Error("failed to make a prediction. Got \n", result, "\n, but expected is\n", expected)
	}
}
