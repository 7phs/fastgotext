package array

//go:generate gotemplate "bitbucket.org/7phs/native/template/array" "IntArray(int, C.int, C.sizeof_int)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "IntArrayPool(IntArray, C.sizeof_int, NewIntArrayInterface)"

//go:generate gotemplate "bitbucket.org/7phs/native/template/matrix" "IntMatrix(int, C.int, C.sizeof_int)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "IntMatrixPool(IntMatrix, C.sizeof_int, NewIntMatrixInterface)"

//go:generate gotemplate "bitbucket.org/7phs/native/template/array" "FloatArray(float32, C.float, C.sizeof_float)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "FloatArrayPool(FloatArray, C.sizeof_float, NewFloatArrayInterface)"

//go:generate gotemplate "bitbucket.org/7phs/native/template/matrix" "FloatMatrix(float32, C.float, C.sizeof_float)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "FloatMatrixPool(FloatMatrix, C.sizeof_float, NewFloatMatrixInterface)"
