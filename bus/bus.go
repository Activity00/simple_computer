package bus

//总线顶层模块
type Bus struct {
	arbiter   *Arbiter
	masterMux *MasterMux
	slaveMux  *SlaveMux
	addrDec   *AddrDec
}

func NewBus() *Bus {
	return &Bus{arbiter: NewArbiter(),
		masterMux: NewMasterMux(),
		slaveMux:  NewSlaveMux(),
		addrDec:   NewAddrDec()}
}
