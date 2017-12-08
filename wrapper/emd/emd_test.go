package emd

import (
	"testing"

	"bitbucket.org/7phs/fastgotext/wrapper/native"
)

func prepareData() (docBow1, docBow2 []float32, distanceMatrix *native.FloatMatrix) {
	docBow1 = []float32{0.5, 0.5, 0., 0.}
	docBow2 = []float32{0., 0., 0.5, 0.5}

	distanceMatrix = native.ToFloatMatrix([][]float32{
		{0., 0., 16.03984642, 22.11830902},
		{0., 0., 17.83054543, 14.92696762},
		{0., 0., 0., 0.},
		{0., 0., 0., 0.},
	})

	return
}

func TestEmd(t *testing.T) {
	docBow1, docBow2, distanceMatrix := prepareData()
	defer distanceMatrix.Free()

	emdWrapperSaved := emdWrapper
	defer func() {
		emdWrapper = emdWrapperSaved
	}()

	testSuites := []struct {
		t string
		f emdFunc
		e float32
	}{
		{t: "emd general", f: emdCalc, e: 15.483407},
		{t: "emd just wrapper", f: emdDumb, e: 1.},
	}

	for _, testCase := range testSuites {
		emdWrapper = testCase.f

		if exist := Emd(docBow1, docBow2, distanceMatrix); exist != testCase.e {
			t.Error("failed to calc emd in a '", testCase.t, "' case. Got", exist, ", but expected is", testCase.e)
		}
	}
}
