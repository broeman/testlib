package testlib

type Testme struct {
	input string
}

// getter
func (t *Testme) Input() string {
	return t.input
}

// setter
func (t *Testme) SetInput(input string) {
	t.input = input
}