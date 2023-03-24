# Go Bindings for OpenBLAS ![version-0.3.21](https://img.shields.io/badge/version-0.3.21-lightgrey.svg) [![GoDoc](https://godoc.org/github.com/blast-go/openblas?status.svg)](https://godoc.org/github.com/blast-go/openblas)

<a href="https://www.buymeacoffee.com/mwahlmann" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>

Package [openblas](https://github.com/blast-go/openblas) provides Go bindings 
for [OpenBLAS](https://www.openblas.net/) — an optimized BLAS library based on 
GotoBLAS2.

## How to use

Using OpenBLAS routines in Go is straightforward just import the package and call any routine.

```go
import "github.com/blast-go/openblas"
```

## Caveats

To use this library you need `CGO_ENABLED=1` and have OpenBLAS library installed
in your system.