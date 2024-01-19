package test

const (
	fint   = 40
	ffloat = 0.025
)

func decayI(n int) int {
	return n - n/fint
}

func decayF(n int) int {
	return int(float64(n) * (1 - ffloat))
}
