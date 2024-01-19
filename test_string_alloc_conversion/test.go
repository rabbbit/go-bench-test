package test

func do(b [10]byte) string {
	var r string
	r = string(b[:9])
	return r
}
