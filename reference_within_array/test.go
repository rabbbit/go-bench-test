package test

type tstruct struct {
	n int
}

func test1() tstruct {
	a := make([]tstruct, 1000)

	for i := 0; i < 1000; i++ {
		a[i] = tstruct{n: i}
	}
	return a[500]
}

func testAppend() tstruct {
	a := make([]tstruct, 0, 1000)

	for i := 0; i < 1000; i++ {
		a = append(a, tstruct{n: i})
	}
	return a[500]
}

func testPointer() *tstruct {
	a := make([]*tstruct, 1000)

	for i := 0; i < 1000; i++ {
		a[i] = &tstruct{n: i}
	}
	return a[500]
}

func testPointer2() *tstruct {
	a := make([]*tstruct, 0, 1000)

	for i := 0; i < 1000; i++ {
		a = append(a, &tstruct{n: i})
	}
	return a[500]
}

func testPointer3() *tstruct {
	a := make([]tstruct, 1000)

	for i := 0; i < 1000; i++ {
		a[i] = tstruct{n: i}
	}
	return &a[500]
}

func testPointer4() *tstruct {
	a := make([]tstruct, 1000)

	for i := 0; i < 1000; i++ {
		a[i] = tstruct{n: i}
	}
	b := a[500]
	return &b
}
