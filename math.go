package math

import (
	"C"
	m "math"
)

func fabs(x C.double) C.double {
	return C.double(m.Abs(float64(x)))
}

func acos(x C.double) C.double {
	return C.double(m.Acos(float64(x)))
}

func acosf(x C.double) C.double {
	return C.double(m.Acos(float64(x)))
}

func acosh(x C.double) C.double {
	return C.double(m.Acosh(float64(x)))
}

func asin(x C.double) C.double {
	return C.double(m.Asin(float64(x)))
}

func asinh(x C.double) C.double {
	return C.double(m.Asinh(float64(x)))
}

func atan(x C.double) C.double {
	return C.double(m.Atan(float64(x)))
}

func atanh(x C.double) C.double {
	return C.double(m.Atanh(float64(x)))
}

func atan2(y C.double, x C.double) C.double {
	return C.double(m.Atan2(float64(y), float64(x)))
}

func cbrt(x C.double) C.double {
	return C.double(m.Cbrt(float64(x)))
}

func ceil(x C.double) C.double {
	return C.double(m.Ceil(float64(x)))
}

func copysign(x C.double, y C.double) C.double {
	return C.double(m.Copysign(float64(x), float64(y)))
}

func cos(x C.double) C.double {
	return C.double(m.Cos(float64(x)))
}

func cosh(x C.double) C.double {
	return C.double(m.Cosh(float64(x)))
}

func fdim(x C.double, y C.double) C.double {
	return C.double(m.Dim(float64(x), float64(y)))
}

func erf(x C.double) C.double {
	return C.double(m.Erf(float64(x)))
}

func erfc(x C.double) C.double {
	return C.double(m.Erfc(float64(x)))
}

func exp(x C.double) C.double {
	return C.double(m.Exp(float64(x)))
}

func exp2(x C.double) C.double {
	return C.double(m.Exp2(float64(x)))
}

func expm1(x C.double) C.double {
	return C.double(m.Expm1(float64(x)))
}

func floor(x C.double) C.double {
	return C.double(m.Floor(float64(x)))
}

func frexp(f C.double, expptr *C.int) C.double {
	frac, exp := m.Frexp(float64(f))
	*expptr = C.int(exp)
	return C.double(frac)
}

func log(x C.double) C.double {
	return C.double(m.Log(float64(x)))
}

func log10(x C.double) C.double {
	return C.double(m.Log10(float64(x)))
}

func fmax(x C.double, y C.double) C.double {
	return C.double(m.Max(float64(x), float64(y)))
}

func fmin(x C.double, y C.double) C.double {
	return C.double(m.Min(float64(x), float64(y)))
}

func pow(x C.double, y C.double) C.double {
	return C.double(m.Pow(float64(x), float64(y)))
}

func round(x C.double) C.double {
	return C.double(m.Round(float64(x)))
}

func sqrt(x C.double) C.double {
	return C.double(m.Sqrt(float64(x)))
}

func sin(x C.double) C.double {
	return C.double(m.Sin(float64(x)))
}

func sinh(x C.double) C.double {
	return C.double(m.Sinh(float64(x)))
}

func tan(x C.double) C.double {
	return C.double(m.Tan(float64(x)))
}

func tanh(x C.double) C.double {
	return C.double(m.Tanh(float64(x)))
}
