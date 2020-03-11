package ccall

// extern double callSqrt(double x);
import "C"

func Sqrt(x float64) float64 {
	return float64(C.callSqrt(C.double(x)))
}
