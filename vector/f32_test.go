package vector

import (
	"math"
	"testing"

	"github.com/7phs/fastgotext/wrapper/array"
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
		{.1, .1 + float32(F32_EPS_DEFAULT/2), 0},
	}

	for _, test := range testSuites {
		if exist := F32Compare(test.f1, test.f2, F32_EPS_DEFAULT); exist != test.expected {
			t.Error("failed to compare ", test.f1, " and ", test.f2, ". Got ", exist, ", but expected is ", test.expected)
		}
	}
}

func TestCFloatToF32(t *testing.T) {
	var (
		expected = []float32{.1, 2.3, 56.7}
		arr      = array.WithFloatArray(expected)
		src      = arr.Pointer()
	)
	defer arr.Free()

	exist := UnmarshalF32(src, int(arr.Dim()[0]))

	if !IsF32Equal(exist, F32Vector(expected)) {
		t.Error("failed to convert from C.float array to []float32. Result is ", exist, ", but expected is ", expected)
	}
}

func TestIsF32Equal(t *testing.T) {
	vec1 := F32Vector{.1, .2, .3, .0}
	vec2 := F32Vector{.1, .2, .3, float32(F32_EPS_DEFAULT) / 2}
	vec3 := F32Vector{.4, .2, .3, .5}
	vec4 := F32Vector{.4, .2, .3, .5, .6}
	vec5 := F32Vector{.4, .2, .3}

	if !IsF32Equal(vec1, vec2) {
		t.Error("failed to check equal ", vec1, " and ", vec2, ", but they are equal")
	}

	if IsF32Equal(vec1, vec3) {
		t.Error("failed to check equal ", vec1, " and ", vec3, ", but they aren't equal")
	}

	if IsF32Equal(vec1, vec4) {
		t.Error("failed to check equal ", vec1, " and ", vec4, " by length, but they aren't equal")
	}

	if IsF32Equal(vec1, vec5) {
		t.Error("failed to check equal ", vec1, " and ", vec4, " by length, but they aren't equal")
	}
}

func TestIsF32EqualExt(t *testing.T) {
	eps := 0.01
	vec1 := F32Vector{.1, .2, .3, .0}
	vec2 := F32Vector{.101, .2, .3, .005}
	vec3 := F32Vector{.1100001, .21, .3, .005}

	if !IsF32EqualExt(vec1, vec2, eps) {
		t.Error("failed to check equal ", vec1, " and ", vec2, ", but they are equal with eps", eps)
	}

	if IsF32EqualExt(vec1, vec3, eps) {
		t.Error("failed to check isn't equal ", vec1, " and ", vec3, " with eps", eps)
	}
}

func TestF32Vector_Add(t *testing.T) {
	vec1 := F32Vector{.1, .2, .3}
	vec2 := F32Vector{.4, .4, .4}
	vec3 := F32Vector{.4, .4}

	expected := F32Vector{.5, .6, .7}

	vec1.Add(vec2)

	if !IsF32Equal(vec1, expected) {
		t.Error("failed to sum vec1 and vec2. Result is", vec1, ", but expected is", expected)
	}

	if err := vec1.Add(vec3); err == nil {
		t.Error("failed to check sum vec1 and vec3. Result is", err, ", but expected is error")
	}
}

func TestF32Vector_Sub(t *testing.T) {
	vec1 := F32Vector{.1, .2, .3}
	vec2 := F32Vector{.4, .4, .4}
	vec3 := F32Vector{.4, .4}

	expected := F32Vector{.3, .2, .1}

	vec2.Sub(vec1)

	if !IsF32Equal(vec2, expected) {
		t.Error("failed to subtract vec2 and vec1. Result is", vec2, ", but expected is", expected)
	}

	if err := vec2.Sub(vec3); err == nil {
		t.Error("failed to check subtract vec1 and vec3. Result is", err, ", but expected is error")
	}
}

func TestF32Vector_Mul(t *testing.T) {
	vec1 := F32Vector{.1, .2, .3}
	vec2 := F32Vector{.4, .4, .4}
	vec3 := F32Vector{.4, .4}

	expected := F32Vector{.04, .08, .12}

	vec1.Mul(vec2)

	if !IsF32Equal(vec1, expected) {
		t.Error("failed to mul vec1 by vec2. Result is", vec1, ", but expected is", expected)
	}

	if err := vec1.Mul(vec3); err == nil {
		t.Error("failed to check mul vec1 and vec3. Result is", err, ", but expected is error")
	}
}

func TestF32Vector_Sum(t *testing.T) {
	var (
		vec              = F32Vector{.1, .2, .3}
		expected float32 = .6
	)

	if exist := vec.Sum(); F32Compare(exist, expected, F32_EPS_DEFAULT) != 0 {
		t.Error("failed to sum vec items. Result is ", exist, ", but expected is ", expected)
	}

}

func TestF32Vector_Pow(t *testing.T) {
	vec := F32Vector{.1, .2, .3}

	expected := F32Vector{.01, .04, .09}

	vec.Pow()

	if !IsF32Equal(vec, expected) {
		t.Error("failed to subtract vec2 and vec1. Result is", vec, ", but expected is", expected)
	}
}

func TestF32Vector_Distance(t *testing.T) {
	vec := F32Vector{.1, .2, .3}

	expected := float32(math.Sqrt(.14))

	if exist := vec.Distance(); exist != expected {
		t.Error("failed to calc unified vec distance. Result is ", exist, ", but expected is ", expected)
	}
}

func TestF32Vector_Normalize(t *testing.T) {
	vec := F32Vector{.12, .9, .6}
	var normalizer float32 = 3

	expected := F32Vector{.04, .3, .2}

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
		vec4 = F32Vector{1., 2., 3., .4}
	)

	expected := F32Vector{1.5, 1.5, 1.5}

	if _, err := F32Mean(); err == nil {
		t.Error("failed to check empty args mean")
	}

	if exist, err := F32Mean(vec1, vec2, vec3); err != nil {
		t.Error("failed to calc a mean of vecs", err)
	} else if !IsF32Equal(exist, expected) {
		t.Error("failed to calc a mean of vecs. Result is", exist, ", but expected is", expected)
	}

	if exist, err := F32Mean(vec1, vec2, vec3, vec4); err == nil {
		t.Error("failed to check vec with different length. Result is ", err, " and ", exist, ", but expected is error and .0")
	}
}

func TestF32Dot(t *testing.T) {
	var (
		vec1             = F32Vector{.2, .2, .3}
		vec2             = F32Vector{.5, .8, .7}
		vec3             = F32Vector{1., 2., 3.}
		vec4             = F32Vector{1., 2.}
		expected float32 = 1.05
	)

	if v := F32Dot(); v != .0 {
		t.Error("failed to check empty args mean")
	}

	if exist := F32Dot(vec1, vec2, vec3); F32Compare(exist, expected, F32_EPS_DEFAULT) != 0 {
		t.Error("failed to calc dot for three vec. Result is ", exist, ", but expected is", expected)
	}

	if exist := F32Dot(vec1, vec2, vec3, vec4); exist != .0 {
		t.Error("failed to check vec with different length. Result is ", exist, ", but expected is .0")
	}
}
