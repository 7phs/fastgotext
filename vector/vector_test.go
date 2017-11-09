package vector

import "testing"

func TestIsEqual(t *testing.T) {
	vec1 := []float32{.1, .2, .3, .0}
	vec2 := []float32{.1, .2, .3, float32(F32_EPS_DEFAULT) / 2}
	vec3 := []float32{.4, .2, .3, .5}
	vec4 := []float32{.4, .2, .3, .5, .6}
	vec5 := []float32{.4, .2, .3}

	if !IsEqual(vec1, vec2) {
		t.Error("failed to check equal ", vec1, " and ", vec2, ", but they are equal")
	}

	if IsEqual(vec1, vec3) {
		t.Error("failed to check equal ", vec1, " and ", vec3, ", but they aren't equal")
	}

	if IsEqual(vec1, vec4) {
		t.Error("failed to check equal ", vec1, " and ", vec4, " by length, but they aren't equal")
	}

	if IsEqual(vec1, vec5) {
		t.Error("failed to check equal ", vec1, " and ", vec4, " by length, but they aren't equal")
	}
}

func TestIsEqualExt(t *testing.T) {
	eps := 0.01
	vec1 := []float32{.1, .2, .3, .0}
	vec2 := []float32{.101, .2, .3, .005}
	vec3 := []float32{.1100001, .21, .3, .005}

	if !IsEqualExt(vec1, vec2, eps) {
		t.Error("failed to check equal ", vec1, " and ", vec2, ", but they are equal with eps", eps)
	}

	if IsEqualExt(vec1, vec3, eps) {
		t.Error("failed to check isn't equal ", vec1, " and ", vec3, " with eps", eps)
	}
}

func TestMean(t *testing.T) {
	var (
		vec1 = []float32{.1, .2, .3}
		vec2 = []float32{.9, .8, .7}
		vec3 = []float32{3.5, 3.5, 3.5}
	)

	expected := []float32{1.5, 1.5, 1.5}

	exist, err := Mean(vec1, vec2, vec3)
	if err != nil {
		t.Error("failed to calc a mean of vecs", err)
		return
	}

	if !IsEqual(exist, expected) {
		t.Error("failed to calc a mean of vecs. Result is", exist, ", but expected is", expected)
	}
}

func TestDot(t *testing.T) {
	var (
		vec1             = []float32{.2, .2, .3}
		vec2             = []float32{.5, .8, .7}
		vec3             = []float32{1., 2., 3.}
		expected float32 = 1.05
	)

	if exist := Dot(vec1, vec2, vec3); F32Compare(exist, expected, F32_EPS_DEFAULT) != 0 {
		t.Error("failed to calc dot for three vec. Result is ", exist, ", but expected is", expected)
	}
}
