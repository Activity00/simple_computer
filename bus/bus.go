package bus

import "../circuit"

const AddrWidth int = 16 // 地址总线位宽
const DataWidth int = 16 // 数据总线位宽
const ContrWidth int = 2 // 控制总线位宽

//总线顶层模块
type Bus struct {
	AddrWires  []circuit.Wire // 地址总线信号
	DataWires  []circuit.Wire // 数据总线信号
	ContrWires []circuit.Wire // 控制总线信号

	arbiter   *Arbiter
	masterMux *MasterMux
	slaveMux  *SlaveMux
	addrDec   *AddrDec
}

func NewBus() *Bus {
	bus := new(Bus)

	bus.AddrWires = make([]circuit.Wire, AddrWidth)
	for i := 0; i < AddrWidth; i++ {
		bus.AddrWires[i] = *circuit.NewWire("addr_"+string(i), false)
	}

	bus.DataWires = make([]circuit.Wire, DataWidth)
	for i := 0; i < DataWidth; i++ {
		bus.DataWires[i] = *circuit.NewWire("data_"+string(i), false)
	}

	bus.ContrWires = make([]circuit.Wire, ContrWidth)
	for i := 0; i < ContrWidth; i++ {
		bus.ContrWires[i] = *circuit.NewWire("contr_"+string(i), false)
	}

	bus.arbiter = NewArbiter(bus)
	bus.masterMux = NewMasterMux(bus)
	bus.slaveMux = NewSlaveMux(bus)
	bus.addrDec = NewAddrDec(bus)
	return bus
}

func (bus *Bus) SetAddrInputWire(index int, value bool) {
	bus.AddrWires[index].Update(value)
}

func (bus *Bus) GetAddrOutputWire(index int) bool {
	return bus.AddrWires[index].Get()
}

func (bus *Bus) SetAddrValue(value uint16) {
	// 总线 16 bit 设置
	var x = 0
	for i := AddrWidth - 1; i >= 0; i-- {
		// i << x 表示 当前位1其他0 eg： 1 << 4  0000 0001 << 4 => 0001 0000
		// value & x 表示取二进制各位&操作	eg： 1111 (15) & 1000 (8) => 1000 (8)
		r := value & (1 << uint16(x))
		if r != 0 {
			bus.SetAddrInputWire(i, true)
		} else {
			bus.SetAddrInputWire(i, false)
		}
		x++
	}
}

func (bus *Bus) String() string {
	result := ""
	for i := 0; i < AddrWidth; i++ {
		if bus.GetAddrOutputWire(i) {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}
