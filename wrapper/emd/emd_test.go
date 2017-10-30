package emd

import (
	"testing"
)

func TestEmd(t *testing.T) {
	docBow1 := []float32{0.1, 0., 1.2, 0., 0., 2.4}
	docBow2 := []float32{0, 0.4, 0., 2., 0., 0.}

	distanceMatrix := [][]float32{
		{0., 1., 2., 1.5, 5.6, 3.},
		{1., 0., 12., .5, .6, 13.},
		{2., 12., 0., 1.5, 6.5, 8.},
		{1.5, .5, 1.5, 0., .66, 33.},
		{5.6, .6, 6.5, .66, 0., 3.},
		{3., 13., 8., 33., 3., 0.},
	}

	exist := Emd(docBow1, docBow2, distanceMatrix)
	expected := float32(12.604165)

	if exist != expected {
		t.Error("failed to calc emd. Got", exist, ", but expected is", expected)
	}
}
