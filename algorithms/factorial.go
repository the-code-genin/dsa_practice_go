package algorithms

func Factorial(x int64) int64 {
	if x == 1 {
		return x
	} else {
		return x * Factorial(x - 1)
	}
}