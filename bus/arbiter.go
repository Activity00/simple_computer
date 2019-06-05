package bus

import (
	"../circuit"
)

/**
总线仲裁器：对总线使用权进行调停

功能：接收总线主控发来的总线使用权请求，并且将使用权赋给合适的总线主控该模块总线仲裁器针对四个总线主控发来的请求进行调停。

实现原理：总线仲裁器根据目前所有者状态，按照有限状态机方式进行控制。总线状态机有四个状态，分别是
		 0号总线主控持有总线使用权， 1号总线主控持有总线使用权， 2号总线主控持有总线使用权，3号总线主控持有总线使用权
		针对总线使用权的调停使用轮询机制（按照请求顺序进行使用权分配，并且平等对待所有总线主控机制）eg:
        当前主控0 优先级 0 1 2 3
              1       1, 2, 3, 0
			  2       2, 3, 0, 1
   		      3       3, 0, 1, 2
**/

const (
	m0Grant = iota
	m1Grant
	m2Grant
	m3Grant
)

const (
	req0Grant = iota
	req1Grant
	req2Grant
	req3Grant
)

type Arbiter struct {
	bus *Bus

	owner uint8 // 当前主控 eg: 0, 1, 2, 3
}

func NewArbiter(bus *Bus) *Arbiter {
	biter := new(Arbiter)
	biter.bus = bus
	biter.owner = 0
	for i := uint8(0); i < MasterCh; i++ {
		bus.grantWire[i] = *circuit.NewWire("mGrant"+string(i), false)
	}
	return biter
}

func (biter *Arbiter) arbitrate() {

	switch biter.owner {
	case m0Grant:
		if biter.bus.reqWires[req0Grant].Get() == true {
			biter.bus.grantWire[m0Grant].Update(true)
			biter.owner = m0Grant
		} else if biter.bus.reqWires[req1Grant].Get() == true {
			biter.bus.grantWire[m1Grant].Update(true)
			biter.owner = m1Grant
		} else if biter.bus.reqWires[req2Grant].Get() == true {
			biter.bus.grantWire[m2Grant].Update(true)
			biter.owner = m2Grant
		} else if biter.bus.reqWires[req3Grant].Get() == true {
			biter.bus.grantWire[m3Grant].Update(true)
			biter.owner = m3Grant
		}
		break
	case m1Grant:
		if biter.bus.reqWires[req1Grant].Get() == true {
			biter.bus.grantWire[m1Grant].Update(true)
			biter.owner = m1Grant
		} else if biter.bus.reqWires[req2Grant].Get() == true {
			biter.bus.grantWire[m2Grant].Update(true)
			biter.owner = m2Grant
		} else if biter.bus.reqWires[req3Grant].Get() == true {
			biter.bus.grantWire[m3Grant].Update(true)
			biter.owner = m3Grant
		} else if biter.bus.reqWires[req0Grant].Get() == true {
			biter.bus.grantWire[m0Grant].Update(true)
			biter.owner = m0Grant
		}
		break
	case m2Grant:
		if biter.bus.reqWires[req2Grant].Get() == true {
			biter.bus.grantWire[m2Grant].Update(true)
			biter.owner = m2Grant
		} else if biter.bus.reqWires[req3Grant].Get() == true {
			biter.bus.grantWire[m3Grant].Update(true)
			biter.owner = m3Grant
		} else if biter.bus.reqWires[req0Grant].Get() == true {
			biter.bus.grantWire[m0Grant].Update(true)
			biter.owner = m0Grant
		} else if biter.bus.reqWires[req1Grant].Get() == true {
			biter.bus.grantWire[m1Grant].Update(true)
			biter.owner = m1Grant
		}
		break
	case m3Grant:
		if biter.bus.reqWires[req3Grant].Get() == true {
			biter.bus.grantWire[m3Grant].Update(true)
			biter.owner = m3Grant
		} else if biter.bus.reqWires[req0Grant].Get() == true {
			biter.bus.grantWire[m0Grant].Update(true)
			biter.owner = m0Grant
		} else if biter.bus.reqWires[req1Grant].Get() == true {
			biter.bus.grantWire[m1Grant].Update(true)
			biter.owner = m1Grant
		} else if biter.bus.reqWires[req2Grant].Get() == true {
			biter.bus.grantWire[m2Grant].Update(true)
			biter.owner = m2Grant
		}
		break
	}
}
