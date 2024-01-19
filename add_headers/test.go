package test

type store1 struct {
	store1 map[string]string
}

func newStore1(capacity int) *store1 {
	s := store1{
		store1: make(map[string]string, 1),
	}
	return &s
}

func (s *store1) with(k, v string) *store1 {
	s.store1[k] = v
	return s
}

func do1(out map[string]string, k, v string) {
	s := newStore1(1).with(k, v)
	out[k] = s.store1[k]
}

type store2 struct {
	store1 map[string]string
	store2 map[string]string
}

func newStore2(capacity int) *store2 {
	s := store2{
		store1: make(map[string]string, 1),
		store2: make(map[string]string, 1),
	}
	return &s
}

func (s *store2) with(k, v string) *store2 {
	s.store1[k] = v
	s.store2[k] = v
	return s
}

func do2(out map[string]string, k, v string) {
	s := newStore2(1).with(k, v)
	out[k] = s.store1[k]
}

type doer struct{}

func (d *doer) doMore(out map[string]string) {
	do2(out, "me", "hello")
}
