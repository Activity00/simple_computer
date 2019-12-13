package cpu

import (

	"simple_computer/bus"
	"simple_computer/circuit"
	"simple_computer/io"
	"simple_computer/memory"
)

type CPU struct {
	master circuit.Master
}

func NewCPU(bus *bus.Bus, memory *memory.Memory64K) *CPU {
	return &CPU{}
}
func (c *CPU) ConnectPeripheral(p io.Peripheral) {

}



