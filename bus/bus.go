package bus

import (
	"../circuit"
	"../utils"
)

const AddrWidth int = 30 // 地址总线位宽
const DataWidth int = 32 // 数据总线位宽
const MasterCh uint8 = 4

//总线顶层模块
type Bus struct {
	/**总线信号线**/
	// 控制总线
	clkWire   circuit.Wire   // 同步信号
	reqWires  []circuit.Wire // 请求总线使用权
	grantWire []circuit.Wire // 总线使用许可信号
	csWire    circuit.Wire   // 从属访问选择信号
	asWire    circuit.Wire   // 访问有效表示信号
	rwWire    circuit.Wire   // 访问方式表示信号
	rdyWire   circuit.Wire   // 访问结束表示信号
	// 地址总线信号
	addrWires []circuit.Wire
	// 数据总线
	rdDataWires []circuit.Wire // 读取数据
	wrDataWires []circuit.Wire // 写入数据

	masters map[uint8]*utils.Master // 主控

	arbiter   *Arbiter //总线仲裁器
	masterMux *MasterMux
	slaveMux  *SlaveMux
	addrDec   *AddrDec
}

func NewBus(width int) *Bus {
	bus := new(Bus)

	bus.clkWire = *circuit.NewWire("clk", false)
	//for i := uint8(0); i < MasterCh; i++ {
	//	bus.reqWires[i] = *circuit.NewWire("req", false)
	//}

	bus.csWire = *circuit.NewWire("cs", false)
	bus.asWire = *circuit.NewWire("as", false)
	bus.rwWire = *circuit.NewWire("rw", false)

	bus.addrWires = make([]circuit.Wire, AddrWidth)
	for i := 0; i < AddrWidth; i++ {
		bus.addrWires[i] = *circuit.NewWire("addr_"+string(i), false)
	}

	bus.rdDataWires = make([]circuit.Wire, DataWidth)
	for i := 0; i < DataWidth; i++ {
		bus.rdDataWires[i] = *circuit.NewWire("rd"+string(i), false)
	}
	bus.wrDataWires = make([]circuit.Wire, DataWidth)
	for i := 0; i < DataWidth; i++ {
		bus.wrDataWires[i] = *circuit.NewWire("wr"+string(i), false)
	}

	bus.arbiter = NewArbiter(bus)
	bus.masterMux = NewMasterMux(bus)
	bus.slaveMux = NewSlaveMux(bus)
	bus.addrDec = NewAddrDec(bus)
	return bus
}

// 同步信号
func (bus *Bus) SyncClockWire(value bool) {
	bus.clkWire.Update(value)
}

// 总线主控输出信号
func (bus *Bus) MasterInputWires(addr int, cs bool, as bool, rw bool, wrData int32) {

}

//总线从属输出信号
func (bus *Bus) SlaveOutputWires(addr int, cs bool, as bool, rw bool, wrData int32) {

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

func (bus *Bus) SetDataValue(value uint16) {
	// 总线 16 bit 设置
	var x = 0
	for i := DataWidth - 1; i >= 0; i-- {
		// i << x 表示 当前位1其他0 eg： 1 << 4  0000 0001 << 4 => 0001 0000
		// value & x 表示取二进制各位&操作	eg： 1111 (15) & 1000 (8) => 1000 (8)
		r := value & (1 << uint16(x))
		if r != 0 {
			//bus.SetDataInputWire(i, true)
		} else {
			//bus.SetDataInputWire(i, false)
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

func (bus *Bus) SetAddrInputWire(index int, value bool) {
	bus.addrWires[index].Update(value)
}

func (bus *Bus) GetAddrOutputWire(index int) bool {
	return bus.addrWires[index].Get()
}
