package fastgotext

import "sort"

type Dictionary []string

func (w Dictionary) Sort() {
	sort.Strings(w)
}

func (w Dictionary) Len() int {
	return len(w)
}

func (w Dictionary) WordIndex(word string) int {
	return sort.SearchStrings(w, word)
}

func (w Dictionary) Doc2Bow(doc []string) []int {
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
