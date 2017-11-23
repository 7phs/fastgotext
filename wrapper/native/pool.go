package native

import (
	"sync"
)

type IntPoolManager struct {
	sync.Map
}

func NewIntPoolManager() *IntPoolManager {
	return &IntPoolManager{}
}

func (o *IntPoolManager) getPool(sz uint) *intPool {
	if pool, ok := o.Load(sz); ok {
		return pool.(*intPool)
	}

	pool, _ := o.LoadOrStore(sz, newIntPool(sz))

	return pool.(*intPool)
}

func (o *IntPoolManager) Get(sz uint) *IntArray {
	return o.getPool(sz).Get()
}

type intPool struct {
	sync.Pool
}

func newIntPool(sz uint) *intPool {
	pool := &intPool{}

	pool.New = func() interface{} {
		return pool.new(sz)
	}

	return pool
}

func (o *intPool) new(sz uint) interface{} {
	return NewIntArrayExt(sz, o)
}

func (o *intPool) Get() *IntArray {
	return o.Pool.Get().(*IntArray).Clear()
}

func (o *intPool) Put(arr *IntArray) {
	o.Pool.Put(arr)
}

type floatPoolManager struct {
	sync.Map
}

func NewFloatPoolManager() *floatPoolManager {
	return &floatPoolManager{}
}

func (o *floatPoolManager) getPool(sz uint) *floatPool {
	if pool, ok := o.Load(sz); ok {
		return pool.(*floatPool)
	}

	pool, _ := o.LoadOrStore(sz, newFloatPool(sz))

	return pool.(*floatPool)
}

func (o *floatPoolManager) Get(sz uint) *FloatArray {
	return o.getPool(sz).Get()
}

type floatPool struct {
	sync.Pool
}

func newFloatPool(sz uint) *floatPool {
	pool := &floatPool{}

	pool.New = func() interface{} {
		return pool.new(sz)
	}

	return pool
}

func (o *floatPool) new(sz uint) interface{} {
	return NewFloatArrayExt(sz, o)
}

func (o *floatPool) Get() *FloatArray {
	return o.Pool.Get().(*FloatArray).Clear()
}

func (o *floatPool) Put(arr *FloatArray) {
	o.Pool.Put(arr)
}
