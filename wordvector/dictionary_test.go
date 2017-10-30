package wordvector

import (
	"reflect"
	"testing"
)

func TestDictionary(t *testing.T) {
	words := Dictionary("word12", "word07", "word01", "word08", "word04", "word08", "word11")
	expected := dictionary{"word01", "word04", "word07", "word08", "word11", "word12"}

	if !reflect.DeepEqual(words, expected) {
		t.Error("failed to create dictionary as sorted strings array.\nResult is", words, ",\nbut expected is", expected)
	}

	words2 := Dictionary("word01", "word02", "word03", "word02", "word01")
	expected2 := dictionary{"word01", "word02", "word03"}

	if !reflect.DeepEqual(words2, expected2) {
		t.Error("failed to create dictionary by words list with repeating.\nResult is", words2, ",\nbut expected is", expected2)
	}
}

func TestDictionary_Join(t *testing.T) {
	words1 := Dictionary("word01", "word02", "word04", "word08", "word11")
	words2 := Dictionary("word03", "word04", "word06", "word07", "word09")
	words3 := Dictionary("word01", "word01", "word01")
	words4 := Dictionary("word01", "word01", "word01")

	expected := dictionary{"word01", "word02", "word03", "word04", "word06", "word07", "word08", "word09", "word11"}
	expected2 := dictionary{"word01"}

	if exist := words1.Join(words2); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to join words1 and words2.\nExist is", exist, "\n, but expected is", expected)
	}

	if exist := words2.Join(words1); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to join words2 and words1.\nExist is", exist, "\n, but expected is", expected)
	}

	if exist := words3.Join(words4); !reflect.DeepEqual(exist, expected2) {
		t.Error("failed to join one word dicts.\nExist is", exist, "\n, but expected is", expected2)
	}
}

func TestDictionary_IsEmpty(t *testing.T) {
	if !Dictionary().IsEmpty() {
		t.Error("failed to check an empty dictionary")
	}

	if Dictionary("word1", "word2").IsEmpty() {
		t.Error("failed to check a dictionary with words")
	}
}

func TestDictionary_Doc2Bow(t *testing.T) {
	doc := []string{"word08", "word01", "word02", "word04", "word08", "word08", "word11", "word04"}

	dict := dictionary{"word01", "word02", "word03", "word04", "word06", "word07", "word08", "word09", "word11"}

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

func TestDictionary_BowNormalize(t *testing.T) {
	doc := []string{"word08", "word01", "word02", "word04", "word08", "word08", "word11", "word04"}

	dict := dictionary{"word01", "word02", "word03", "word04", "word06", "word07", "word08", "word09", "word11"}

	expected := map[int]float32{0: .125, 1: .125, 3: .25, 6: .375, 8: .125}

	exist := dict.BowNormalize(doc)

	for wordIndex, existBow := range exist {
		if existBow == 0. {
			continue
		}

		if _, ok := expected[wordIndex]; !ok {
			t.Error("found unexpected word bow for ", dict[wordIndex], " = ", existBow)
		}
	}

}
