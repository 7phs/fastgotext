package wordvector

import (
	"reflect"
	"testing"

	"github.com/7phs/fastgotext/vector"
)

var (
	testWords = map[string]struct {
		index  int
		vector vector.F32Vector
	}{
		"word01": {
			index:  1,
			vector: vector.F32Vector{.1, .2, .3, .4},
		},
		"word02": {
			index:  2,
			vector: vector.F32Vector{13.1, 34.2, 82.3, 12.4},
		},
		"word03": {
			index:  3,
			vector: vector.F32Vector{-.1, -2., 3., 2.4},
		},
		"word04": {
			index:  4,
			vector: vector.F32Vector{27.1, -82., 53., .9},
		},
		"word05": {
			index:  5,
			vector: vector.F32Vector{.145, .89, .23, .89},
		},
		"word06": {
			index:  6,
			vector: vector.F32Vector{24.1, 82., 153., 19.9},
		},
		"word07": {
			index:  7,
			vector: vector.F32Vector{7.1, -8., 3., 4.6},
		},
		"word08": {
			index:  8,
			vector: vector.F32Vector{33.1, 13., -67., 123.},
		},
		"word09": {
			index:  9,
			vector: vector.F32Vector{-43.89, 155., -98., .345},
		},
		"word10": {
			index:  10,
			vector: vector.F32Vector{11.89, 1.55, -9.8, .45},
		},
	}
)

type testWordVectorDictionary struct{}

func (o *testWordVectorDictionary) Find(word string) int {
	if info, ok := testWords[word]; !ok {
		return 0
	} else {
		return info.index
	}
}

type testWordVectorModel struct{}

func (o *testWordVectorModel) GetDictionary() WordVectorDictionary {
	return &testWordVectorDictionary{}
}

func (o *testWordVectorModel) WordToVector(word string) (res vector.F32Vector) {
	if info, ok := testWords[word]; !ok {
		return nil
	} else {
		res = append(res, info.vector...)

		return
	}
}

func TestWordVector(t *testing.T) {
	model := WordVector(&testWordVectorModel{})

	if model == nil {
		t.Error("failed to create word vector model")
	}
}

func TestWordVector_filterDoc(t *testing.T) {
	model := WordVector(&testWordVectorModel{})

	test := []string{"word01", "unknown", "word03", "word12", "word08", "unknown2"}
	expected := []string{"word01", "word03", "word08"}

	exist := model.filterDoc(test)

	if !reflect.DeepEqual(exist, expected) {
		t.Error("failed to filter words. Result is ", exist, ", but expected is ", expected)
	}
}

func TestWordVector_DocToVectors(t *testing.T) {
	model := WordVector(&testWordVectorModel{})

	doc := []string{"word09", "word01", "word10", "word08", "unknown", "word09"}
	expected := [][]float32{
		{-43.89, 155., -98., .345},
		{.1, .2, .3, .4},
		{11.89, 1.55, -9.8, .45},
		{33.1, 13., -67., 123.},
		nil,
		{-43.89, 155., -98., .345},
	}

	exist := model.DocToVectors(doc)
	if !reflect.DeepEqual(exist, expected) {
		t.Error("failed to get a vectorized doc. Result is ", exist, ", but expected is ", expected)
	}
}

func TestWordVector_WordsDistance(t *testing.T) {
	model := WordVector(&testWordVectorModel{})

	var (
		exist            = model.WordsDistance("word01", "word03")
		expected float32 = 4.02119385257662
	)

	if vector.F32Compare(exist, expected, vector.F32_EPS_DEFAULT) != 0 {
		t.Error("failed to calc words distance. Result is ", exist, ", but expected is", expected)
	}
}

func TestWordVector_WMDistance(t *testing.T) {
	model := WordVector(&testWordVectorModel{})

	var (
		doc1     = []string{"word09", "word01", "word10", "word08", "unknown", "word07"}
		doc2     = []string{"word07", "word08", "word05", "unknown", "word01", "word02", "word04"}
		docEmpty = []string{"unknown14", "unknwown21", "unknown"}
		docWord1 = []string{"word01", "unknown14", "unknwown21", "unknown", "word01"}
		docWord2 = []string{"word01", "unknown", "word01"}

		expected float32 = 65.4372
	)

	if _, err := model.WMDistance(doc1, docEmpty); err == nil {
		t.Error("failed to check empty docs for calc WM distance.")
	}

	if _, err := model.WMDistance(docEmpty, doc1); err == nil {
		t.Error("failed to check empty docs (2) for calc WM distance.")
	}

	if distance, err := model.WMDistance(docWord1, docWord2); err != nil {
		t.Error("failed to check one word docs for calc WM distance. Got ", err)
	} else if vector.F32Compare(distance, 1., vector.F32_EPS_DEFAULT) != 0 {
		t.Error("failed to calc wm distance. Result is", distance, ", but expected is", 1.)
	}

	if distance, err := model.WMDistance(doc1, doc2); err != nil {
		t.Error("failed to calc WM distance. Gor error", err)
	} else if vector.F32Compare(distance, expected, vector.F32_EPS_DEFAULT) != 0 {
		t.Error("failed to calc wm distance. Result is", distance, ", but expected is", expected)
	}
}

func TestWordVector_docToUnitVec(t *testing.T) {
	model := WordVector(&testWordVectorModel{})

	var (
		doc      = []string{"word09", "word01", "word10", "word08", "unknown", "word07"}
		docEmpty = []string{"unknown01", "unknown12"}

		expected = vector.F32Vector{.030882472, .6018363, -.638114, .479218}
	)

	if _, err := model.docToUnitVec(docEmpty); err == nil {
		t.Error("failed to check empty doc")
	}

	if vec, err := model.docToUnitVec(doc); err != nil {
		t.Error("failed to calc WM distance. Gor error", err)
	} else if !vector.IsF32Equal(vec, expected) {
		t.Error("failed to get a doc unit vector. Result is", vec, ", but expected is", expected)
	}
}

func TestWordVector_Similarity(t *testing.T) {
	model := WordVector(&testWordVectorModel{})

	var (
		doc1     = []string{"word09", "word01", "word10", "word08", "unknown", "word07"}
		doc2     = []string{"word07", "word08", "word05", "unknown", "word01", "word02", "word04"}
		doc3     = []string{"word09", "word10", "word08"}
		docEmpty = []string{"unknown01", "unknown12"}

		expected1 float32 = -0.0016786456
		expected3 float32 = 0.9991214
	)

	if _, err := model.Similarity(docEmpty, doc1); err == nil {
		t.Error("failed to check empty doc as first args")
	}

	if _, err := model.Similarity(doc1, docEmpty); err == nil {
		t.Error("failed to check empty doc as second args")
	}

	if similarity, err := model.Similarity(doc1, doc2); err != nil {
		t.Error("failed to calc similarity doc1 and doc2. Gor error", err)
	} else if vector.F32Compare(similarity, expected1, vector.F32_EPS_DEFAULT) != 0 {
		t.Error("failed to calc similarity doc1 and doc2. Result is", similarity, ", but expected is", expected1)
	}

	if similarity, err := model.Similarity(doc1, doc3); err != nil {
		t.Error("failed to calc similarity doc1 and doc3. Gor error", err)
	} else if vector.F32Compare(similarity, expected3, vector.F32_EPS_DEFAULT) != 0 {
		t.Error("failed to calc similarity doc1 and doc3. Result is", similarity, ", but expected is", expected3)
	}
}
