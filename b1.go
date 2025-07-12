package gomod1B

import gomod1C_v6 "github.com/relax-space/gomod1C/v6"

func AddB(a, b int) int {
	return a + b
}

func MultiB(a, b int) int {
	return gomod1C_v6.MultiC(a, b)
}
