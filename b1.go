package gomod1B

import gomod1C_v4 "github.com/relax-space/gomod1C/v4"

func AddB(a, b int) int {
	return a + b
}

func MultiB(a, b int) int {
	return gomod1C_v4.MultiC(a, b)
}
