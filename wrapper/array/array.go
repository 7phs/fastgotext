package array

//go:generate gotemplate "bitbucket.org/7phs/native/template/array" "IntArray(int, C.int, C.sizeof_int)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "IntArrayPool(IntArray, C.sizeof_int, IntArrayPoolKey, NewIntArrayInterface)"

//go:generate gotemplate "bitbucket.org/7phs/native/template/matrix" "IntMatrix(int, C.int, C.sizeof_int)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "IntMatrixPool(IntMatrix, C.sizeof_int, IntMatrixPoolKey, NewIntMatrixInterface)"

//go:generate gotemplate "bitbucket.org/7phs/native/template/array" "FloatArray(float32, C.float, C.sizeof_float)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "FloatArrayPool(FloatArray, C.sizeof_float, FloatArrayPoolKey, NewFloatArrayInterface)"

//go:generate gotemplate "bitbucket.org/7phs/native/template/matrix" "FloatMatrix(float32, C.float, C.sizeof_float)"
//go:generate gotemplate "bitbucket.org/7phs/native/template/pool" "FloatMatrixPool(FloatMatrix, C.sizeof_float, FloatMatrixPoolKey, NewFloatMatrixInterface)"
