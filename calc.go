package goself

func Add(d ...int) int {
	sum := 0
	for _, v := range d {
		sum += v
	}
	return sum
}
