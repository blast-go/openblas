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

func main() {
	fmt.Println(GetConfig())
}
