package fibonacho

func Fibonalgas(hastaAqui int) (int, []int) {
	a := 1
	b := 0
	temp := 0
	var sequence []int

	for i := 1; i <= hastaAqui; i++ {
		temp = a
		a = a + b
		b = temp
		sequence = append(sequence, b)
	}
	return b, sequence
}
