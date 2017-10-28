package vector

const (
	F32_EPS float64 = .0000001
)

func CastF32(vectors ...[]float32) []F32Vector {
	res := make([]F32Vector, 0, len(vectors))

	for _, vec := range vectors {
		res = append(res, F32Vector(vec))
	}

	return res
}

func IsEqual(vec1, vec2 []float32) bool {
	F32vecs := CastF32(vec1, vec2)

	return IsF32Equal(F32vecs[0], F32vecs[1])
}

func Mean(vectors ...[]float32) ([]float32, error) {
	return F32Mean(CastF32(vectors...)...)
}

func Dot(vectors ...[]float32) float32 {
	return F32Dot(CastF32(vectors...)...)
}
