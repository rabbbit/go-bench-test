package main

func getSlice() []string {
	return []string{"a", "b", "c"}
}

//go:noinline
func test(s string) bool {
	for _, ss := range getSlice() {
		if s == ss {
			return true
		}
	}
	return false
}

//go:noinline
func test2(s string) bool {
	for _, ss := range append(getSlice(), "d") {
		if s == ss {
			return true
		}
	}
	return false
}
