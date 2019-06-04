package cpu

import "../circuit"

type CPU struct {
	master circuit.Master
}

func NewCPU() *CPU {
	return &CPU{}
}
