package emd

import "bitbucket.org/7phs/fastgotext/wrapper/native"

func Emd(docBow1, docBow2 []float32, distanceMatrix *native.FloatMatrix) float32 {
	var (
		signature1 = newSignatureT(docBow1)
		signature2 = newSignatureT(docBow2)
		distance   = newDistFeatureT(distanceMatrix)
	)
	defer signature1.free()
	defer signature2.free()

	return float32(emdWrapper(signature1, signature2, distance))
}
