package fastgotext

// #include <stdlib.h>
import "C"
import (
	"bytes"
	"encoding/binary"
	"strings"
	"unsafe"
)

func Join(sortedWords1, sortedWords2 []string) (res []string) {
	var (
		index1 = 0
		index2 = 0
	)

	res = make([]string, 0, len(sortedWords1)+len(sortedWords2))

	for index1 < len(sortedWords1) && index2 < len(sortedWords2) {
		switch strings.Compare(sortedWords1[index1], sortedWords2[index2]) {
		case 0:
			res = append(res, sortedWords1[index1])

			index1++
			index2++

		case -1:
			res = append(res, sortedWords1[index1])
			index1++

		case 1:
			res = append(res, sortedWords2[index2])
			index2++
		}
	}

	switch {
	case index1 < len(sortedWords1) && index2 >= len(sortedWords2):
		res = append(res, sortedWords1[index1:]...)

	case index2 < len(sortedWords2) && index1 >= len(sortedWords1):
		res = append(res, sortedWords2[index2:]...)
	}

	return
}

func CastFloatMatrix(matrix [][]float32) unsafe.Pointer {
	res := &bytes.Buffer{}

	for _, row := range matrix {
		for _, value := range row {
			binary.Write(res, binary.LittleEndian, value)
		}
	}

	return C.CBytes(res.Bytes())
}

func FreeFloatMatrix(matrix unsafe.Pointer) {
	C.free(matrix)
}
