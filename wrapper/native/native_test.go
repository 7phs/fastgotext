package native

import (
	"reflect"
	"testing"
)

func TestToIntArray(t *testing.T) {
	expected := []int{0, 2, 4, 6, 8}

	d := ToIntArray(expected)

	if exist := d.Marshal(); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to make copy as native array. Got", exist, ", but expected is", expected)
	}
}

func TestIntArray(t *testing.T) {
	len := uint(5)
	d := NewIntArray(len)
	defer d.Free()

	data := d.Slice()

	for i := range data {
		data[i] = ToCint(int(i) * 2)
	}

	expected := []int{0, 2, 4, 6, 8}

	if exist := d.Marshal(); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to change native array. Got", exist, ", but expected is", expected)
	}
}

func TestIntArray_Free(t *testing.T) {
	len := uint(5)
	d := NewIntArray(len)
	d.Free()

	count := 0
	for i := range d.Slice() {
		count += i
	}

	if count != 0 {
		t.Error("failed to range empty slices")
	}
}

func TestToFloatArray(t *testing.T) {
	expected := []float32{0, 2.904, 5.808, 8.712, 11.616, 14.52}

	d := ToFloatArray(expected)

	if exist := d.Marshal(); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to make copy as native array. Got", exist, ", but expected is", expected)
	}
}

func TestFloatArray(t *testing.T) {
	len := uint(6)
	d := NewFloatArray(len)
	defer d.Free()

	data := d.Slice()

	for i := range data {
		data[i] = ToCfloat(float32(i) * float32(2.904))
	}

	expected := []float32{0, 2.904, 5.808, 8.712, 11.616, 14.52}

	if exist := d.Marshal(); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to change native array. Got", exist, ", but expected is", expected)
	}
}

func TestFloatArray_Free(t *testing.T) {
	len := uint(7)
	d := NewFloatArray(len)
	d.Free()

	count := 0
	for i := range d.Slice() {
		count += i
	}

	if count != 0 {
		t.Error("failed to range empty slices")
	}
}

func TestToFloatMatrix(t *testing.T) {
	expected := [][]float32{
		{0., 0., 0., 0., 0.},
		{1.04, 0.52, 0.34666666, 0.26, 0.20799999},
		{2.08, 1.04, 0.6933333, 0.52, 0.41599998},
		{3.12, 1.56, 1.04, 0.78, 0.624},
	}

	d := ToFloatMatrix(expected)

	if exist := d.Marshal(); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to make copy as native matrix. Got", exist, ", but expected is", expected)
	}
}

func TestFloatMatrix(t *testing.T) {
	lenRow := uint(4)
	lenCol := uint(5)
	d := NewFloatMatrix(lenRow, lenCol)
	defer d.Free()

	data := d.Slice()

	for i, row := range data {
		for j := range row {
			data[i][j] = ToCfloat(float32(i) / float32(j+1) * float32(1.04))
		}
	}

	expected := [][]float32{
		{0., 0., 0., 0., 0.},
		{1.04, 0.52, 0.34666666, 0.26, 0.20799999},
		{2.08, 1.04, 0.6933333, 0.52, 0.41599998},
		{3.12, 1.56, 1.04, 0.78, 0.624},
	}

	if exist := d.Marshal(); !reflect.DeepEqual(exist, expected) {
		t.Error("failed to change native array. Got", exist, ", but expected is", expected)
	}
}

func TestFloatMatrix_Free(t *testing.T) {
	lenRow := uint(4)
	lenCol := uint(5)
	d := NewFloatMatrix(lenRow, lenCol)
	d.Free()

	count := 0
	for i, row := range d.Slice() {
		for j := range row {
			count += i + j
		}
	}

	if count != 0 {
		t.Error("failed to range empty slices")
	}
}
