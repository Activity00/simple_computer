package utils

import "../bus"

type MasterInterface interface {
	SendRequestWire(value bool)
	GrandSuccess()
}

// 主控设备抽象 所有主控设备如cpu的抽象
type Master struct {
	bus *bus.Bus
}

func (master *Master) SendRequestWires(value bool) {
	master.bus.ArbitrateWire(master, value)
}

func (master *Master) GrandSuccess() {

}
