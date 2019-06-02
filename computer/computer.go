package computer

type Computer struct {
	memory *memory.Memory64K
	cpu    *cpu.CPU
	bus    *Bus

	displayAdapter  *io.DisplayAdapter
	screenControl   *io.ScreenControl
	keyboardAdapter *io.KeyboardAdapter

	screenChannel chan *[160][240]byte
	quitChannel   chan bool
}
