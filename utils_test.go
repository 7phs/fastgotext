package fastgotext

import (
	"testing"
	"reflect"
)

func TestJoin(t *testing.T) {
	words1:=[]string{"word01", "word02", "word04", "word08", "word11"}
	words2:=[]string{"word03", "word04", "word06", "word07", "word09"}

	expected := []string{"word01", "word02", "word03", "word04", "word06", "word07", "word08", "word09", "word11"}

	if exist := Join(words1, words2); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to join words1 and words2.\nExist is", exist, "\n, but expected is", expected)
	}

	if exist := Join(words2, words1); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to join words2 and words1.\nExist is", exist, "\n, but expected is", expected)
	}
}