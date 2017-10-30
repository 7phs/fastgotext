package wordvector

import (
	"math"
	"os"

	"bitbucket.org/7phs/fastgotext/vector"
	"bitbucket.org/7phs/fastgotext/wrapper/emd"
)

type WordVectorDictionary interface {
	Find(string) int
}

type WordVectorModel interface {
	GetDictionary() WordVectorDictionary
	WordToVector(word string) vector.F32Vector
}

type wordVector struct {
	model WordVectorModel
}

func WordVector(model WordVectorModel) *wordVector {
	return &wordVector{
		model: model,
	}
}

func (w *wordVector) filterDoc(doc []string) (res []string) {
	dict := w.model.GetDictionary()
	res = make([]string, 0, len(doc))

	for _, word := range doc {
		if dict.Find(word) > 0 {
			res = append(res, word)
		}
	}

	return
}

func (w *wordVector) DocToVectors(doc []string) [][]float32 {
	res := make([][]float32, 0, len(doc))

	for _, word := range doc {
		res = append(res, w.model.WordToVector(word))
	}

	return res
}

func (w *wordVector) WordsDistance(word1, word2 string) float32 {
	vec := w.model.WordToVector(word1)

	vec.Sub(w.model.WordToVector(word2))
	vec.Pow()

	return float32(math.Sqrt(float64(vec.Sum())))
}

func (w *wordVector) WMDistance(doc1, doc2 []string) (float32, error) {
	doc1 = w.filterDoc(doc1)
	doc2 = w.filterDoc(doc2)

	dict1 := Dictionary(doc1...)
	dict2 := Dictionary(doc2...)

	if dict1.IsEmpty() || dict2.IsEmpty() {
		return float32(math.Inf(1)), os.ErrInvalid
	}

	dict := dict1.Join(dict2)
	dictLen := dict.Len()
	if dictLen <= 1 {
		return 1., nil
	}

	distanceMatrix := make([][]float32, dictLen)
	for i, word1 := range dict {
		distanceMatrix[i] = make([]float32, dictLen)

		for j, word2 := range dict {
			if dict1.WordIndex(word1) >= 0 && dict2.WordIndex(word2) >= 0 {
				distanceMatrix[i][j] = w.WordsDistance(word1, word2)
			}
		}
	}

	return emd.Emd(dict.BowNormalize(doc1), dict.BowNormalize(doc2), distanceMatrix), nil
}

func (w *wordVector) docToUnitVec(doc []string) (vector.F32Vector, error) {
	doc = w.filterDoc(doc)

	core, err := vector.Mean(w.DocToVectors(doc)...)
	if err != nil {
		// TODO error wrap
		return nil, err
	}

	coreF := vector.F32Vector(core)

	if veclen := coreF.Distance(); veclen > .0 {
		coreF.Normalize(veclen)
	}

	return coreF, nil
}

func (w *wordVector) Similarity(doc1, doc2 []string) (float32, error) {
	unitCore1, err := w.docToUnitVec(doc1)
	if err != nil {
		// TODO error wrap
		return .0, err
	}

	unitCore2, err := w.docToUnitVec(doc2)
	if err != nil {
		// TODO error wrap
		return .0, err
	}

	return vector.F32Dot(unitCore1, unitCore2), nil
}
