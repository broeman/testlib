package testlib

import (
	"fmt"	
)

type Testme struct {
	Input string	// uppercase becomes public
}

func (t Testme) Output() string {
	fmt.Println(t)
	return t.Input
}