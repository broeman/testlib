package testlib

type Testme struct {
	input string
}

func (t Testme) output() string {
	return t.input
}