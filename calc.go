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

func Plus(i, j int) int {
	return i + j
}

func Oh() {}

func Test() {}

func Self() string {
	return "this is master"
}

func Hey() {}

func Print() {}
