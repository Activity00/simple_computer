package bus

/**
总线仲裁器：对总线使用权进行调停

功能：接收总线主控发来的总线使用权请求，并且将使用权赋给合适的总线主控该模块总线仲裁器针对四个总线主控发来的请求进行调停。

实现原理：总线仲裁器根据目前所有者状态，按照有限状态机方式进行控制。总线状态机有四个状态，分别是
		 0号总线主控持有总线使用权， 1号总线主控持有总线使用权， 2号总线主控持有总线使用权，3号总线主控持有总线使用权
		针对总线使用权的调停使用轮询机制（按照请求顺序进行使用权分配，并且平等对待所有总线主控机制）
**/

type Arbiter struct {
	m0Grant bool
	m1Grant bool
	m2Grant bool
	m3Grant bool

	bus *Bus
}

func NewArbiter(bus *Bus) *Arbiter {
	biter := new(Arbiter)
	biter.m0Grant = false
	biter.m1Grant = false
	biter.m2Grant = false
	biter.m3Grant = false
	biter.bus = bus
	return biter
}

func (biter *Arbiter) update(owner int) {
	switch owner {
	case 0:
		biter.m0Grant = true
		break
	case 1:
		biter.m0Grant = true
		break
	case 2:
		biter.m0Grant = true
		break
	case 3:
		biter.m0Grant = true
		break
	default:
		break

	}
}
