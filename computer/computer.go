package computer

import (
	"log"
	"simple_computer/bus"
	"simple_computer/cpu"
	"simple_computer/io"
	"simple_computer/memory"
)

type Computer struct {
	bus    *bus.Bus
	cpu    *cpu.CPU
	memory *memory.Memory64K

	displayAdapter  *io.DisplayAdapter
	screenControl   *io.ScreenControl
	keyboardAdapter *io.KeyboardAdapter

	screenChannel chan *[160][240]byte
	quitChannel   chan bool

}

func NewComputer(screenChannel chan *[160][240]byte, quitChannel chan bool) *Computer  {
	c := new(Computer)
	c.screenChannel = screenChannel
	c.quitChannel = quitChannel

	c.bus = bus.NewBus(16)
	c.memory = memory.NewMemory64K(c.bus)
	c.cpu = cpu.NewCPU(c.bus, c.memory)

	c.keyboardAdapter = io.NewKeyboardAdapter()
	c.cpu.ConnectPeripheral(c.keyboardAdapter)

	c.displayAdapter = io.NewDisplaydAdapter()
	c.screenControl = io.NewScreenControl(c.displayAdapter, c.screenChannel, c.quitChannel)
	c.cpu.ConnectPeripheral(c.displayAdapter)

	return c

}

func (c *Computer) PowerUp(){
	log.Println(" power up ...")
	

}
