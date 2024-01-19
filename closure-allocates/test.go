package test

type inner struct {
	a int
}

type outer struct {
	in []*inner
}

//go:noinline
func (o *outer) doSmth(f func(*inner)) {
	for _, i := range o.in {
		f(i)
	}
}

type worker struct {
	in *inner
}

func (w *worker) add(in *inner) {
	w.in.a += in.a
}

func run(x *outer, w *worker) int {
	x.doSmth(func(i *inner) {
		w.add(i)
	})
	return w.in.a
}
