// MIT
package main

/*
#cgo CFLAGS: -I/usr/local/opt/openblas/include
#cgo LDFLAGS: -L/usr/local/opt/openblas/lib -lopenblas
#include "cblas.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func SetNumThreads(numThreads int32) {
	C.openblas_set_num_threads((C.int)(numThreads))

}

func GotoSetNumThreads(numThreads int32) {
	C.goto_set_num_threads((C.int)(numThreads))
}

func GetNumThreads() int32 {
	return int32(C.openblas_get_num_threads())
}

func GetNumProcs() int32 {
	return int32(C.openblas_get_num_procs())
}

func GetConfig() string {
	return C.GoString(C.openblas_get_config())
}

func GetCorename() string {
	return C.GoString(C.openblas_get_corename())
}

func GetParallel() int32 {
	return int32(C.openblas_get_parallel())
}

// Dsdot returns the dot product of two single-precision vectors.
func Dsdot(n int, x []float32, incX int, y []float32, incY int) float64 {
	return float64(C.cblas_dsdot(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX), (*C.float)(unsafe.Pointer(&y[0])), C.int(incY)))
}

// Sdsdot returns the dot product of two single-precision vectors with a correction term.
func Sdsdot(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32 {
	return float32(C.cblas_sdsdot(C.int(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX), (*C.float)(unsafe.Pointer(&y[0])), C.int(incY)))
}

// Ddot returns the dot product of two double-precision vectors.
func Ddot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return float64(C.cblas_ddot(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX), (*C.double)(unsafe.Pointer(&y[0])), C.int(incY)))
}

// Sdot returns the dot product of two single-precision vectors.
func Sdot(n int, x []float32, incX int, y []float32, incY int) float32 {
	return float32(C.cblas_sdot(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX), (*C.float)(unsafe.Pointer(&y[0])), C.int(incY)))
}

// Cdotu returns the dot product of two complex single-precision vectors.
func Cdotu(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	return complex64(C.cblas_cdotu(C.int(n), (unsafe.Pointer(&x[0])), C.int(incX), (unsafe.Pointer(&y[0])), C.int(incY)))
}

// Cdotc returns the dot product of two complex single-precision vectors conjugating the first.
func Cdotc(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	return complex64(C.cblas_cdotc(C.int(n), (unsafe.Pointer(&x[0])), C.int(incX), (unsafe.Pointer(&y[0])), C.int(incY)))
}

// Zdotu returns the dot product of two complex double-precision vectors.
func Zdotu(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	return complex128(C.cblas_zdotu(C.int(n), (unsafe.Pointer(&x[0])), C.int(incX), (unsafe.Pointer(&y[0])), C.int(incY)))
}

// Zdotc returns the dot product of two complex double-precision vectors conjugating the first.
func Zdotc(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	return complex128(C.cblas_zdotc(C.int(n), (unsafe.Pointer(&x[0])), C.int(incX), (unsafe.Pointer(&y[0])), C.int(incY)))
}

// CdotuSub returns the dot product of two complex single-precision vectors and stores the result in ret.
func CdotuSub(n int, x []complex64, incX int, y []complex64, incY int, ret *complex64) {
	C.cblas_cdotu_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(ret))
}

// CdotcSub returns the dot product of two complex single-precision vectors conjugating the first and stores the result in ret.
func CdotcSub(n int, x []complex64, incX int, y []complex64, incY int, ret *complex64) {
	C.cblas_cdotc_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(ret))
}

// ZdotuSub returns the dot product of two complex double-precision vectors and stores the result in ret.
func ZdotuSub(n int, x []complex128, incX int, y []complex128, incY int, ret *complex128) {
	C.cblas_zdotu_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(ret))
}

// ZdotcSub returns the dot product of two complex double-precision vectors conjugating the first and stores the result in ret.
func ZdotcSub(n int, x []complex128, incX int, y []complex128, incY int, ret *complex128) {
	C.cblas_zdotc_sub(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(ret))
}

// Sasum returns the sum of the absolute values of a single-precision vector.
func Sasum(n int, x []float32, incX int) float32 {
	return float32(C.cblas_sasum(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX)))
}

// Dasum returns the sum of the absolute values of a double-precision vector.
func Dasum(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dasum(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX)))
}

// Scasum returns the sum of the absolute values of the real and imaginary components of a complex single-precision vector.
func Scasum(n int, x []complex64, incX int) float32 {
	return float32(C.cblas_scasum(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}

// Dzasum returns the sum of the absolute values of the real and imaginary components of a complex double-precision vector.
func Dzasum(n int, x []complex128, incX int) float64 {
	return float64(C.cblas_dzasum(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}

// Ssum returns the sum of a single-precision vector.
func Ssum(n int, x []float32, incX int) float32 {
	return float32(C.cblas_ssum(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX)))
}

// Dsum returns the sum of a double-precision vector.
func Dsum(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dsum(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX)))
}

// Scsum returns the sum of the real and imaginary components of a complex single-precision vector.
func Scsum(n int, x []complex64, incX int) float32 {
	return float32(C.cblas_scsum(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}

// Dzsum returns the sum of the real and imaginary components of a complex double-precision vector.
func Dzsum(n int, x []complex128, incX int) float64 {
	return float64(C.cblas_dzsum(C.int(n), unsafe.Pointer(&x[0]), C.int(incX)))
}

// Snrm2 returns the Euclidean norm (2-norm) of a single-precision vector.
func Snrm2(N int, X []float32, incX int) float32 {
	return float32(C.cblas_snrm2(C.int(N), (*C.float)(unsafe.Pointer(&X[0])), C.int(incX)))
}

// Dnrm2 returns the Euclidean norm (2-norm) of a double-precision vector.
func Dnrm2(N int, X []float64, incX int) float64 {
	return float64(C.cblas_dnrm2(C.int(N), (*C.double)(unsafe.Pointer(&X[0])), C.int(incX)))
}

// Scnrm2 returns the Euclidean norm (2-norm) of a complex single-precision vector.
func Scnrm2(N int, X []complex64, incX int) float32 {
	return float32(C.cblas_scnrm2(C.int(N), unsafe.Pointer(&X[0]), C.int(incX)))
}

// Dznrm2 returns the Euclidean norm (2-norm) of a complex double-precision vector.
func Dznrm2(N int, X []complex128, incX int) float64 {
	return float64(C.cblas_dznrm2(C.int(N), unsafe.Pointer(&X[0]), C.int(incX)))
}

// Isamax returns the index of the element with the largest absolute value in a single-precision vector.
func Isamax(n int, x []float32, incx int) int {
	return int(C.cblas_isamax(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Idamax returns the index of the element with the largest absolute value in a double-precision vector.
func Idamax(n int, x []float64, incx int) int {
	return int(C.cblas_idamax(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Icamax returns the index of the element with the largest absolute value in a complex single-precision vector.
func Icamax(n int, x []complex64, incx int) int {
	return int(C.cblas_icamax(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Izamax returns the index of the element with the largest absolute value in a complex double-precision vector.
func Izamax(n int, x []complex128, incx int) int {
	return int(C.cblas_izamax(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Isamin returns the index of the element with the smallest absolute value in a single-precision vector.
func Isamin(n int, x []float32, incx int) int {
	return int(C.cblas_isamin(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Idamin returns the index of the element with the smallest absolute value in a double-precision vector.
func Idamin(n int, x []float64, incx int) int {
	return int(C.cblas_idamin(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Icamin returns the index of the element with the smallest absolute value in a complex single-precision vector.
func Icamin(n int, x []complex64, incx int) int {
	return int(C.cblas_icamin(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Izamin returns the index of the element with the smallest absolute value in a complex double-precision vector.
func Izamin(n int, x []complex128, incx int) int {
	return int(C.cblas_izamin(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Ismax returns the index of the element with the largest absolute value in a single-precision vector.
func Ismax(n int, x []float32, incx int) int {
	return int(C.cblas_ismax(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Idmax returns the index of the element with the largest absolute value in a double-precision vector.
func Idmax(n int, x []float64, incx int) int {
	return int(C.cblas_idmax(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Icmax returns the index of the element with the largest absolute value in a complex single-precision vector.
func Icmax(n int, x []complex64, incx int) int {
	return int(C.cblas_icmax(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Izmax returns the index of the element with the largest absolute value in a complex double-precision vector.
func Izmax(n int, x []complex128, incx int) int {
	return int(C.cblas_izmax(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Ismin returns the index of the element with the smallest value in a single-precision vector.
func Ismin(n int, x []float32, incx int) int {
	return int(C.cblas_ismin(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Idmin returns the index of the element with the smallest value in a double-precision vector.
func Idmin(n int, x []float64, incx int) int {
	return int(C.cblas_idmin(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx)))
}

// Icmin returns the index of the element with the smallest value in a complex single-precision vector.
func Icmin(n int, x []complex64, incx int) int {
	return int(C.cblas_icmin(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Izmin returns the index of the element with the smallest value in a complex double-precision vector.
func Izmin(n int, x []complex128, incx int) int {
	return int(C.cblas_izmin(C.int(n), unsafe.Pointer(&x[0]), C.int(incx)))
}

// Saxpy adds a multiple of a single-precision vector to another single-precision vector.
func Saxpy(n int, alpha float32, x []float32, incx int, y []float32, incy int) {
	C.cblas_saxpy(C.int(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx), (*C.float)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Daxpy adds a multiple of a double-precision vector to another double-precision vector.
func Daxpy(n int, alpha float64, x []float64, incx int, y []float64, incy int) {
	C.cblas_daxpy(C.int(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx), (*C.double)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Caxpy adds a multiple of a complex single-precision vector to another complex single-precision vector.
func Caxpy(n int, alpha complex64, x []complex64, incx int, y []complex64, incy int) {
	C.cblas_caxpy(C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&y[0]), C.int(incy))
}

// Zaxpy adds a multiple of a complex double-precision vector to another complex double-precision vector.
func Zaxpy(n int, alpha complex128, x []complex128, incx int, y []complex128, incy int) {
	C.cblas_zaxpy(C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&y[0]), C.int(incy))
}

// Scopy copies a single-precision vector x to a single-precision vector y.
func Scopy(n int, x []float32, incx int, y []float32, incy int) {
	C.cblas_scopy(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx), (*C.float)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Dcopy copies a double-precision vector x to a double-precision vector y.
func Dcopy(n int, x []float64, incx int, y []float64, incy int) {
	C.cblas_dcopy(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx), (*C.double)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Ccopy copies a complex single-precision vector x to a complex single-precision vector y.
func Ccopy(n int, x []complex64, incx int, y []complex64, incy int) {
	C.cblas_ccopy(C.int(n), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&y[0]), C.int(incy))
}

// Zcopy copies a complex double-precision vector x to a complex double-precision vector y.
func Zcopy(n int, x []complex128, incx int, y []complex128, incy int) {
	C.cblas_zcopy(C.int(n), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&y[0]), C.int(incy))
}

// Sswap interchanges two single-precision vectors.
func Sswap(n int, x []float32, incx int, y []float32, incy int) {
	C.cblas_sswap(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx), (*C.float)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Dswap interchanges two double-precision vectors.
func Dswap(n int, x []float64, incx int, y []float64, incy int) {
	C.cblas_dswap(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx), (*C.double)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Cswap interchanges two complex single-precision vectors.
func Cswap(n int, x []complex64, incx int, y []complex64, incy int) {
	C.cblas_cswap(C.int(n), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&y[0]), C.int(incy))
}

// Zswap interchanges two complex double-precision vectors.
func Zswap(n int, x []complex128, incx int, y []complex128, incy int) {
	C.cblas_zswap(C.int(n), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&y[0]), C.int(incy))
}

// Srot applies a plane rotation to vectors x and y.
func Srot(n int, x []float32, incX int, y []float32, incY int, c float32, s float32) {
	C.cblas_srot(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX), (*C.float)(unsafe.Pointer(&y[0])), C.int(incY), C.float(c), C.float(s))
}

// Drot applies a plane rotation to vectors x and y.
func Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	C.cblas_drot(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX), (*C.double)(unsafe.Pointer(&y[0])), C.int(incY), C.double(c), C.double(s))
}

// Csrot applies a plane rotation to complex single-precision vectors x and y.
func Csrot(n int, x []complex64, incX int, y []complex64, incY int, c float32, s float32) {
	//C.cblas_csrot(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), C.float(c), C.float(s))
}

// Zdrot applies a plane rotation to complex double-precision vectors x and y.
func Zdrot(n int, x []complex128, incX int, y []complex128, incY int, c float64, s float64) {
	//C.cblas_zdrot(C.int(n), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), C.double(c), C.double(s))
}

// cblas_srotg computes the parameters for a Givens rotation matrix.
func Srotg(a *float32, b *float32, c *float32, s *float32) {
	C.cblas_srotg((*C.float)(unsafe.Pointer(a)), (*C.float)(unsafe.Pointer(b)), (*C.float)(unsafe.Pointer(c)), (*C.float)(unsafe.Pointer(s)))
}

// cblas_drotg computes the parameters for a Givens rotation matrix.
func Drotg(a *float64, b *float64, c *float64, s *float64) {
	C.cblas_drotg((*C.double)(unsafe.Pointer(a)), (*C.double)(unsafe.Pointer(b)), (*C.double)(unsafe.Pointer(c)), (*C.double)(unsafe.Pointer(s)))
}

// cblas_crotg computes the parameters for a Givens rotation matrix.
func Crotg(a *complex64, b *complex64, c *float32, s *complex64) {
	//C.cblas_crotg(unsafe.Pointer(a), unsafe.Pointer(b), (*C.float)(unsafe.Pointer(c)), unsafe.Pointer(s))
}

// cblas_zrotg computes the parameters for a Givens rotation matrix.
func Zrotg(a *complex128, b *complex128, c *float64, s *complex128) {
	//C.cblas_zrotg(unsafe.Pointer(a), unsafe.Pointer(b), (*C.double)(unsafe.Pointer(c)), unsafe.Pointer(s))
}

// Rotates the points (X,Y) in the plane through the angle THETA.
func Srotm(n int32, x []float32, incX int32, y []float32, incY int32, p []float32) {
	C.cblas_srotm(C.int(n), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX), (*C.float)(unsafe.Pointer(&y[0])), C.int(incY), (*C.float)(unsafe.Pointer(&p[0])))
}

// Rotates the points (X,Y) in the plane through the angle THETA.
func Drotm(n int32, x []float64, incX int32, y []float64, incY int32, p []float64) {
	C.cblas_drotm(C.int(n), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX), (*C.double)(unsafe.Pointer(&y[0])), C.int(incY), (*C.double)(unsafe.Pointer(&p[0])))
}

// Given the scalars d1, d2, and b2, constructs the modified Givens transformation matrix H which zeros the second component of the 2-vector transpose(d1,d2)
func Srotmg(d1, d2, b1 *float32, b2 float32, p []float32) {
	C.cblas_srotmg((*C.float)(d1), (*C.float)(d2), (*C.float)(b1), C.float(b2), (*C.float)(unsafe.Pointer(&p[0])))
}

// Given the scalars d1, d2, and b2, constructs the modified Givens transformation matrix H which zeros the second component of the 2-vector transpose(d1,d2)
func Drotmg(d1, d2, b1 *float64, b2 float64, p []float64) {
	C.cblas_drotmg((*C.double)(d1), (*C.double)(d2), (*C.double)(b1), C.double(b2), (*C.double)(unsafe.Pointer(&p[0])))
}

// Scalar multiplication of a single-precision float vector X by a scalar alpha.
func Sscal(n int, alpha float32, x []float32, incX int) {
	C.cblas_sscal(C.int(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX))
}

// Scalar multiplication of a double-precision float vector X by a scalar alpha.
func Dscal(n int, alpha float64, x []float64, incX int) {
	C.cblas_dscal(C.int(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX))
}

// Scalar multiplication of a complex single-precision float vector X by a complex scalar alpha.
func Cscal(n int, alpha complex64, x []complex64, incX int) {
	C.cblas_cscal(C.int(n), (unsafe.Pointer(&alpha)), unsafe.Pointer(&x[0]), C.int(incX))
}

// Scalar multiplication of a complex double-precision float vector X by a complex scalar alpha.
func Zscal(n int, alpha complex128, x []complex128, incX int) {
	C.cblas_zscal(C.int(n), (unsafe.Pointer(&alpha)), unsafe.Pointer(&x[0]), C.int(incX))
}

// Scalar multiplication of a complex single-precision float vector X by a real scalar alpha.
func Csscal(n int, alpha float32, x []complex64, incX int) {
	C.cblas_csscal(C.int(n), C.float(alpha), unsafe.Pointer(&x[0]), C.int(incX))
}

// Scalar multiplication of a complex double-precision float vector X by a real scalar alpha.
func Zdscal(n int, alpha float64, x []complex128, incX int) {
	C.cblas_zdscal(C.int(n), C.double(alpha), unsafe.Pointer(&x[0]), C.int(incX))
}

// Sgemv performs a matrix-vector operation with a general rectangular matrix, for real single precision elements.
func Sgemv(order Order, trans Transpose, m, n int, alpha float32, a []float32, lda int, x []float32, incx int, beta float32, y []float32, incy int) {
	C.cblas_sgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.int(m), C.int(n), C.float(alpha), (*C.float)(unsafe.Pointer(&a[0])), C.int(lda), (*C.float)(unsafe.Pointer(&x[0])), C.int(incx), C.float(beta), (*C.float)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Dgemv performs a matrix-vector operation with a general rectangular matrix, for real double precision elements.
func Dgemv(order Order, trans Transpose, m, n int, alpha float64, a []float64, lda int, x []float64, incx int, beta float64, y []float64, incy int) {
	C.cblas_dgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.int(m), C.int(n), C.double(alpha), (*C.double)(unsafe.Pointer(&a[0])), C.int(lda), (*C.double)(unsafe.Pointer(&x[0])), C.int(incx), C.double(beta), (*C.double)(unsafe.Pointer(&y[0])), C.int(incy))
}

// Cgemv performs a matrix-vector operation with a general rectangular matrix, for complex single precision elements.
func Cgemv(order Order, trans Transpose, m, n int, alpha complex64, a []complex64, lda int, x []complex64, incx int, beta complex64, y []complex64, incy int) {
	C.cblas_cgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incy))
}

// Zgemv performs a matrix-vector operation with a general rectangular matrix, for complex double precision elements.
func Zgemv(order Order, trans Transpose, m, n int, alpha complex128, a []complex128, lda int, x []complex128, incx int, beta complex128, y []complex128, incy int) {
	C.cblas_zgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incx), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.int(incy))
}

// Sger performs a rank-1 update of matrix A. A = alpha * X * Y^T + A
func Sger(order Order, m, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	C.cblas_sger(C.enum_CBLAS_ORDER(order), C.int(m), C.int(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX), (*C.float)(unsafe.Pointer(&y[0])), C.int(incY), (*C.float)(unsafe.Pointer(&a[0])), C.int(lda))
}

// Dger performs a rank-1 update of matrix A. A = alpha * X * Y^T + A
func Dger(order Order, m, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	C.cblas_dger(C.enum_CBLAS_ORDER(order), C.int(m), C.int(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX), (*C.double)(unsafe.Pointer(&y[0])), C.int(incY), (*C.double)(unsafe.Pointer(&a[0])), C.int(lda))
}

// Cgeru performs a rank-1 update of matrix A. A = alpha * X * Y^H + A
func Cgeru(order Order, m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cgeru(C.enum_CBLAS_ORDER(order), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}

// Cgerc performs a rank-1 update of matrix A. A = alpha * X * conj(Y)^T + A
func Cgerc(order Order, m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cgerc(C.enum_CBLAS_ORDER(order), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}

// Zgeru performs a rank-1 update of matrix A. A = alpha * X * Y^H + A
func Zgeru(order Order, m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zgeru(C.enum_CBLAS_ORDER(order), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}

// Zgerc performs a rank-1 update of matrix A. A = alpha * X * conj(Y)^T + A
func Zgerc(order Order, m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zgerc(C.enum_CBLAS_ORDER(order), C.int(m), C.int(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&y[0]), C.int(incY), unsafe.Pointer(&a[0]), C.int(lda))
}

// Strsv solves a system of linear equations with a triangular matrix stored in packed format.
func Strsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float32, lda int, x []float32, incX int) {
	C.cblas_strsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.int(n), (*C.float)(unsafe.Pointer(&a[0])), C.int(lda), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX))
}

// Dtrsv solves a system of linear equations with a triangular matrix stored in packed format.
func Dtrsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.int(n), (*C.double)(unsafe.Pointer(&a[0])), C.int(lda), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX))
}

// Ctrsv solves a system of linear equations with a triangular matrix stored in packed format.
func Ctrsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctrsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}

// Ztrsv solves a system of linear equations with a triangular matrix stored in packed format.
func Ztrsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztrsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}

// Strmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Strmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float32, lda int, x []float32, incX int) {
	C.cblas_strmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.int(n), (*C.float)(unsafe.Pointer(&a[0])), C.int(lda), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX))
}

// Dtrmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Dtrmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.int(n), (*C.double)(unsafe.Pointer(&a[0])), C.int(lda), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX))
}

// Ctrmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Ctrmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctrmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}

// Ztrmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Ztrmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztrmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&x[0]), C.int(incX))
}

// Ssyr performs a symmetric rank-1 update of a real symmetric matrix.
func Ssyr(order Order, upLo UpLo, n int, alpha float32, x []float32, incX int, a []float32, lda int) {
	C.cblas_ssyr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.int(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.int(incX), (*C.float)(unsafe.Pointer(&a[0])), C.int(lda))
}

// Dsyr performs a symmetric rank-1 update of a real symmetric matrix.
func Dsyr(order Order, upLo UpLo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	C.cblas_dsyr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.int(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.int(incX), (*C.double)(unsafe.Pointer(&a[0])), C.int(lda))
}

// Cher performs a hermitian rank-1 update of a complex hermitian matrix.
func Cher(order Order, upLo UpLo, n int, alpha float32, x []complex64, incX int, a []complex64, lda int) {
	C.cblas_cher(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.int(n), C.float(alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&a[0]), C.int(lda))
}

// Zher performs a hermitian rank-1 update of a complex hermitian matrix.
func Zher(order Order, upLo UpLo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int) {
	C.cblas_zher(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.int(n), C.double(alpha), unsafe.Pointer(&x[0]), C.int(incX), unsafe.Pointer(&a[0]), C.int(lda))
}

func main() {
	list := []float64{1.1, 10, 1, 2, 3, 4, 5}
	fmt.Println(Idmin(len(list), list, 1))
}
