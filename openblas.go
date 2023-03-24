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
	return float64(C.cblas_dsdot(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// Sdsdot returns the dot product of two single-precision vectors with a correction term.
func Sdsdot(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32 {
	return float32(C.cblas_sdsdot(C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// Ddot returns the dot product of two double-precision vectors.
func Ddot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return float64(C.cblas_ddot(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// Sdot returns the dot product of two single-precision vectors.
func Sdot(n int, x []float32, incX int, y []float32, incY int) float32 {
	return float32(C.cblas_sdot(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// Cdotu returns the dot product of two complex single-precision vectors.
func Cdotu(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	return complex64(C.cblas_cdotu(C.blasint(n), (unsafe.Pointer(&x[0])), C.blasint(incX), (unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// Cdotc returns the dot product of two complex single-precision vectors conjugating the first.
func Cdotc(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	return complex64(C.cblas_cdotc(C.blasint(n), (unsafe.Pointer(&x[0])), C.blasint(incX), (unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// Zdotu returns the dot product of two complex double-precision vectors.
func Zdotu(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	return complex128(C.cblas_zdotu(C.blasint(n), (unsafe.Pointer(&x[0])), C.blasint(incX), (unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// Zdotc returns the dot product of two complex double-precision vectors conjugating the first.
func Zdotc(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	return complex128(C.cblas_zdotc(C.blasint(n), (unsafe.Pointer(&x[0])), C.blasint(incX), (unsafe.Pointer(&y[0])), C.blasint(incY)))
}

// CdotuSub returns the dot product of two complex single-precision vectors and stores the result in ret.
func CdotuSub(n int, x []complex64, incX int, y []complex64, incY int, ret *complex64) {
	C.cblas_cdotu_sub(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(ret))
}

// CdotcSub returns the dot product of two complex single-precision vectors conjugating the first and stores the result in ret.
func CdotcSub(n int, x []complex64, incX int, y []complex64, incY int, ret *complex64) {
	C.cblas_cdotc_sub(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(ret))
}

// ZdotuSub returns the dot product of two complex double-precision vectors and stores the result in ret.
func ZdotuSub(n int, x []complex128, incX int, y []complex128, incY int, ret *complex128) {
	C.cblas_zdotu_sub(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(ret))
}

// ZdotcSub returns the dot product of two complex double-precision vectors conjugating the first and stores the result in ret.
func ZdotcSub(n int, x []complex128, incX int, y []complex128, incY int, ret *complex128) {
	C.cblas_zdotc_sub(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(ret))
}

// Sasum returns the sum of the absolute values of a single-precision vector.
func Sasum(n int, x []float32, incX int) float32 {
	return float32(C.cblas_sasum(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX)))
}

// Dasum returns the sum of the absolute values of a double-precision vector.
func Dasum(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dasum(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX)))
}

// Scasum returns the sum of the absolute values of the real and imaginary components of a complex single-precision vector.
func Scasum(n int, x []complex64, incX int) float32 {
	return float32(C.cblas_scasum(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX)))
}

// Dzasum returns the sum of the absolute values of the real and imaginary components of a complex double-precision vector.
func Dzasum(n int, x []complex128, incX int) float64 {
	return float64(C.cblas_dzasum(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX)))
}

// Ssum returns the sum of a single-precision vector.
func Ssum(n int, x []float32, incX int) float32 {
	return float32(C.cblas_ssum(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX)))
}

// Dsum returns the sum of a double-precision vector.
func Dsum(n int, x []float64, incX int) float64 {
	return float64(C.cblas_dsum(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX)))
}

// Scsum returns the sum of the real and imaginary components of a complex single-precision vector.
func Scsum(n int, x []complex64, incX int) float32 {
	return float32(C.cblas_scsum(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX)))
}

// Dzsum returns the sum of the real and imaginary components of a complex double-precision vector.
func Dzsum(n int, x []complex128, incX int) float64 {
	return float64(C.cblas_dzsum(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX)))
}

// Snrm2 returns the Euclidean norm (2-norm) of a single-precision vector.
func Snrm2(N int, X []float32, incX int) float32 {
	return float32(C.cblas_snrm2(C.blasint(N), (*C.float)(unsafe.Pointer(&X[0])), C.blasint(incX)))
}

// Dnrm2 returns the Euclidean norm (2-norm) of a double-precision vector.
func Dnrm2(N int, X []float64, incX int) float64 {
	return float64(C.cblas_dnrm2(C.blasint(N), (*C.double)(unsafe.Pointer(&X[0])), C.blasint(incX)))
}

// Scnrm2 returns the Euclidean norm (2-norm) of a complex single-precision vector.
func Scnrm2(N int, X []complex64, incX int) float32 {
	return float32(C.cblas_scnrm2(C.blasint(N), unsafe.Pointer(&X[0]), C.blasint(incX)))
}

// Dznrm2 returns the Euclidean norm (2-norm) of a complex double-precision vector.
func Dznrm2(N int, X []complex128, incX int) float64 {
	return float64(C.cblas_dznrm2(C.blasint(N), unsafe.Pointer(&X[0]), C.blasint(incX)))
}

// Isamax returns the index of the element with the largest absolute value in a single-precision vector.
func Isamax(n int, x []float32, incx int) int {
	return int(C.cblas_isamax(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Idamax returns the index of the element with the largest absolute value in a double-precision vector.
func Idamax(n int, x []float64, incx int) int {
	return int(C.cblas_idamax(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Icamax returns the index of the element with the largest absolute value in a complex single-precision vector.
func Icamax(n int, x []complex64, incx int) int {
	return int(C.cblas_icamax(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Izamax returns the index of the element with the largest absolute value in a complex double-precision vector.
func Izamax(n int, x []complex128, incx int) int {
	return int(C.cblas_izamax(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Isamin returns the index of the element with the smallest absolute value in a single-precision vector.
func Isamin(n int, x []float32, incx int) int {
	return int(C.cblas_isamin(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Idamin returns the index of the element with the smallest absolute value in a double-precision vector.
func Idamin(n int, x []float64, incx int) int {
	return int(C.cblas_idamin(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Icamin returns the index of the element with the smallest absolute value in a complex single-precision vector.
func Icamin(n int, x []complex64, incx int) int {
	return int(C.cblas_icamin(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Izamin returns the index of the element with the smallest absolute value in a complex double-precision vector.
func Izamin(n int, x []complex128, incx int) int {
	return int(C.cblas_izamin(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Ismax returns the index of the element with the largest absolute value in a single-precision vector.
func Ismax(n int, x []float32, incx int) int {
	return int(C.cblas_ismax(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Idmax returns the index of the element with the largest absolute value in a double-precision vector.
func Idmax(n int, x []float64, incx int) int {
	return int(C.cblas_idmax(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Icmax returns the index of the element with the largest absolute value in a complex single-precision vector.
func Icmax(n int, x []complex64, incx int) int {
	return int(C.cblas_icmax(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Izmax returns the index of the element with the largest absolute value in a complex double-precision vector.
func Izmax(n int, x []complex128, incx int) int {
	return int(C.cblas_izmax(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Ismin returns the index of the element with the smallest value in a single-precision vector.
func Ismin(n int, x []float32, incx int) int {
	return int(C.cblas_ismin(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Idmin returns the index of the element with the smallest value in a double-precision vector.
func Idmin(n int, x []float64, incx int) int {
	return int(C.cblas_idmin(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx)))
}

// Icmin returns the index of the element with the smallest value in a complex single-precision vector.
func Icmin(n int, x []complex64, incx int) int {
	return int(C.cblas_icmin(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Izmin returns the index of the element with the smallest value in a complex double-precision vector.
func Izmin(n int, x []complex128, incx int) int {
	return int(C.cblas_izmin(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx)))
}

// Saxpy adds a multiple of a single-precision vector to another single-precision vector.
func Saxpy(n int, alpha float32, x []float32, incx int, y []float32, incy int) {
	C.cblas_saxpy(C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Daxpy adds a multiple of a double-precision vector to another double-precision vector.
func Daxpy(n int, alpha float64, x []float64, incx int, y []float64, incy int) {
	C.cblas_daxpy(C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Caxpy adds a multiple of a complex single-precision vector to another complex single-precision vector.
func Caxpy(n int, alpha complex64, x []complex64, incx int, y []complex64, incy int) {
	C.cblas_caxpy(C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Zaxpy adds a multiple of a complex double-precision vector to another complex double-precision vector.
func Zaxpy(n int, alpha complex128, x []complex128, incx int, y []complex128, incy int) {
	C.cblas_zaxpy(C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Scopy copies a single-precision vector x to a single-precision vector y.
func Scopy(n int, x []float32, incx int, y []float32, incy int) {
	C.cblas_scopy(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Dcopy copies a double-precision vector x to a double-precision vector y.
func Dcopy(n int, x []float64, incx int, y []float64, incy int) {
	C.cblas_dcopy(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Ccopy copies a complex single-precision vector x to a complex single-precision vector y.
func Ccopy(n int, x []complex64, incx int, y []complex64, incy int) {
	C.cblas_ccopy(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Zcopy copies a complex double-precision vector x to a complex double-precision vector y.
func Zcopy(n int, x []complex128, incx int, y []complex128, incy int) {
	C.cblas_zcopy(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Sswap interchanges two single-precision vectors.
func Sswap(n int, x []float32, incx int, y []float32, incy int) {
	C.cblas_sswap(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Dswap interchanges two double-precision vectors.
func Dswap(n int, x []float64, incx int, y []float64, incy int) {
	C.cblas_dswap(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Cswap interchanges two complex single-precision vectors.
func Cswap(n int, x []complex64, incx int, y []complex64, incy int) {
	C.cblas_cswap(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Zswap interchanges two complex double-precision vectors.
func Zswap(n int, x []complex128, incx int, y []complex128, incy int) {
	C.cblas_zswap(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Srot applies a plane rotation to vectors x and y.
func Srot(n int, x []float32, incX int, y []float32, incY int, c float32, s float32) {
	C.cblas_srot(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY), C.float(c), C.float(s))
}

// Drot applies a plane rotation to vectors x and y.
func Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	C.cblas_drot(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY), C.double(c), C.double(s))
}

// Csrot applies a plane rotation to complex single-precision vectors x and y.
func Csrot(n int, x []complex64, incX int, y []complex64, incY int, c float32, s float32) {
	//C.cblas_csrot(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), C.float(c), C.float(s))
}

// Zdrot applies a plane rotation to complex double-precision vectors x and y.
func Zdrot(n int, x []complex128, incX int, y []complex128, incY int, c float64, s float64) {
	//C.cblas_zdrot(C.blasint(n), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), C.double(c), C.double(s))
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
	C.cblas_srotm(C.blasint(n), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.float)(unsafe.Pointer(&p[0])))
}

// Rotates the points (X,Y) in the plane through the angle THETA.
func Drotm(n int32, x []float64, incX int32, y []float64, incY int32, p []float64) {
	C.cblas_drotm(C.blasint(n), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.double)(unsafe.Pointer(&p[0])))
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
	C.cblas_sscal(C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Scalar multiplication of a double-precision float vector X by a scalar alpha.
func Dscal(n int, alpha float64, x []float64, incX int) {
	C.cblas_dscal(C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Scalar multiplication of a complex single-precision float vector X by a complex scalar alpha.
func Cscal(n int, alpha complex64, x []complex64, incX int) {
	C.cblas_cscal(C.blasint(n), (unsafe.Pointer(&alpha)), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Scalar multiplication of a complex double-precision float vector X by a complex scalar alpha.
func Zscal(n int, alpha complex128, x []complex128, incX int) {
	C.cblas_zscal(C.blasint(n), (unsafe.Pointer(&alpha)), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Scalar multiplication of a complex single-precision float vector X by a real scalar alpha.
func Csscal(n int, alpha float32, x []complex64, incX int) {
	C.cblas_csscal(C.blasint(n), C.float(alpha), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Scalar multiplication of a complex double-precision float vector X by a real scalar alpha.
func Zdscal(n int, alpha float64, x []complex128, incX int) {
	C.cblas_zdscal(C.blasint(n), C.double(alpha), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Sgemv performs a matrix-vector operation with a general rectangular matrix, for real single precision elements.
func Sgemv(order Order, trans Transpose, m, n int, alpha float32, a []float32, lda int, x []float32, incx int, beta float32, y []float32, incy int) {
	C.cblas_sgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.blasint(m), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incx), C.float(beta), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Dgemv performs a matrix-vector operation with a general rectangular matrix, for real double precision elements.
func Dgemv(order Order, trans Transpose, m, n int, alpha float64, a []float64, lda int, x []float64, incx int, beta float64, y []float64, incy int) {
	C.cblas_dgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.blasint(m), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incx), C.double(beta), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incy))
}

// Cgemv performs a matrix-vector operation with a general rectangular matrix, for complex single precision elements.
func Cgemv(order Order, trans Transpose, m, n int, alpha complex64, a []complex64, lda int, x []complex64, incx int, beta complex64, y []complex64, incy int) {
	C.cblas_cgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.blasint(m), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Zgemv performs a matrix-vector operation with a general rectangular matrix, for complex double precision elements.
func Zgemv(order Order, trans Transpose, m, n int, alpha complex128, a []complex128, lda int, x []complex128, incx int, beta complex128, y []complex128, incy int) {
	C.cblas_zgemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(trans), C.blasint(m), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incx), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incy))
}

// Sger performs a rank-1 update of matrix A. A = alpha * X * Y^T + A
func Sger(order Order, m, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	C.cblas_sger(C.enum_CBLAS_ORDER(order), C.blasint(m), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda))
}

// Dger performs a rank-1 update of matrix A. A = alpha * X * Y^T + A
func Dger(order Order, m, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	C.cblas_dger(C.enum_CBLAS_ORDER(order), C.blasint(m), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda))
}

// Cgeru performs a rank-1 update of matrix A. A = alpha * X * Y^H + A
func Cgeru(order Order, m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cgeru(C.enum_CBLAS_ORDER(order), C.blasint(m), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Cgerc performs a rank-1 update of matrix A. A = alpha * X * conj(Y)^T + A
func Cgerc(order Order, m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cgerc(C.enum_CBLAS_ORDER(order), C.blasint(m), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Zgeru performs a rank-1 update of matrix A. A = alpha * X * Y^H + A
func Zgeru(order Order, m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zgeru(C.enum_CBLAS_ORDER(order), C.blasint(m), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Zgerc performs a rank-1 update of matrix A. A = alpha * X * conj(Y)^T + A
func Zgerc(order Order, m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zgerc(C.enum_CBLAS_ORDER(order), C.blasint(m), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Strsv solves a system of linear equations with a triangular matrix stored in packed format.
func Strsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float32, lda int, x []float32, incX int) {
	C.cblas_strsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Dtrsv solves a system of linear equations with a triangular matrix stored in packed format.
func Dtrsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Ctrsv solves a system of linear equations with a triangular matrix stored in packed format.
func Ctrsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctrsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ztrsv solves a system of linear equations with a triangular matrix stored in packed format.
func Ztrsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztrsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Strmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Strmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float32, lda int, x []float32, incX int) {
	C.cblas_strmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Dtrmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Dtrmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtrmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Ctrmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Ctrmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctrmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ztrmv performs one of the matrix-vector operations x = A*x or x = Aᵀ*x where A is a triangular matrix.
func Ztrmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztrmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ssyr performs a symmetric rank-1 update of a real symmetric matrix.
func Ssyr(order Order, upLo UpLo, n int, alpha float32, x []float32, incX int, a []float32, lda int) {
	C.cblas_ssyr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda))
}

// Dsyr performs a symmetric rank-1 update of a real symmetric matrix.
func Dsyr(order Order, upLo UpLo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	C.cblas_dsyr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda))
}

// Cher performs a hermitian rank-1 update of a complex hermitian matrix.
func Cher(order Order, upLo UpLo, n int, alpha float32, x []complex64, incX int, a []complex64, lda int) {
	C.cblas_cher(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Zher performs a hermitian rank-1 update of a complex hermitian matrix.
func Zher(order Order, upLo UpLo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int) {
	C.cblas_zher(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Ssyr2 performs the symmetric rank 2 operation A := alpha*x*y^T + alpha*y*x^T + A, where alpha is a scalar, x and y are n element vectors, and A is an n by n symmetric matrix, supplied in packed form.
func Ssyr2(order Order, upLo UpLo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	C.cblas_ssyr2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda))
}

// Dsyr2 performs the symmetric rank 2 operation A := alpha*x*y^T + alpha*y*x^T + A, where alpha is a scalar, x and y are n element vectors, and A is an n by n symmetric matrix, supplied in packed form.
func Dsyr2(order Order, upLo UpLo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	C.cblas_dsyr2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda))
}

// Cher2 performs the Hermitian rank-2 update
func Cher2(order Order, upLo UpLo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	C.cblas_cher2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Zher2 performs the Hermitian rank-2 update
func Zher2(order Order, upLo UpLo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	C.cblas_zher2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&a[0]), C.blasint(lda))
}

// Sgbmv applies a general banded matrix to a vector (single precision).
func Sgbmv(order Order, transA Transpose, m int, n int, kl int, ku int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_sgbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(transA), C.blasint(m), C.blasint(n), C.blasint(kl), C.blasint(ku), C.float(alpha), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), C.float(beta), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Dgbmv applies a general banded matrix to a vector (double precision).
func Dgbmv(order Order, transA Transpose, m int, n int, kl int, ku int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dgbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(transA), C.blasint(m), C.blasint(n), C.blasint(kl), C.blasint(ku), C.double(alpha), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), C.double(beta), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Cgbmv applies a general banded matrix to a vector (single precision).
func Cgbmv(order Order, transA Transpose, m int, n int, kl int, ku int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_cgbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(transA), C.blasint(m), C.blasint(n), C.blasint(kl), C.blasint(ku), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// Zgbmv applies a general banded matrix to a vector (double precision).
func Zgbmv(order Order, transA Transpose, m int, n int, kl int, ku int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zgbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_TRANSPOSE(transA), C.blasint(m), C.blasint(n), C.blasint(kl), C.blasint(ku), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// Ssbmv computes the matrix-vector product using a symmetric band matrix.
func Ssbmv(order Order, upLo UpLo, n, k int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_ssbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.blasint(k), C.float(alpha), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), C.float(beta), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Dsbmv computes the matrix-vector product using a symmetric band matrix.
func Dsbmv(order Order, upLo UpLo, n, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dsbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.blasint(k), C.double(alpha), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), C.double(beta), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Stbmv multiplies a triangular banded matrix with a vector.
func Stbmv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []float32, lda int, x []float32, incX int) {
	C.cblas_stbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), C.blasint(k), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Dtbmv multiplies a triangular banded matrix with a vector.
func Dtbmv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), C.blasint(k), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Ctbmv multiplies a triangular banded matrix with a vector.
func Ctbmv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), C.blasint(k), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ztbmv multiplies a triangular banded matrix with a vector.
func Ztbmv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), C.blasint(k), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Stbsv solves a banded triangular system of equations with a general matrix.
func Stbsv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []float32, lda int, x []float32, incX int) {
	C.cblas_stbsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), C.blasint(k), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Dtbsv solves a banded triangular system of equations with a general matrix.
func DTbsv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []float64, lda int, x []float64, incX int) {
	C.cblas_dtbsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), C.blasint(k), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Ctbsv solves a banded triangular system of equations with a general matrix.
func CTbsv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []complex64, lda int, x []complex64, incX int) {
	C.cblas_ctbsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), C.blasint(k), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ztbsv solves a banded triangular system of equations with a general matrix.
func Ztbsv(order Order, upLo UpLo, trans Transpose, diag Diag, n, k int, a []complex128, lda int, x []complex128, incX int) {
	C.cblas_ztbsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), C.blasint(k), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Stpmv Multiplies a packed triangular matrix by a vector.
func Stpmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []float32, x []float32, incX int) {
	C.cblas_stpmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), (*C.float)(unsafe.Pointer(&ap[0])), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Dtpmv Multiplies a packed triangular matrix by a vector.
func Dtpmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []float64, x []float64, incX int) {
	C.cblas_dtpmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), (*C.double)(unsafe.Pointer(&ap[0])), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Ctpmv Multiplies a packed triangular matrix by a vector.
func Ctpmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []complex64, x []complex64, incX int) {
	C.cblas_ctpmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ztpmv Multiplies a packed triangular matrix by a vector.
func Ztpmv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []complex128, x []complex128, incX int) {
	C.cblas_ztpmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag), C.blasint(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Stpsv Multiplies a packed triangular matrix by a vector
func Stpsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []float32, x []float32, incX int) {
	C.cblas_stpsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), (*C.float)(unsafe.Pointer(&ap[0])), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Dtpsv Multiplies a packed triangular matrix by a vector
func Dtpsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []float64, x []float64, incX int) {
	C.cblas_dtpsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), (*C.double)(unsafe.Pointer(&ap[0])), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX))
}

// Ctpsv Multiplies a packed triangular matrix by a vector
func Ctpsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []complex64, x []complex64, incX int) {
	C.cblas_ctpsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ztpsv Multiplies a packed triangular matrix by a vector
func Ztpsv(order Order, upLo UpLo, trans Transpose, diag Diag, n int, ap []complex128, x []complex128, incX int) {
	C.cblas_ztpsv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.enum_CBLAS_TRANSPOSE(trans), C.enum_CBLAS_DIAG(diag),
		C.blasint(n), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.blasint(incX))
}

// Ssymv computes y := alpha*A*x + beta*y for real symmetric matrix A
func Ssymv(order Order, upLo UpLo, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_ssymv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), C.float(beta), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Dsymv computes y := alpha*A*x + beta*y for real symmetric matrix A
func Dsymv(order Order, upLo UpLo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dsymv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), C.double(beta), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Chemv computes y := alpha*A*x + beta*y for complex Hermitian matrix A
func Chemv(order Order, upLo UpLo, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_chemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// Zhemv computes y := alpha*A*x + beta*y for complex Hermitian matrix A
func Zhemv(order Order, upLo UpLo, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zhemv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// Sspmv performs a symmetric packed matrix-vector multiplication.
func Sspmv(order Order, upLo UpLo, n int, alpha float32, ap []float32, x []float32, incX int, beta float32, y []float32, incY int) {
	C.cblas_sspmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&ap[0])), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), C.float(beta), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Dspmv performs a symmetric packed matrix-vector multiplication.
func Dspmv(order Order, upLo UpLo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
	C.cblas_dspmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&ap[0])), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), C.double(beta), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY))
}

// Sspr Perform a symmetric rank-1 update of a symmetric packed matrix with a real symmetric vector
func Sspr(order Order, upLo UpLo, n int, alpha float32, x []float32, incX int, ap []float32) {
	C.cblas_sspr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&ap[0])))
}

// Dspr Perform a symmetric rank-1 update of a symmetric packed matrix with a real symmetric vector
func Dspr(order Order, upLo UpLo, n int, alpha float64, x []float64, incX int, ap []float64) {
	C.cblas_dspr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&ap[0])))
}

// Chpr Perform a Hermitian rank-1 update of a Hermitian packed matrix with a complex Hermitian vector
func Chpr(order Order, upLo UpLo, n int, alpha float32, x []complex64, incX int, ap []complex64) {
	C.cblas_chpr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&ap[0]))
}

// Zhpr Perform a Hermitian rank-1 update of a Hermitian packed matrix with a complex Hermitian vector
func Zhpr(order Order, upLo UpLo, n int, alpha float64, x []complex128, incX int, ap []complex128) {
	C.cblas_zhpr(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&ap[0]))
}

// Sspr2 computes the symmetric rank 2 operation A := alpha*x*y**T + alpha*y*x**T + A,
func Sspr2(order Order, upLo UpLo, n int, alpha float32, x []float32, incX int, y []float32, incY int, ap []float32) {
	C.cblas_sspr2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.float(alpha), (*C.float)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.float)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.float)(unsafe.Pointer(&ap[0])))
}

// Dspr2 computes the symmetric rank 2 operation A := alpha*x*y**T + alpha*y*x**T + A,
func Dspr2(order Order, upLo UpLo, n int, alpha float64, x []float64, incX int, y []float64, incY int, ap []float64) {
	C.cblas_dspr2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.double(alpha), (*C.double)(unsafe.Pointer(&x[0])), C.blasint(incX), (*C.double)(unsafe.Pointer(&y[0])), C.blasint(incY), (*C.double)(unsafe.Pointer(&ap[0])))
}

// Chpr2 performs the Hermitian rank 2 operation A := alpha*x*y**H + conjg(alpha)*y*x**H + A,
func Chpr2(order Order, upLo UpLo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, ap []complex64) {
	C.cblas_chpr2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&ap[0]))
}

// Zhpr2 performs the Hermitian rank 2 operation A := alpha*x*y**H + conjg(alpha)*y*x**H + A,
func Zhpr2(order Order, upLo UpLo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, ap []complex128) {
	C.cblas_zhpr2(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&y[0]), C.blasint(incY), unsafe.Pointer(&ap[0]))
}

// CHbmv performs the matrix-vector operation: y = alpha*A*x + beta*y.
func Chbmv(order Order, upLo UpLo, n, k int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_chbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.blasint(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// ZHbmv performs the matrix-vector operation: y = alpha*A*x + beta*y.
func Zhbmv(order Order, upLo UpLo, n, k int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zhbmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), C.blasint(k), unsafe.Pointer(&alpha), unsafe.Pointer(&a[0]), C.blasint(lda), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// CHpmv performs the matrix-vector operation: y = alpha*A*x + beta*y.
func Chpmv(order Order, upLo UpLo, n int, alpha complex64, ap []complex64, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	C.cblas_chpmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// ZHpmv performs the matrix-vector operation: y = alpha*A*x + beta*y.
func Zhpmv(order Order, upLo UpLo, n int, alpha complex128, ap []complex128, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	C.cblas_zhpmv(C.enum_CBLAS_ORDER(order), C.enum_CBLAS_UPLO(upLo), C.blasint(n), unsafe.Pointer(&alpha), unsafe.Pointer(&ap[0]), unsafe.Pointer(&x[0]), C.blasint(incX), unsafe.Pointer(&beta), unsafe.Pointer(&y[0]), C.blasint(incY))
}

// Sgemm computes the matrix-matrix product where the matrices are
func Sgemm(Order Order, TransA Transpose, TransB Transpose, m int, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	C.cblas_sgemm(C.enum_CBLAS_ORDER(Order), C.enum_CBLAS_TRANSPOSE(TransA), C.enum_CBLAS_TRANSPOSE(TransB), C.blasint(m), C.blasint(n), C.blasint(k), C.float(alpha), (*C.float)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.float)(unsafe.Pointer(&b[0])), C.blasint(ldb), C.float(beta), (*C.float)(unsafe.Pointer(&c[0])), C.blasint(ldc))
}

// Dgemm computes the matrix-matrix product where the matrices are
func Dgemm(Order Order, TransA Transpose, TransB Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	C.cblas_dgemm(C.enum_CBLAS_ORDER(Order), C.enum_CBLAS_TRANSPOSE(TransA), C.enum_CBLAS_TRANSPOSE(TransB), C.blasint(m), C.blasint(n), C.blasint(k), C.double(alpha), (*C.double)(unsafe.Pointer(&a[0])), C.blasint(lda), (*C.double)(unsafe.Pointer(&b[0])), C.blasint(ldb), C.double(beta), (*C.double)(unsafe.Pointer(&c[0])), C.blasint(ldc))
}

func main() {
	list := []float64{1.1, 10, 1, 2, 3, 4, 5}
	fmt.Println(Idmin(len(list), list, 1))
}
