package wordvector

import (
	"sort"
	"strings"
)

type dictionary []string

func Dictionary(words ...string) dictionary {
	return dictionary(words).Sort()
}

func (w dictionary) Join(other dictionary) (res dictionary) {
	var (
		index1 = 0
		index2 = 0

		len1 = w.Len()
		len2 = other.Len()
	)

	res = make([]string, 0, len1+len2)

	for index1 < len1 && index2 < len2 {
		switch strings.Compare(w[index1], other[index2]) {
		case 0:
			res = append(res, w[index1])

			index1++
			index2++

		case -1:
			res = append(res, w[index1])
			index1++

		case 1:
			res = append(res, other[index2])
			index2++
		}
	}

	switch {
	case index1 < len1 && index2 >= len2:
		res = append(res, w[index1:]...)

	case index2 < len1 && index1 >= len2:
		res = append(res, other[index2:]...)
	}

	return

}

func (w dictionary) Sort() dictionary {
	sort.Strings(w)

	return w
}

func (w dictionary) Len() int {
	return len(w)
}

func (w dictionary) IsEmpty() bool {
	return w.Len() == 0
}

func (w dictionary) WordIndex(word string) int {
	return sort.SearchStrings(w, word)
}

func (w dictionary) Doc2Bow(doc []string) []int {
	counter := make(map[string]int)

	for _, word := range doc {
		if _, ok := counter[word]; !ok {
			counter[word] = 1
		} else {
			counter[word]++
		}
	}

	res := make([]int, w.Len())

	for word, freq := range counter {
		if index := w.WordIndex(word); index >= 0 {
			res[index] = freq
		}
	}

	return res
}
