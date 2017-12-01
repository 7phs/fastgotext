// +build !dumb

package emd

import (
	"testing"

	"bitbucket.org/7phs/fastgotext/wrapper/native"
)

func TestEmd(t *testing.T) {
	docBow1 := []float32{0.5, 0.5, 0., 0.}
	docBow2 := []float32{0., 0., 0.5, 0.5}

	distanceMatrix := [][]float32{
		{0., 0., 16.03984642, 22.11830902},
		{0., 0., 17.83054543, 14.92696762},
		{0., 0., 0., 0.},
		{0., 0., 0., 0.},
	}

	distanceMarshaled := native.ToFloatMatrix(distanceMatrix)
	defer distanceMarshaled.Free()

	exist := Emd(docBow1, docBow2, distanceMarshaled)
	expected := float32(15.483407)

	if exist != expected {
		t.Error("failed to calc emd. Got", exist, ", but expected is", expected)
	}
}
