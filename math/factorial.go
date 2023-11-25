package math

func Factorial(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
