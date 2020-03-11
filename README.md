# cgo-math
Generate libm bridge for cgo

# Example

( Put same example in `_example` directory )

### Working Tree

```
$ tree
.
├── ccall
│   ├── ccall.c
│   └── ccall.go
├── go.mod
├── go.sum
└── main.go
```

### Contents

- `go.mod`

```
module example

go 1.12
```

- `main.go`

Only call `ccall.Sqrt(float64)float64` .

```go
package main

import (
	"example/ccall"
	"fmt"
)

func main() {
	fmt.Println(ccall.Sqrt(4))
}
```

- `ccall/ccall.go`

define `Sqrt(float64)float64` and call C API ( `double callSqrt(double x)` )

```go
package ccall

// extern double callSqrt(double x);
import "C"

func Sqrt(x float64) float64 {
	return float64(C.callSqrt(C.double(x)))
}
```

- `ccall/ccall.g`

define `double callSqrt(double)` and call `sqrt`

```c
#include "_cgo_export.h"

double callSqrt(double x) {
  return sqrt(x);
}
```

### Run example

```bash
$ go run main.go
# example/ccall
ccall.c:4:10: warning: implicitly declaring library function 'sqrt' with type 'double (double)' [-Wimplicit-function-declaration]
ccall.c:4:10: note: include the header <math.h> or explicitly provide a declaration for 'sqrt'
2
```

Produce **warning: implicitly declaring** because this example doesn't link `libm` .

### Executes cgo-math-bindgen

#### Install CLI

```bash
go get github.com/goccy/cgo-math/cmd/cgo-math-bindgen
```

#### Run generator

```bash
$ cgo-math-bindgen --package ccall --output ccall
```

Generate bridge source to `ccall/bridge_math.go` as `ccall` package

### Try to run example again

```bash
$ go run main.go
2
```

Yeah ! we resolved reference to `sqrt` symbol !
