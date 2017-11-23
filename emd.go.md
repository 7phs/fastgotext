/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: emd.i

package emd

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_emd_f9fc16b73f1936b0(swig_intgo arg1);
extern void _wrap_signature_t_n_set_emd_f9fc16b73f1936b0(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_signature_t_n_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern void _wrap_signature_t_Features_set_emd_f9fc16b73f1936b0(uintptr_t arg1, swig_voidp arg2);
extern swig_voidp _wrap_signature_t_Features_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern void _wrap_signature_t_Weights_set_emd_f9fc16b73f1936b0(uintptr_t arg1, swig_voidp arg2);
extern swig_voidp _wrap_signature_t_Weights_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern uintptr_t _wrap_new_signature_t_emd_f9fc16b73f1936b0(void);
extern void _wrap_delete_signature_t_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern void _wrap_flow_t_from_set_emd_f9fc16b73f1936b0(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_flow_t_from_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern void _wrap_flow_t_to_set_emd_f9fc16b73f1936b0(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_flow_t_to_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern void _wrap_flow_t_amount_set_emd_f9fc16b73f1936b0(uintptr_t arg1, float arg2);
extern float _wrap_flow_t_amount_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern uintptr_t _wrap_new_flow_t_emd_f9fc16b73f1936b0(void);
extern void _wrap_delete_flow_t_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern void _wrap_dist_features_t_dim_set_emd_f9fc16b73f1936b0(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_dist_features_t_dim_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern void _wrap_dist_features_t_distanceMatrix_set_emd_f9fc16b73f1936b0(uintptr_t arg1, swig_voidp arg2);
extern swig_voidp _wrap_dist_features_t_distanceMatrix_get_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern uintptr_t _wrap_new_dist_features_t_emd_f9fc16b73f1936b0(void);
extern void _wrap_delete_dist_features_t_emd_f9fc16b73f1936b0(uintptr_t arg1);
extern float _wrap_emd_emd_f9fc16b73f1936b0(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3, uintptr_t arg4, swig_voidp arg5);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_emd_f9fc16b73f1936b0(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrSignature_t uintptr

func (p SwigcptrSignature_t) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSignature_t) SwigIsSignature_t() {
}

func (arg1 SwigcptrSignature_t) SetN(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_signature_t_n_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrSignature_t) GetN() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_signature_t_n_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSignature_t) SetFeatures(arg2 *int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_signature_t_Features_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (arg1 SwigcptrSignature_t) GetFeatures() (_swig_ret *int) {
	var swig_r *int
	_swig_i_0 := arg1
	swig_r = (*int)(C._wrap_signature_t_Features_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSignature_t) SetWeights(arg2 *float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_signature_t_Weights_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (arg1 SwigcptrSignature_t) GetWeights() (_swig_ret *float32) {
	var swig_r *float32
	_swig_i_0 := arg1
	swig_r = (*float32)(C._wrap_signature_t_Weights_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewSignature_t() (_swig_ret Signature_t) {
	var swig_r Signature_t
	swig_r = (Signature_t)(SwigcptrSignature_t(C._wrap_new_signature_t_emd_f9fc16b73f1936b0()))
	return swig_r
}

func DeleteSignature_t(arg1 Signature_t) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_signature_t_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0))
}

type Signature_t interface {
	Swigcptr() uintptr
	SwigIsSignature_t()
	SetN(arg2 int)
	GetN() (_swig_ret int)
	SetFeatures(arg2 *int)
	GetFeatures() (_swig_ret *int)
	SetWeights(arg2 *float32)
	GetWeights() (_swig_ret *float32)
}

type SwigcptrFlow_t uintptr

func (p SwigcptrFlow_t) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFlow_t) SwigIsFlow_t() {
}

func (arg1 SwigcptrFlow_t) SetFrom(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_flow_t_from_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrFlow_t) GetFrom() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_flow_t_from_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrFlow_t) SetTo(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_flow_t_to_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrFlow_t) GetTo() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_flow_t_to_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrFlow_t) SetAmount(arg2 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_flow_t_amount_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.float(_swig_i_1))
}

func (arg1 SwigcptrFlow_t) GetAmount() (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	swig_r = (float32)(C._wrap_flow_t_amount_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewFlow_t() (_swig_ret Flow_t) {
	var swig_r Flow_t
	swig_r = (Flow_t)(SwigcptrFlow_t(C._wrap_new_flow_t_emd_f9fc16b73f1936b0()))
	return swig_r
}

func DeleteFlow_t(arg1 Flow_t) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_flow_t_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0))
}

type Flow_t interface {
	Swigcptr() uintptr
	SwigIsFlow_t()
	SetFrom(arg2 int)
	GetFrom() (_swig_ret int)
	SetTo(arg2 int)
	GetTo() (_swig_ret int)
	SetAmount(arg2 float32)
	GetAmount() (_swig_ret float32)
}

type SwigcptrDist_features_t uintptr

func (p SwigcptrDist_features_t) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDist_features_t) SwigIsDist_features_t() {
}

func (arg1 SwigcptrDist_features_t) SetDim(arg2 uint) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_dist_features_t_dim_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrDist_features_t) GetDim() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_dist_features_t_dim_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDist_features_t) SetDistanceMatrix(arg2 *float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_dist_features_t_distanceMatrix_set_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (arg1 SwigcptrDist_features_t) GetDistanceMatrix() (_swig_ret *float32) {
	var swig_r *float32
	_swig_i_0 := arg1
	swig_r = (*float32)(C._wrap_dist_features_t_distanceMatrix_get_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewDist_features_t() (_swig_ret Dist_features_t) {
	var swig_r Dist_features_t
	swig_r = (Dist_features_t)(SwigcptrDist_features_t(C._wrap_new_dist_features_t_emd_f9fc16b73f1936b0()))
	return swig_r
}

func DeleteDist_features_t(arg1 Dist_features_t) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_dist_features_t_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0))
}

type Dist_features_t interface {
	Swigcptr() uintptr
	SwigIsDist_features_t()
	SetDim(arg2 uint)
	GetDim() (_swig_ret uint)
	SetDistanceMatrix(arg2 *float32)
	GetDistanceMatrix() (_swig_ret *float32)
}

func Emd(arg1 Signature_t, arg2 Signature_t, arg3 Dist_features_t, arg4 Flow_t, arg5 *int) (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3.Swigcptr()
	_swig_i_3 := arg4.Swigcptr()
	_swig_i_4 := arg5
	swig_r = (float32)(C._wrap_emd_emd_f9fc16b73f1936b0(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3), C.swig_voidp(_swig_i_4)))
	return swig_r
}