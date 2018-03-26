package emd

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/7phs/fastgotext/wrapper/array"
)

func randomFloat32(probability, max float32) float32 {
	if rand.Float32() < probability {
		return .0
	}

	return rand.Float32() * max
}

func BenchmarkEmd(b *testing.B) {
	for sz := 10; sz <= 100; sz += 10 {
		b.Run(fmt.Sprint("Emd/", sz), func(b *testing.B) {
			var (
				max             float32 = 99.
				probabilityZero float32 = .3

				docBow1  = make([]float32, sz)
				docBow2  = make([]float32, sz)
				distance = make([][]float32, sz)
			)

			for i := 0; i < sz; i++ {
				docBow1[i] = randomFloat32(probabilityZero, max)
				docBow2[i] = randomFloat32(probabilityZero, max)
			}

			for i := 0; i < sz; i++ {
				distance[i] = make([]float32, sz)
			}

			for i := 0; i < sz; i++ {
				distance[i][i] = 0.

				for j := i + 1; j < sz; j++ {
					value := randomFloat32(probabilityZero, max)

					distance[i][j] = value
					distance[j][i] = value
				}
			}

			distanceMatrix := array.WithFloatMatrix(distance)
			defer distanceMatrix.Free()

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				Emd(docBow1, docBow2, distanceMatrix)
			}
		})
	}
}
