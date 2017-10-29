package vector

import (
	"bitbucket.org/7phs/fastgotext/marshal"
	"testing"
)

func TestF32Compare(t *testing.T) {
	testSuites := []struct {
		f1       float32
		f2       float32
		expected int
	}{
		{.1, .1, 0},
		{.2, .1, 1},
		{.1, 2., -1},
		{.1, .1 + float32(F32_EPS/2), 0},
	}

	for _, test := range testSuites {
		if exist := F32Compare(test.f1, test.f2); exist != test.expected {
			t.Error("failed to compare ", test.f1, " and ", test.f2, ". Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestCFloatToF32(t *testing.T) {
	var (
		expected = []float32{.1, 2.3, 56.7}
		arr      = (&marshal.FloatArray{}).Marshal(expected)
		src      = arr.Pointer()
	)
	defer marshal.FreePointer(src)

	exist := UnmarshalF32(src, arr.Len())

	if !IsF32Equal(exist, F32Vector(expected)) {
		t.Error("failed to convert from C.float array to []float32. Result is ", exist, ", but expected is ", expected)
	}
}

func TestIsF32Equal(t *testing.T) {
	vec1 := F32Vector{.1, .2, .3, .0}
	vec2 := F32Vector{.1, .2, .3, float32(F32_EPS) / 2}
	vec3 := F32Vector{.4, .2, .3, .5}

	if !IsF32Equal(vec1, vec2) {
		t.Error("failed to check equal ", vec1, " and ", vec2, ", but they are equal")
	}

	if IsF32Equal(vec1, vec3) {
		t.Error("failed to check equal ", vec1, " and ", vec3, ", but they aren't equal")
	}
}

func TestF32Vector_Add(t *testing.T) {
	vec1 := F32Vector{.1, .2, .3}
	vec2 := F32Vector{.4, .4, .4}

	expected := F32Vector{.5, .6, .7}

	vec1.Add(vec2)

	if !IsF32Equal(vec1, expected) {
		t.Error("failed to sum vec1 and vec2. Result is", vec1, ", but expected is", expected)
	}
}

func TestF32Vector_Mul(t *testing.T) {
	vec1 := F32Vector{.1, .2, .3}
	vec2 := F32Vector{.4, .4, .4}

	expected := F32Vector{.04, .08, .12}

	vec1.Mul(vec2)

	if !IsF32Equal(vec1, expected) {
		t.Error("failed to mul vec1 by vec2. Result is", vec1, ", but expected is", expected)
	}
}

func TestF32Vector_Sum(t *testing.T) {
	var (
		vec              = F32Vector{.1, .2, .3}
		expected float32 = .6
	)

	if exist := vec.Sum(); F32Compare(exist, expected) != 0 {
		t.Error("failed to sum vec items. Result is ", exist, ", but expected is ", expected)
	}

}

func TestF32Vector_Normalize(t *testing.T) {
	vec := F32Vector([]float32{.12, .9, .6})
	var normalizer float32 = 3

	expected := F32Vector([]float32{.04, .3, .2})

	vec.Normalize(normalizer)

	if !IsF32Equal(vec, expected) {
		t.Error("failed to normalize with ", normalizer, ". Result is", vec, ", but expected is", expected)
	}
}

func TestF32Mean(t *testing.T) {
	var (
		vec1 = F32Vector{.1, .2, .3}
		vec2 = F32Vector{.9, .8, .7}
		vec3 = F32Vector{3.5, 3.5, 3.5}
	)

	expected := F32Vector{1.5, 1.5, 1.5}

	exist, err := F32Mean(vec1, vec2, vec3)
	if err != nil {
		t.Error("failed to calc a mean of vecs", err)
		return
	}

	if !IsF32Equal(exist, expected) {
		t.Error("failed to calc a mean of vecs. Result is", exist, ", but expected is", expected)
	}
}

func TestF32Dot(t *testing.T) {
	var (
		vec1             = F32Vector{.2, .2, .3}
		vec2             = F32Vector{.5, .8, .7}
		vec3             = F32Vector{1., 2., 3.}
		expected float32 = 1.05
	)

	if exist := F32Dot(vec1, vec2, vec3); F32Compare(exist, expected) != 0 {
		t.Error("failed to calc dot for three vec. Result is ", exist, ", but expected is", expected)
	}
}
