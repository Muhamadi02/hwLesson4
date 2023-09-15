package begin

import "math"

func Calc(a, b float64) (float64, float64, float64, float64) {

	var sum, min, mult, div float64
	sum = math.Abs(a) + math.Abs(b)
	min = math.Abs(a) - math.Abs(b)
	mult = math.Abs(a) * math.Abs(b)
	div = math.Abs(a) / math.Abs(b)

	return sum, min, mult, div
}

func Triangle(a, b float64) (float64, float64) {

	var c, P float64
	c = math.Sqrt((a * a) + (b * b))
	P = a + b + c
	return c, P
}

func Circle(l float64) (float64, float64) {
	var R, S float64
	R = l / (2 * 3.14)
	S = 2 * 3.14 * (R * R)
	return R, S

}
