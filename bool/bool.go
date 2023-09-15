package boolean

func EvenNumber(a, b int) bool {

	var bl1, bl2 bool

	if a%2 == 0 {
		bl1 = true
	}
	if b%2 == 0 {
		bl2 = true
	}

	return bl2 == bl1
}
