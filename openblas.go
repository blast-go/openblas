// MIT
package main

/*
#include "cblas.h"
#include <stdlib.h>
*/
import "C"

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
