package fastgotext

import (
	"testing"
)

func TestDictionary_Doc2Bow(t *testing.T) {
	doc := []string{"word08", "word01", "word02", "word04", "word08", "word08", "word11", "word04"}

	dict := Dictionary{"word01", "word02", "word03", "word04", "word06", "word07", "word08", "word09", "word11"}

	expected := map[int]int{0: 1, 1: 1, 3: 2, 6: 3, 8: 1}

	exist := dict.Doc2Bow(doc)

	for wordIndex, expectedFreq := range expected {
		if wordIndex < 0 || wordIndex >= len(exist) {
			t.Error("failed to calc word freq. Index ", wordIndex, " is out of exist scope")
			continue
		}
		if existFreq := exist[wordIndex]; existFreq != expectedFreq {
			t.Error("failed to calc word freq for ", dict[wordIndex], ". Exist is", existFreq, ", but expected is ", expectedFreq)
		}
	}

	for wordIndex, existFreq := range exist {
		if existFreq == 0 {
			continue
		}

		if _, ok := expected[wordIndex]; !ok {
			t.Error("found unexpected word freq for ", dict[wordIndex], " = ", existFreq)
		}
	}
}
