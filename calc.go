package goself

func Add(d ...int) int {
	sum := 0
	for _, v := range d {
		sum += v
	}
	return sum
}

func Sub(i, j int) int {
	return i - j
}