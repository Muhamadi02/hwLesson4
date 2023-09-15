package integer

func Sum(a int) (int, int) {

	var b, c, d, e, sum, mult int
	b = a / 100
	c = a % 100
	d = c / 10
	e = c % 10

	sum = b + d + e
	mult = b * d * e

	return sum, mult

}

func Reverse(a int) int {
	var b, c, d, e, rev int
	b = a / 100
	c = a % 100
	d = c / 10
	e = c % 10

	rev = e*100 + d*10 + b
	return rev

}

func Int123132(a int) int {

	var b, c, d, e, mult int
	b = a / 100
	c = a % 100
	d = c / 10
	e = c % 10

	mult = b*100 + e*10 + d
	return mult
}
